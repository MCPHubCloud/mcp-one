package registry

import "mcphub.cloud/mcp-one/pkg/types"

// 定义mcpserver 的注册信息
type ServerRegistryInfo struct {
	Enable    bool               `json:"enable" yaml:"enable"`
	Name      string             `json:"name" yaml:"name"`
	TransType types.TransortType `json:"trans_type" yaml:"transType"`
	//for sse
	Url string `json:"url" yaml:"url"`

	//for stdio
	Command string   `json:"command" yaml:"command"`
	Env     []string `json:"env" yaml:"env"`
	Args    []string `json:"args" yaml:"args"`
}

type RegistryProvider interface {
	GetRegisteredServers() []ServerRegistryInfo
}
