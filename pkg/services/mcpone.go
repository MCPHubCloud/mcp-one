package services

import (
	"context"
	"errors"
	"fmt"
	mcpclient "github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	mcpserver "github.com/mark3labs/mcp-go/server"
	"log"
	"log/slog"
	"mcphub.cloud/mcp-one/pkg/config"
	"mcphub.cloud/mcp-one/pkg/registry"
	"mcphub.cloud/mcp-one/pkg/types"
	"strings"
	"time"
)

//AllInOne MCpSerice that implement mcp service proto

const (
	MCPONE_TOOL_PREFIX string = "mcpone@"
)

func extractToolName(s string) (string, error) {
	if strings.HasPrefix(s, MCPONE_TOOL_PREFIX) {
		return strings.TrimPrefix(s, MCPONE_TOOL_PREFIX), nil
	}
	return "", errors.New("tool name not recognized by mcpone")
}

type MCPOneServer struct {
	provider     registry.RegistryProvider
	serverConfig *config.McpOneConfig
	server       *mcpserver.MCPServer

	instances map[string]*types.ServerInstance //all managed servers
	tools     map[string]*types.ServerInstance //toolname -> instance
}

func NewMCPOneServer(oneServerConfig *config.McpOneConfig) *MCPOneServer {
	oneServer := &MCPOneServer{
		server:       mcpserver.NewMCPServer(oneServerConfig.Name, "1.0.0"),
		instances:    make(map[string]*types.ServerInstance),
		serverConfig: oneServerConfig,
		tools:        make(map[string]*types.ServerInstance),
	}

	if oneServerConfig.ProviderType == config.LocalProvider {
		oneServer.provider = registry.NewConfigProvider(oneServerConfig.McpServerConfigFile)
	}

	return oneServer
}

func (m *MCPOneServer) Start() {
	baseUrl := m.serverConfig.GetBaseUrlOrDefault("http://localhost:9090")
	sse := mcpserver.NewSSEServer(m.server, mcpserver.WithBaseURL(baseUrl))
	log.Printf("MCP-One server listening on: %s", baseUrl)
	if err := sse.Start(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func (m *MCPOneServer) GetActiveServers() {
	//TODO user can disable server in progress
}

func (m *MCPOneServer) LoadAllServers() {
	registeredServer := m.provider.GetRegisteredServers()
	for _, registryInfo := range registeredServer {
		m.registerServer(registryInfo)
	}
}

func (m *MCPOneServer) getClient(toolName string) mcpclient.MCPClient {
	instance, ok := m.tools[toolName]
	if ok {
		return instance.Client
	}
	return nil
}

func (m *MCPOneServer) addToolForMcpOneServer(origTool mcp.Tool, instance *types.ServerInstance) {
	newTool := origTool
	newTool.Name = MCPONE_TOOL_PREFIX + newTool.Name
	m.tools[newTool.Name] = instance
	m.server.AddTool(newTool, m.callTool)
}

func (m *MCPOneServer) callTool(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := request.Params.Name
	client := m.getClient(request.Params.Name)

	backendToolName, err := extractToolName(toolName)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Failed call tool  %s!", err)), err
	}

	request.Params.Name = backendToolName
	log.Printf("Call tool name: %+v", request)
	if client != nil {
		ret, err := client.CallTool(ctx, request)
		if err != nil {
			log.Fatalf("Call tool error: %v", err)
		} else {
			log.Printf("Call tool result: %+v", ret)
			return ret, nil
		}
	}
	return mcp.NewToolResultText(fmt.Sprintf("Failed call tool  %s!", request.Params.Name)), nil
}

func (m *MCPOneServer) registerServer(registry registry.ServerRegistryInfo) {
	//建立客户端，更新链接状态
	var client mcpclient.MCPClient
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	switch registry.TransType {
	case types.TransportSSE:
		{
			if c, err := mcpclient.NewSSEMCPClient(registry.Url); err != nil {
				slog.Error("failed register mcp server for %s, %s", registry.Name, err.Error())
				return
			} else {
				if err := c.Start(ctx); err != nil {
					slog.Error("failed start sse client for %s, %s", registry.Name, err.Error())
					return
				}
				client = c
			}
		}

	case types.TransportStdio:
		if c, err := mcpclient.NewStdioMCPClient(registry.Command, registry.Env, registry.Args...); err != nil {
			slog.Error("failed register mcp server for ", registry.Name)
			return
		} else {
			client = c
		}
	}

	//add to server instance info lists
	serverInstance := types.NewServerInstance(registry.Name)

	//connect
	fmt.Println("Initializing client...")
	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mcpone-client-" + registry.Name,
		Version: "1.0.0",
	}

	initResult, err := client.Initialize(ctx, initRequest)
	if err != nil {
		slog.Error("Failed to initialize: %v", err)
	}

	serverInstance.SetConnected(client)

	fmt.Printf(
		"Initialized with server: %s %s\n\n",
		initResult.ServerInfo.Name,
		initResult.ServerInfo.Version,
	)

	//fetch tools for this server
	fmt.Println("Listing available tools...")
	toolsRequest := mcp.ListToolsRequest{}
	tools, err := client.ListTools(ctx, toolsRequest)
	if err != nil {
		log.Fatalf("Failed to list tools: %v", err)
	}
	for _, tool := range tools.Tools {
		fmt.Printf("- %s: %s\n", tool.Name, tool.Description)
		serverInstance.AddTools(tool)

		//add modified tool mcpone
		m.addToolForMcpOneServer(tool, serverInstance)
	}
	fmt.Println()
}
