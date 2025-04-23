# mcp-one
A unified All-in-One entrance for mcpservers, which manages various types of rated mcpservers.

🤔What you need is only one mcpserver! 🎉

# How to use
```bash
mcpone -c mcpserver-config.yaml

Usage:
  mcpone -c mcpserver-config.yaml [flags]

Flags:
      -- string                          mcpServers list config
      --add_dir_header                   If true, adds the file directory to the header of the log messages
      --alsologtostderr                  log to standard error as well as files (no effect when -logtostderr=true)
      --baseurl string                   mcpoone server listen address
  -c, --config string                    config file of mcpone (default "mcpone-config.yaml")
```

### start server
`./mcp-one -c ./mcpone-config.yaml --baseurl 0.0.0.0:9090 `
```
Add sse client to your client, http://localhost:9090/sse. You can also hosted in your server.

# Build
```bash
git clone https://github.com/MCPHubCloud/mcp-one.git
cd mcp-one
make build

# target output in _output/mcp-one





# Mcp Servers
mcp-one already supported servers, you can add more in mcpserver-config.yaml just like client
- time-mcp
- mcp-server-fetch
