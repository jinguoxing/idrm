# IDRM 项目

基于 Go-Zero 的微服务架构项目，包含数据视图、数据理解和资源目录等业务模块。

## 📁 项目结构

```
idrm/
├── cmd/                    # 服务入口（3个独立服务）
│   ├── api-server/        # HTTP API服务
│   ├── job-server/        # 定时任务服务
│   └── consumer-server/   # 消息队列消费服务
├── app/                    # 应用层
│   ├── api/               # API层（简单HTTP接口）
│   ├── bff/               # BFF层（复杂业务聚合）
│   ├── job/               # 定时任务
│   ├── consumer/          # 消息队列消费者
│   └── rpc/               # RPC服务（可选）
├── domain/                 # 领域层（核心业务逻辑）
│   ├── data_view/         # 数据视图领域
│   ├── data_understanding/ # 数据理解领域
│   └── resource_catalog/  # 资源目录领域
├── infrastructure/         # 基础设施层
│   ├── persistence/       # 数据持久化
│   └── mq/                # 消息队列
└── pkg/                    # 公共包
    ├── config/            # 配置
    ├── errorx/            # 错误处理
    ├── response/          # 响应格式
    └── middleware/        # 中间件
```

## 🏗️ 架构设计

### 三层架构
- **API层**: 处理简单的CRUD操作，直接调用Domain服务
- **BFF层**: 处理复杂业务聚合，组合多个Domain逻辑
- **Domain层**: 核心业务逻辑，可被API、BFF、Job、Consumer共享

### 三个独立服务
| 服务 | 端口 | 职责 |
|-----|------|-----|
| api-server | 8080 | HTTP API服务 |
| job-server | - | 定时任务 |
| consumer-server | - | Kafka消息消费 |

### Domain模块
- **data_view**: 数据视图
- **data_understanding**: 数据理解  
- **resource_catalog**: 资源目录

## 🚀 快速开始

### 环境要求
- Go 1.21+
- MySQL 8.0+
- Redis 6.0+
- Kafka 2.8+

### 安装依赖
```bash
go mod download
```

### 配置数据库
```sql
CREATE DATABASE idrm_data_view DEFAULT CHARACTER SET utf8mb4;
CREATE DATABASE idrm_data_understanding DEFAULT CHARACTER SET utf8mb4;
CREATE DATABASE idrm_resource_catalog DEFAULT CHARACTER SET utf8mb4;
```

### 启动服务

#### 1. 启动API服务
```bash
make run-api
# 或
go run cmd/api-server/main.go -f cmd/api-server/etc/api-server.yaml
```

#### 2. 启动定时任务服务
```bash
make run-job
# 或
go run cmd/job-server/main.go -f cmd/job-server/etc/job-server.yaml
```

#### 3. 启动消费者服务
```bash
make run-consumer
# 或
go run cmd/consumer-server/main.go -f cmd/consumer-server/etc/consumer-server.yaml
```

## 📝 开发指南

### 添加新的Domain模块

1. 创建目录结构：
```bash
mkdir -p domain/new_domain/{entity,repository,service}
```

2. 定义实体（entity）：定义业务对象
3. 定义仓储接口（repository）：数据访问抽象
4. 实现领域服务（service）：核心业务逻辑
5. 在infrastructure层实现仓储

### 添加新的API接口

1. 在`app/api/[module]/handler/`创建handler
2. 在`app/api/[module]/routes.go`注册路由
3. 调用Domain服务处理业务逻辑

### 添加新的定时任务

1. 在`app/job/`创建任务目录
2. 实现任务逻辑
3. 在`cmd/job-server/main.go`注册任务

### 添加新的消费者

1. 在`app/consumer/`创建消费者目录
2. 实现MessageHandler接口
3. 在`cmd/consumer-server/main.go`注册消费者

## 🔧 配置说明

### 多数据库配置
每个Domain使用独立数据库，在配置文件中配置：
```yaml
DataSources:
  DataView:
    Driver: mysql
    Source: root:password@tcp(127.0.0.1:3306)/idrm_data_view
  DataUnderstanding:
    Driver: mysql
    Source: root:password@tcp(127.0.0.1:3306)/idrm_data_understanding
  ResourceCatalog:
    Driver: mysql
    Source: root:password@tcp(127.0.0.1:3306)/idrm_resource_catalog
```

### Kafka配置
```yaml
Kafka:
  Brokers:
    - 127.0.0.1:9092
  Group: idrm-consumer-group
  Topics:
    - data_view_events
    - catalog_events
```

## 📦 构建和部署

### 本地构建
```bash
make build
```

### Docker构建
```bash
make docker-build
```

### 部署到Kubernetes
```bash
kubectl apply -f deploy/k8s/
```

## 🧪 测试

```bash
# 运行单元测试
make test

# 运行集成测试
make test-integration
```

## 📖 API文档

API文档位于 `docs/api/` 目录，启动服务后访问：
- Swagger UI: http://localhost:8080/swagger

## 🤝 贡献指南

1. Fork 项目
2. 创建特性分支
3. 提交变更
4. 推送到分支
5. 创建 Pull Request

## 📄 许可证

MIT License
