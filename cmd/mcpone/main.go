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
	var provider string
	var mcpServerConfigFile string
	var baseUrl string
	var name string

	var rootCmd = &cobra.Command{
		Use:   "mcpone -c mcpserver-config.yaml",
		Short: "mcpone -c mcpserver-config.yaml",
		Run: func(cmd *cobra.Command, args []string) {
			conf, err := utils.ReadAndParseFile[config.McpOneConfig](filePath)

			if name != "" {
				conf.Name = name
			}

			//commandline params overwrite config params
			if mcpServerConfigFile != "" {
				conf.McpServerConfigFile = mcpServerConfigFile
			}

			if provider != "" {
				conf.ProviderType = config.ProviderType(provider)
			}

			if baseUrl != "" {
				conf.BaseUrl = baseUrl
			}

			if err != nil {
				fmt.Println(err)
				return
			}

			mcpServer := services.NewMCPOneServer(conf)
			mcpServer.LoadAllServers()
			mcpServer.Start()
		},
	}

	rootCmd.Flags().StringVarP(&name, "name", "n", "mcpone", "config file of mcpone")
	rootCmd.Flags().StringVarP(&filePath, "config", "c", "mcpone-config.yaml", "config file of mcpone")
	rootCmd.Flags().StringVarP(&mcpServerConfigFile, "", "", "", "mcpServers list config")
	rootCmd.Flags().StringVarP(&provider, "provider", "p", "local", "current only support [local] provider")
	rootCmd.Flags().StringVarP(&baseUrl, "baseurl", "", "", "mcpoone server listen address")

	if err := rootCmd.MarkFlagRequired("config"); err != nil {
		fmt.Println("needed config file for mcpone:", err)
		return
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("failed run mcpone sever", err)
	}
}
