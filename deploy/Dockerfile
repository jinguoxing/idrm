# 多阶段构建
FROM golang:1.24-alpine AS builder

# 设置 Go 代理加速(构建时参数)
ARG GOPROXY=https://goproxy.cn,direct
ENV GOPROXY=${GOPROXY}

# 设置工作目录
WORKDIR /app

# 安装必要工具
RUN apk add --no-cache git

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o idrm-api ./api/api.go

# 运行时镜像
FROM alpine:latest

# 安装必要工具
RUN apk --no-cache add ca-certificates tzdata curl

# 设置时区
ENV TZ=Asia/Shanghai

# 创建用户
RUN addgroup -g 1000 idrm && \
    adduser -D -u 1000 -G idrm idrm

# 设置工作目录
WORKDIR /app

# 从builder复制二进制文件
COPY --from=builder /app/idrm-api .

# 创建日志目录
RUN mkdir -p /app/logs && chown -R idrm:idrm /app

# 切换用户
USER idrm

# 暴露端口
EXPOSE 8888

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s --start-period=40s --retries=3 \
    CMD curl -f http://localhost:8888/health || exit 1

# 启动命令
CMD ["./idrm-api", "-f", "etc/api.yaml"]
