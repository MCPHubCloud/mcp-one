# 第一阶段：编译阶段
FROM golang:1.21-alpine AS builder

# 安装必要的工具
RUN apk add --no-cache git

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译二进制文件
RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mcp-one cmd/mcpone/main.go

# 第二阶段：运行阶段
FROM alpine:latest

# 安装依赖
RUN apk --no-cache add ca-certificates

# 创建工作目录
WORKDIR /app

# 从编译阶段复制二进制文件
COPY --from=builder /app/mcp-one .
COPY --from=builder /app/examples .

# 运行 mcp-one 命令
CMD ["./mcp-one", "--log_file=/tmp/mcpone.log", "-c", "/app/examples/mcpone-config.yaml", "--logtostderr=false", "-v", "5"]
