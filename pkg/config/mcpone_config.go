package config

type ProviderType string

const (
	LocalProvider = "local"
	CloudProvider = "cloud"
)

type McpOneConfig struct {
	ProviderType        ProviderType `json:"provider_type" yaml:"provider_type"`
	McpServerConfigFile string       `json:"config_file" yaml:"mcpserver_config"`
}

func NewDefaultMcpOneConfig() *McpOneConfig {
	config := &McpOneConfig{
		ProviderType:        LocalProvider,
		McpServerConfigFile: "config.yaml",
	}
	return config
}
