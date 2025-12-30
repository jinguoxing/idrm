.PHONY: all build clean test lint run

# 项目名称
PROJECT_NAME := {{PROJECT_NAME}}

# Go 相关
GO := go
GOFLAGS := -v

# 目录
API_DIR := api
RPC_DIR := rpc

# 默认目标
all: build

# 编译
build:
	$(GO) build $(GOFLAGS) -o bin/$(PROJECT_NAME)-api $(API_DIR)/api.go

# 清理
clean:
	rm -rf bin/
	rm -rf logs/

# 测试
test:
	$(GO) test -v ./...

# 测试覆盖率
cover:
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html

# 代码检查
lint:
	golangci-lint run ./...

# 运行 API 服务
run:
	$(GO) run $(API_DIR)/api.go -f $(API_DIR)/etc/api.yaml

# 生成 API 代码
gen-api:
	goctl api go -api $(API_DIR)/doc/api.api -dir $(API_DIR)/

# 格式化代码
fmt:
	gofmt -s -w .
	goimports -w .

# 安装依赖
deps:
	$(GO) mod tidy
	$(GO) mod download

# 帮助
help:
	@echo "可用命令:"
	@echo "  make build    - 编译项目"
	@echo "  make clean    - 清理编译产物"
	@echo "  make test     - 运行测试"
	@echo "  make cover    - 生成测试覆盖率报告"
	@echo "  make lint     - 运行代码检查"
	@echo "  make run      - 运行 API 服务"
	@echo "  make gen-api  - 生成 API 代码"
	@echo "  make fmt      - 格式化代码"
	@echo "  make deps     - 安装依赖"
