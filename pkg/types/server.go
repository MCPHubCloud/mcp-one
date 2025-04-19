package types

import (
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

type TransortType string
type ServerStatus string

const (
	TransportSSE   TransortType = "sse"
	TransportStdio TransortType = "stdio"

	ServerConnecting ServerStatus = "connecting"
	ServerConnected  ServerStatus = "connected"
)

// 表示一个被托管的 mcpserver 实例, 状态管理/同步均反映在这里
type ServerInstance struct {
	Name       string
	Status     ServerStatus
	Client     client.MCPClient
	CreateTime time.Time
	Tools      map[string]mcp.Tool
	//mcp info, TODO resource/prompts
}

func NewServerInstance(name string) *ServerInstance {
	return &ServerInstance{
		Name:       name,
		Status:     ServerConnecting,
		Client:     nil,
		CreateTime: time.Now(),
		Tools:      make(map[string]mcp.Tool),
	}
}

func (s *ServerInstance) SetConnected(c client.MCPClient) {
	s.Client = c
	s.Status = ServerConnected
}

func (s *ServerInstance) AddTools(tool mcp.Tool) {
	s.Tools[tool.Name] = tool
}
