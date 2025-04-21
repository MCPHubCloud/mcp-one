package registry

import (
	"k8s.io/klog/v2"
	"mcphub.cloud/mcp-one/pkg/utils"
	"sync"
	"time"
)

type McpServerConfig struct {
	McpServers map[string]ServerRegistryInfo `json:"mcpServers" yaml:"mcpServers"`
}

type ConfigProvider struct {
	lock       sync.RWMutex
	configFile string
	config     *McpServerConfig
}

func NewConfigProvider(configFile string) *ConfigProvider {
	configProvider := &ConfigProvider{
		lock:       sync.RWMutex{},
		configFile: configFile,
		config: &McpServerConfig{
			McpServers: make(map[string]ServerRegistryInfo),
		},
	}

	configProvider.loadOnce()

	stopChan := make(chan struct{})
	go configProvider.StartMonitor(stopChan)

	return configProvider
}

func (p *ConfigProvider) GetRegisteredServers() []ServerRegistryInfo {
	p.lock.RLock()
	defer p.lock.RUnlock()

	copyServers := make([]ServerRegistryInfo, len(p.config.McpServers))
	i := 0
	for _, v := range p.config.McpServers {
		copyServers[i] = v
		i++
	}
	return copyServers
}

func (p *ConfigProvider) loadOnce() {
	p.lock.Lock()
	defer p.lock.Unlock()

	conf, err := utils.ReadAndParseFile[McpServerConfig](p.configFile)
	if err != nil {
		klog.Errorf("failed load config from %s ", p.configFile)
	}

	p.config = conf
}

func (p *ConfigProvider) StartMonitor(stop chan struct{}) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			p.loadOnce()
		}
	}
}
