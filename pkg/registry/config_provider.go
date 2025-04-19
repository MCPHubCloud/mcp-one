package registry

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"mcphub.cloud/mcp-one/pkg/config"
	"strings"
	"sync"
	"time"
)

type ConfigProvider struct {
	lock   sync.RWMutex
	config config.McpOneConfig
}

func (p *ConfigProvider) get_registy() []ServerRegistryInfo {
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

// 判断文件类型
func getFileType(filePath string) string {
	lowerCasePath := strings.ToLower(filePath)
	if strings.HasSuffix(lowerCasePath, ".json") {
		return "json"
	} else if strings.HasSuffix(lowerCasePath, ".yaml") || strings.HasSuffix(lowerCasePath, ".yml") {
		return "yaml"
	}
	return ""
}

func (p *ConfigProvider) loadOnce(filePath string) {
	p.lock.Lock()
	defer p.lock.Unlock()

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("read file %s error: %v", filePath, err)
	}

	fileType := getFileType(filePath)
	switch fileType {
	case "json":
		err = json.Unmarshal(data, &p.config)
	case "yaml":
		err = yaml.Unmarshal(data, &p.config)
	default:
		log.Errorf("unsupport file type")
	}
}

func (p *ConfigProvider) StartMonitor(stop chan struct{}) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			p.loadOnce("config.json")
		}
	}
}
