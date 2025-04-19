package config

import (
	"mcphub.cloud/mcp-one/pkg/registry"
	"mcphub.cloud/mcp-one/pkg/types"
)

type McpOneConfig struct {
	McpServers map[string]registry.ServerRegistryInfo `json:"mcpServers"`
}

func NewMcpOneConfig() *McpOneConfig {
	config := &McpOneConfig{
		McpServers: make(map[string]registry.ServerRegistryInfo),
	}

	/*
		config.McpServers = append(config.McpServers, types.ServerRegistryInfo{
			Enable:    true,
			Name:      "fetch",
			TransType: types.TransportStdio,
			Command:   "/Users/barry/UserApps/anaconda3/anaconda3/bin/uvx",
			Args:      []string{"mcp-server-fetch"},
		})
	*/

	config.McpServers["mcp-fetch"] = registry.ServerRegistryInfo{
		Enable:    true,
		Name:      "mcp-fetch",
		TransType: types.TransportSSE,
		Url:       "http://101.200.75.13:8080/sse",
	}

	config.McpServers["mcp-timeserver"] = registry.ServerRegistryInfo{
		Enable:    true,
		Name:      "mcp-timeserver",
		TransType: types.TransportStdio,
		Command:   "/Users/barry/UserApps/anaconda3/anaconda3/bin/python3",
		Args:      []string{"-m", "mcp_simple_timeserver"},
	}

	return config
}
