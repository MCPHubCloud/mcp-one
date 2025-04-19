package main

import (
	"mcphub.cloud/mcp-one/pkg/config"
	"mcphub.cloud/mcp-one/pkg/services"
)

func main() {
	mcpserver := services.NewMCPOneServer("mcp-one", config.NewMcpOneConfig())
	mcpserver.LoadAllServers()
	mcpserver.Start()
}
