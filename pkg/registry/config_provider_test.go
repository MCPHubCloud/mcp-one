package registry

import (
	"github.com/stretchr/testify/assert"
	"mcphub.cloud/mcp-one/pkg/types"
	"sync"
	"testing"
)

func TestConfigProvider_loadOnce(t *testing.T) {
	type fields struct {
		lock   sync.RWMutex
		config McpServerConfig
	}
	type args struct {
		filePath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add examples cases.
		{
			name: "load_once",
			args: args{
				filePath: "../../examples/config.yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ConfigProvider{
				configFile: tt.args.filePath,
			}
			p.loadOnce()
			assert.Equal(t, 2, len(p.config.McpServers))
			assert.Equal(t, "mcp-fetch", p.config.McpServers["mcp-fetch"].Name)
			assert.Equal(t, types.TransportSSE, p.config.McpServers["mcp-fetch"].TransType)

			assert.Equal(t, "mcp-timeserver", p.config.McpServers["mcp-timeserver"].Name)
			assert.Equal(t, types.TransportStdio, p.config.McpServers["mcp-timeserver"].TransType)
			assert.Equal(t, 2, len(p.config.McpServers["mcp-timeserver"].Args))
		})
	}
}
