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
		Name:      "mcp-timeserver",
		TransType: types.TransportStdio,
		Command:   "/Users/barry/UserApps/anaconda3/anaconda3/bin/uvx",
		Args:      []string{"mcp-timeserver"},
	})

	return config
}
