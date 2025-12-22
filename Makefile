.PHONY: help build run-api run-job run-consumer test clean docker-build

# 默认目标
help:
	@echo "IDRM 项目 Makefile"
	@echo ""
	@echo "可用命令:"
	@echo "  make build           - 构建所有服务"
	@echo "  make run-api         - 运行API服务"
	@echo "  make run-job         - 运行定时任务服务"
	@echo "  make run-consumer    - 运行消费者服务"
	@echo "  make test            - 运行测试"
	@echo "  make clean           - 清理构建文件"
	@echo "  make docker-build    - 构建Docker镜像"
	@echo "  make init-db         - 初始化数据库"

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
