# IDRM - Intelligent Data Resource Management

基于 Go-Zero 框架的企业级数据资源管理平台，提供数据目录、数据视图和数据理解等核心功能。

![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?logo=docker)

---

## � 目录

- [特性](#-特性)
- [架构设计](#️-架构设计)
- [项目结构](#-项目结构)
- [快速开始](#-快速开始)
- [开发指南](#-开发指南)
- [部署](#-部署)
- [贡献](#-贡献)

---

## ✨ 特性

### 核心功能
- 🗂️ **资源目录管理** - 数据资源分类、组织和检索
- 👁️ **数据视图** - 灵活的数据查询和展示
- 🧠 **数据理解** - 智能数据分析和洞察

### 技术特性
- ⚡ **Dual ORM**  - GORM优先，SQLx降级，灵活切换
- 📊 **完整可观测性** - 日志、链路追踪、审计三位一体
- 🔧 **中间件栈** - RequestID、Trace、CORS、Logger、Recovery
- 🐳 **容器化部署** - Docker Compose 一键启动全栈服务
- 🔍 **ELK日志系统** - Elasticsearch + Kibana + Filebeat
- 📈 **分布式追踪** - OpenTelemetry + Jaeger

---

## 🏗️ 架构设计

### 分层架构

```
┌───────────────────────────────────────────────────────┐
│             Application Services                      │
│  ┌───────────┐  ┌───────────┐  ┌───────────┐        │
│  │API Server │  │Job Server │  │Consumer   │        │
│  │  (8888)   │  │  (Cron)   │  │  (Kafka)  │        │
│  └───────────┘  └───────────┘  └───────────┘        │
├───────────────────────────────────────────────────────┤
│          API Layer (HTTP)                             │
│  ┌──────────┐  ┌──────────┐                          │
│  │ Handler  │→ │  Logic   │                          │
│  └──────────┘  └──────────┘                          │
├───────────────────────────────────────────────────────┤
│          Model Layer (Dual ORM)                       │
│  ┌──────────┐  ┌──────────┐                          │
│  │  GORM    │  │  SQLx    │ (Fallback)              │
│  └──────────┘  └──────────┘                          │
├───────────────────────────────────────────────────────┤
│           Infrastructure Layer                        │
│  MySQL │ Redis │ Kafka │ Jaeger │ ELK Stack         │
└───────────────────────────────────────────────────────┘
```

### 技术栈

| 类别 | 技术选型 |
|------|---------|
| **Web 框架** | go-zero 1.6+ |
| **ORM** | GORM 1.31+ / SQLx |
| **数据库** | MySQL 8.0 |
| **缓存** | Redis 7.0 |
| **链路追踪** | OpenTelemetry + Jaeger |
| **日志系统** | ELK Stack 8.11 |
| **消息队列** | Kafka (可选) |
| **容器化** | Docker + Docker Compose |

---

## 📁 项目结构

```
idrm/
├── api/                          # API 服务
│   ├── doc/                      # API 定义文件 (.api)
│   │   ├── api.api              # 主入口
│   │   ├── resource_catalog/    # 资源目录 API
│   │   └── data_view/           # 数据视图 API
│   ├── etc/                      # 配置文件
│   │   ├── api.yaml             # 本地配置
│   │   └── api.docker.yaml      # Docker配置
│   ├── internal/
│   │   ├── config/              # 配置结构
│   │   ├── handler/             # HTTP 处理器
│   │   │   ├── routes.go        # 路由注册
│   │   │   ├── resource_catalog/
│   │   │   └── data_view/
│   │   ├── logic/               # 业务逻辑
│   │   ├── svc/                 # 服务上下文
│   │   └── types/               # 类型定义
│   └── api.go                    # 服务入口
├── job/                          # 定时任务服务
│   ├── etc/                      # 任务配置
│   ├── internal/                 # 任务实现
│   │   ├── handler/             # 任务处理器
│   │   └── scheduler/           # 调度器
│   └── job.go                    # 服务入口
├── consumer/                     # 消息队列消费服务
│   ├── etc/                      # 消费者配置
│   ├── internal/                 # 消费者实现
│   │   ├── handler/             # 消息处理器
│   │   └── listener/            # 监听器
│   └── consumer.go               # 服务入口
├── rpc/                          # RPC 服务（可选）
│   └── resource_catalog/        # 资源目录 RPC
├── model/                        # 数据模型层
│   ├── resource_catalog/        # 资源目录模型
│   │   ├── interface.go         # 接口定义
│   │   ├── factory.go           # ORM工厂
│   │   ├── gorm/                # GORM实现
│   │   ├── sqlx/                # SQLx实现
│   │   └── types.go             # 数据类型
│   ├── data_view/               # 数据视图模型
│   └── data_understanding/      # 数据理解模型
├── pkg/                          # 公共包
│   ├── config/                  # 配置定义
│   ├── db/                      # 数据库工具
│   ├── middleware/              # 中间件
│   │   ├── recovery.go          # Panic恢复
│   │   ├── requestid.go         # 请求ID
│   │   ├── trace.go             # 链路追踪
│   │   ├── cors.go              # 跨域
│   │   └── logger.go            # 日志
│   ├── response/                # 响应格式
│   ├── telemetry/               # 可观测性
│   │   ├── log/                 # 日志系统
│   │   ├── trace/               # 链路追踪
│   │   └── audit/               # 审计日志
│   ├── validator/               # 请求验证
│   └── utils/                   # 工具函数
├── deploy/                       # 部署配置
│   ├── docker-compose.yml       # Docker编排
│   ├── Dockerfile               # 容器镜像
│   ├── README.md                # 部署文档
│   └── config/                  # 部署配置
│       ├── mysql/init.sql       # 数据库初始化
│       ├── filebeat/            # 日志收集
│       └── api/                 # API配置
├── docs/                         # 文档
│   └── architecture/            # 架构文档
├── migrations/                   # 数据库迁移
├── scripts/                      # 脚本工具
├── logs/                         # 日志目录
├── go.mod                        # Go模块
├── Makefile                      # 构建脚本
└── README.md                     # 本文档
```

---

## 🚀 快速开始

### 前置要求

- **Go** 1.24+
- **Docker** 20.10+
- **Docker Compose** 2.0+
- **Make** (可选，用于快捷命令)

### 方式一：Docker Compose（推荐）

一键启动所有服务（MySQL、Redis、Jaeger、ELK、API）：

```bash
# 1. 克隆项目
git clone <repository-url>
cd idrm

# 2. 启动所有服务
cd deploy
docker-compose up -d

# 3. 查看服务状态
docker-compose ps

# 4. 查看API日志
docker-compose logs -f idrm-api
```

**访问地址**：
- API服务: http://localhost:8888
- Kibana日志: http://localhost:5601
- Jaeger追踪: http://localhost:16686

详细部署文档：[deploy/README.md](deploy/README.md)

### 方式二：本地开发

```bash
# 1. 安装依赖
go mod download

# 2. 配置数据库
mysql -u root -p < deploy/config/mysql/init.sql

# 3. 修改配置
cp api/etc/api.yaml api/etc/api-local.yaml
# 编辑 api-local.yaml 修改数据库连接信息

# 4. 启动服务
go run api/api.go -f api/etc/api-local.yaml

# 或使用 Makefile
make run
```

### 服务说明

| 服务 | 端口 | 用途 | 状态 |
|------|------|------|------|
| **api-server** | 8888 | HTTP API服务 | ✅ 运行中 |
| **job-server** | - | 定时任务调度 | 🔧 开发中 |
| **consumer-server** | - | Kafka消息消费 | 🔧 开发中 |
| MySQL | 3306 | 主数据库 | ✅ 运行中 |
| Redis | 6379 | 缓存 | ✅ 运行中 |
| Elasticsearch | 9200 | 日志存储 | ✅ 运行中 |
| Kibana | 5601 | 日志查询 | ✅ 运行中 |
| Jaeger | 16686 | 链路追踪 | ✅ 运行中 |

### 快速测试

#### API服务测试

```bash
# 健康检查
curl http://localhost:8888/health

# 获取分类列表
curl http://localhost:8888/api/v1/catalog/categories

# 创建分类
curl -X POST http://localhost:8888/api/v1/catalog/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"测试分类","description":"这是一个测试"}'
```

#### Job服务测试（开发中）

```bash
# 启动Job服务
go run job/job.go -f job/etc/job.yaml

# 查看定时任务日志
tail -f logs/job.log
```

#### Consumer服务测试（开发中）

```bash
# 启动Consumer服务
go run consumer/consumer.go -f consumer/etc/consumer.yaml

# 发送测试消息到Kafka
kafka-console-producer --topic catalog_events --bootstrap-server localhost:9092
```

---

## 📝 开发指南

### 添加新的API接口

#### 1. 定义API

创建 `api/doc/resource_catalog/new_feature.api`:

```api
syntax = "v1"

type (
    CreateFeatureReq {
        Name string `json:"name" validate:"required"`
    }
    
    FeatureResp {
        Id   int64  `json:"id"`
        Name string `json:"name"`
    }
)

@server(
    group: resource_catalog/feature
    prefix: /api/v1/catalog
)
service idrm-api {
    @doc "创建功能"
    @handler CreateFeature
    post /features (CreateFeatureReq) returns (FeatureResp)
}
```

#### 2. 导入到主文件

编辑 `api/doc/api.api`:

```api
import "resource_catalog/new_feature.api"
```

#### 3. 生成代码

```bash
cd api
goctl api go -api doc/api.api -dir . --style=goZero
```

#### 4. 实现业务逻辑

编辑生成的 `logic/resource_catalog/feature/createfeaturelogic.go`:

```go
func (l *CreateFeatureLogic) CreateFeature(req *types.CreateFeatureReq) (*types.FeatureResp, error) {
    // 实现业务逻辑
    return &types.FeatureResp{
        Id:   1,
        Name: req.Name,
    }, nil
}
```

### 数据模型开发

#### 使用工厂模式

```go
// ServiceContext会自动选择ORM
type ServiceContext struct {
    CategoryModel resource_catalog.CategoryModel  // 接口类型
}

func NewServiceContext(c config.Config) *ServiceContext {
    // 工厂会自动选择GORM或SQLx
    categoryModel := resource_catalog.NewCategoryModel(sqlConn, gormDB)
    
    return &ServiceContext{
        CategoryModel: categoryModel,
    }
}
```

#### Model层使用示例

```go
// 在Logic中使用
func (l *GetCategoryLogic) GetCategory(req *types.CategoryReq) (*types.CategoryResp, error) {
    // 通过接口调用，具体ORM实现透明
    category, err := l.svcCtx.CategoryModel.FindOne(l.ctx, req.Id)
    if err != nil {
        return nil, err
    }
    
    return &types.CategoryResp{
        Id:   category.Id,
        Name: category.Name,
    }, nil
}
```

### 添加定时任务

#### 1. 创建任务处理器

在`job/internal/handler/`创建任务handler:

```go
// job/internal/handler/sync_data_job.go
package handler

import (
    "context"
    "idrm/model/resource_catalog"
    "github.com/zeromicro/go-zero/core/logx"
)

type SyncDataJob struct {
    categoryModel resource_catalog.CategoryModel
}

func NewSyncDataJob(categoryModel resource_catalog.CategoryModel) *SyncDataJob {
    return &SyncDataJob{
        categoryModel: categoryModel,
    }
}

func (j *SyncDataJob) Run(ctx context.Context) error {
    logx.Info("开始执行数据同步任务")
    
    // 实现任务逻辑
    categories, err := j.categoryModel.FindAll(ctx)
    if err != nil {
        return err
    }
    
    logx.Infof("同步了 %d 条数据", len(categories))
    return nil
}
```

#### 2. 在调度器中注册

编辑`job/internal/scheduler/scheduler.go`:

```go
func (s *Scheduler) RegisterJobs() {
    // 注册数据同步任务
    s.cron.AddFunc("0 */1 * * *", func() {  // 每小时执行
        if err := s.syncDataJob.Run(context.Background()); err != nil {
            logx.Errorf("数据同步任务失败: %v", err)
        }
    })
}
```

#### 3. 配置任务

编辑`job/etc/job.yaml`:

```yaml
Jobs:
  SyncData:
    Cron: "0 */1 * * *"  # 每小时执行
    Enabled: true
```

### 添加消息消费者

#### 1. 创建消费者Handler

在`consumer/internal/handler/`创建handler:

```go
// consumer/internal/handler/catalog_consumer.go
package handler

import (
    "context"
    "encoding/json"
    "idrm/model/resource_catalog"
    "github.com/zeromicro/go-zero/core/logx"
)

type CatalogConsumer struct {
    categoryModel resource_catalog.CategoryModel
}

func NewCatalogConsumer(categoryModel resource_catalog.CategoryModel) *CatalogConsumer {
    return &CatalogConsumer{
        categoryModel: categoryModel,
    }
}

func (c *CatalogConsumer) Consume(ctx context.Context, key, val string) error {
    logx.Infof("收到消息: key=%s, val=%s", key, val)
    
    var event struct {
        Action string `json:"action"`
        Data   interface{} `json:"data"`
    }
    
    if err := json.Unmarshal([]byte(val), &event); err != nil {
        return err
    }
    
    // 处理消息逻辑
    switch event.Action {
    case "create":
        // 处理创建事件
    case "update":
        // 处理更新事件
    }
    
    return nil
}
```

#### 2. 配置消费者

编辑`consumer/etc/consumer.yaml`:

```yaml
Kafka:
  Brokers:
    - localhost:9092
  Consumers:
    CatalogConsumer:
      Group: idrm-catalog-group
      Topics:
        - catalog_events
      Workers: 3
      AutoCommit: true
```

### 中间件使用

中间件已在`api/api.go`中全局注册：

```go
// 已注册的中间件（按顺序）
server.Use(middleware.Recovery())   // 1. Panic恢复
server.Use(middleware.RequestID())  // 2. 请求ID
server.Use(middleware.Trace())      // 3. 链路追踪
server.Use(middleware.CORS())       // 4. 跨域
server.Use(middleware.Logger())     // 5. 请求日志
```

#### 在Logic中使用Trace

```go
import "idrm/pkg/telemetry/trace"

func (l *Logic) Handle(req *Req) error {
    // 创建子Span
    ctx, span := trace.StartInternal(l.ctx)
    defer span.End()
    
    // 添加属性
    span.SetAttributes(
        attribute.String("operation", "create"),
    )
    
    // 业务逻辑...
    
    return nil
}
```

#### 获取RequestID

```go
import "idrm/pkg/middleware"

requestID := middleware.GetRequestID(ctx)
logx.Infof("Request ID: %s", requestID)
```

---

## 🧪 测试

```bash
# 单元测试
make test

# 集成测试
make test-integration

# 测试覆盖率
make test-coverage

# 性能测试
make benchmark
```

---

## 🐳 部署

### Docker 部署

详见 [deploy/README.md](deploy/README.md)

```bash
# 构建镜像
docker-compose build

# 启动服务
docker-compose up -d

# 停止服务
docker-compose down

# 查看日志
docker-compose logs -f idrm-api
```

### Kubernetes 部署

```bash
# 应用配置
kubectl apply -f deploy/k8s/

# 查看状态
kubectl get pods -n idrm

# 查看日志
kubectl logs -f <pod-name> -n idrm
```

---

## 📊 监控和日志

### 日志查询（Kibana）

1. 访问 http://localhost:5601
2. 创建Index Pattern: `idrm-api-*`
3. 选择时间字段: `@timestamp`
4. 开始查询日志

### 链路追踪（Jaeger）

1. 访问 http://localhost:16686
2. 选择服务: `idrm-api`
3. 查看Trace详情

### 日志文件

本地日志位于 `logs/` 目录：
- `access.log` - 访问日志
- `error.log` - 错误日志
- `slow.log` - 慢查询日志
- `stat.log` - 统计日志

---

## 🔧 配置说明

### 数据库配置

```yaml
DB:
  ResourceCatalog:
    Host: mysql                 # Docker: mysql, 本地: 127.0.0.1
    Port: 3306
    Database: idrm_resource_catalog
    Username: root
    Password: idrm@2024
    MaxIdleConns: 10
    MaxOpenConns: 100
    ConnMaxLifetime: 3600
```

### Telemetry配置

```yaml
Telemetry:
  ServiceName: idrm-api
  Environment: production
  
  Log:
    Level: info               # debug, info, warn, error
    Mode: file                # file, console
    
  Trace:
    Enabled: true
    Endpoint: jaeger:4317     # OTLP gRPC endpoint
    Sampler: 0.1              # 采样率 (0.0-1.0)
    
  Audit:
    Enabled: true
```

---

## � API文档

### Swagger文档

启动服务后访问：
- Swagger UI: http://localhost:8888/swagger (开发中)

### API示例

#### 创建分类

```bash
curl -X POST http://localhost:8888/api/v1/catalog/categories \
  -H "Content-Type: application/json" \
  -H "X-Request-ID: test-001" \
  -d '{
    "name": "数据库",
    "description": "数据库相关资源"
  }'
```

#### 获取分类列表

```bash
curl http://localhost:8888/api/v1/catalog/categories
```

#### 获取分类详情

```bash
curl http://localhost:8888/api/v1/catalog/categories/1
```

---

## 🛠️ 常用命令

```bash
# 开发
make run                # 启动服务
make build              # 编译
make test               # 测试

# Docker
make docker-build       # 构建镜像
make docker-up          # 启动容器
make docker-down        # 停止容器
make docker-logs        # 查看日志

# 代码质量
make lint               # 代码检查
make fmt                # 代码格式化

# 生成
make gen-api            # 生成API代码
make gen-model          # 生成Model代码
```

---

## 🤝 贡献

我们欢迎所有形式的贡献！

### 贡献流程

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

### 开发规范

- 遵循 Go 代码规范
- 编写单元测试
- 更新相关文档
- Commit message 使用规范格式

---

## 📄 许可证

本项目采用 MIT 许可证。详见 [LICENSE](LICENSE.md) 文件。

---

## 📞 联系方式

- 项目地址: [GitHub Repository]
- 问题反馈: [Issues]
- 文档: [Wiki]

---

## 🙏 致谢

感谢以下开源项目：

- [go-zero](https://github.com/zeromicro/go-zero) - 微服务框架
- [GORM](https://gorm.io/) - ORM库
- [OpenTelemetry](https://opentelemetry.io/) - 可观测性
- [Jaeger](https://www.jaegertracing.io/) - 分布式追踪
- [Elastic Stack](https://www.elastic.co/) - 日志系统

---

**Happy Coding! 🚀**
