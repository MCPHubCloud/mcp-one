# 第一阶段：编译阶段
FROM golang:1.24-alpine AS builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# 安装必要的工具
RUN apk update
RUN apk add --no-cache git

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY ../go.mod ../go.sum ./

# 下载依赖
ENV GO111MODULE "on"
ENV GOPROXY "https://goproxy.cn,direct"
RUN go mod download

# 复制源代码
COPY . .

# 编译二进制文件
RUN  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mcp-one cmd/mcpone/main.go

# 第二阶段：运行阶段
#FROM alpine:latest
#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
#RUN apk --no-cache add ca-certificates
FROM ghcr.io/astral-sh/uv:debian
#RUN pip install uvx -i https://pypi.tuna.tsinghua.edu.cn/simple

# 创建工作目录
WORKDIR /app

# 从编译阶段复制二进制文件
COPY --from=builder /app/mcp-one .
COPY --from=builder /app/examples .

ENV UV_DEFAULT_INDEX="https://mirrors.tuna.tsinghua.edu.cn/pypi/web/simple"
RUN echo 'Types: deb'  > /etc/apt/sources.list.d/debian.sources && \
    echo 'URIs: https://mirrors.tuna.tsinghua.edu.cn/debian' >>  /etc/apt/sources.list.d/debian.sources && \
    echo 'Suites: bookworm bookworm-updates bookworm-backports' >>  /etc/apt/sources.list.d/debian.sources && \
    echo 'Components: main contrib non-free non-free-firmware' >>  /etc/apt/sources.list.d/debian.sources && \
    echo 'Signed-By: /usr/share/keyrings/debian-archive-keyring.gpg' >>  /etc/apt/sources.list.d/debian.sources

RUN apt update && apt install -y pip nodejs npm

# 运行 mcp-one 命令
CMD ["./mcp-one", "--log_file=mcpone.log", "-c", "mcpone-config.yaml", "--logtostderr=true", "-v", "5"]
