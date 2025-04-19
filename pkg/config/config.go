package config

import "mcphub.cloud/mcp-one/pkg/types"

type McpOneConfig struct {
	McpServers []types.ServerRegistryInfo `json:"mcp_servers"`
}

func NewMcpOneConfig() *McpOneConfig {
	config := &McpOneConfig{
		McpServers: []types.ServerRegistryInfo{},
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

	config.McpServers = append(config.McpServers, types.ServerRegistryInfo{
		Enable:    true,
		Name:      "mcp-fetch",
		TransType: types.TransportSSE,
		Url:       "http://101.200.75.13:8080/sse",
	})

	config.McpServers = append(config.McpServers, types.ServerRegistryInfo{
		Enable:    true,
		Name:      "mcp-timeserver",
		TransType: types.TransportStdio,
		Command:   "/Users/barry/UserApps/anaconda3/anaconda3/bin/python3",
		Args:      []string{"-m", "mcp_simple_timeserver"},
	})

	return config
}
