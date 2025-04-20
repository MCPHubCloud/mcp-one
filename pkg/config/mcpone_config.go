package config

type ProviderType string

const (
	LocalProvider = "local"
	CloudProvider = "cloud"
)

type McpOneConfig struct {
	Name                string       `yaml:"name"`
	ProviderType        ProviderType `json:"provider_type" yaml:"provider_type"`
	McpServerConfigFile string       `json:"config_file" yaml:"mcpserver_config"`
	BaseUrl             string       `json:"base_url" yaml:"base_url"`
}

func NewDefaultMcpOneConfig() *McpOneConfig {
	config := &McpOneConfig{
		ProviderType:        LocalProvider,
		McpServerConfigFile: "mcpserver-config.yaml",
		BaseUrl:             "http://localhost:9090",
	}
	return config
}

func (c McpOneConfig) GetBaseUrlOrDefault(url string) string {
	if c.BaseUrl == "" {
		return url
	}

	return c.BaseUrl
}
