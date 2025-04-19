package types

//定义mcpserver 的注册信息

type ServerRegistryInfo struct {
	Enable    bool         `json:"enable"`
	Name      string       `json:"name"`
	TransType TransortType `json:"trans_type"`
	//for sse
	Url string `json:"url"`

	//for stdio
	Command string   `json:"command"`
	Env     []string `json:"env"`
	Args    []string `json:"args"`
}
