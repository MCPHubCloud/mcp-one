package main

import (
	"fmt"
	"mcphub.cloud/mcp-one/pkg/config"
	"mcphub.cloud/mcp-one/pkg/services"
	"mcphub.cloud/mcp-one/pkg/utils"

	"github.com/spf13/cobra"
)

func main() {
	var filePath string

	var rootCmd = &cobra.Command{
		Use:   "mcpone -c mcpserver-config.yaml",
		Short: "mcpone -c mcpserver-config.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			conf, err := utils.ReadAndParseFile[config.McpOneConfig](filePath)
			if err != nil {
				fmt.Println(err)
				return
			}
			mcpserver := services.NewMCPOneServer("mcpone", conf)
			mcpserver.LoadAllServers()
			mcpserver.Start()
		},
	}

	rootCmd.Flags().StringVarP(&filePath, "config", "c", "mcpone-mcpserver-config.yaml", "config file  of mcpone server")
	if err := rootCmd.MarkFlagRequired("config"); err != nil {
		fmt.Println("needed config file:", err)
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("failed run mcpone sever", err)
	}
}
