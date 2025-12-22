.PHONY: help build run-api run-job run-consumer test clean docker-build api-gen swagger run

# 默认目标
help: ## 显示帮助信息
	@echo "可用的命令:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: api-gen
api-gen: ## 生成 API 代码
	cd api && goctl api go -api doc/api/api.api -dir .
	@echo "API 代码生成完成"

.PHONY: swagger
swagger: ## 生成 Swagger 文档
	@echo "生成 Swagger 文档..."
	@if command -v goctl-swagger >/dev/null 2>&1; then \
		cd api && goctl api plugin -plugin goctl-swagger="swagger -filename idrm.json -basepath /" \
			-api doc/api/api.api -dir doc/swagger; \
		echo "Swagger 文档已生成: api/doc/swagger/idrm.json"; \
	else \
		echo "错误: goctl-swagger 未安装"; \
		echo "请运行: go install github.com/zeromicro/goctl-swagger@latest"; \
		exit 1; \
	fi

.PHONY: run
run: ## 运行 API 服务
	cd api && go run api.go -f etc/api.yaml

# 构建所有服务
build:
	@echo "Building all services..."
	@go build -o bin/api-server cmd/api-server/main.go
	@go build -o bin/job-server cmd/job-server/main.go
	@go build -o bin/consumer-server cmd/consumer-server/main.go
	@echo "Build completed!"

# 运行API服务
run-api:
	@echo "Starting API server..."
	@go run cmd/api-server/main.go -f cmd/api-server/etc/api-server.yaml

# 运行定时任务服务
run-job:
	@echo "Starting Job server..."
	@go run cmd/job-server/main.go -f cmd/job-server/etc/job-server.yaml

# 运行消费者服务
run-consumer:
	@echo "Starting Consumer server..."
	@go run cmd/consumer-server/main.go -f cmd/consumer-server/etc/consumer-server.yaml

# 运行测试
test:
	@echo "Running tests..."
	@go test -v ./...

# 运行集成测试
test-integration:
	@echo "Running integration tests..."
	@go test -v ./test/integration/...

# 清理构建文件
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -rf logs/
	@echo "Clean completed!"

# 构建Docker镜像
docker-build:
	@echo "Building Docker images..."
	@docker build -f deploy/docker/Dockerfile.api-server -t idrm-api-server:latest .
	@docker build -f deploy/docker/Dockerfile.job-server -t idrm-job-server:latest .
	@docker build -f deploy/docker/Dockerfile.consumer-server -t idrm-consumer-server:latest .
	@echo "Docker build completed!"

# 初始化数据库
init-db:
	@echo "Initializing databases..."
	@sh scripts/init-db.sh
	@echo "Database initialization completed!"

# 生成代码
gen-proto:
	@echo "Generating proto code..."
	@sh scripts/gen-proto.sh
	@echo "Proto generation completed!"

# 代码格式化
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Format completed!"

# 代码检查
lint:
	@echo "Running linter..."
	@golangci-lint run
	@echo "Lint completed!"

# 安装依赖
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies installed!"

# 运行所有服务（开发模式）
dev:
	@echo "Starting all services in development mode..."
	@make -j3 run-api run-job run-consumer
