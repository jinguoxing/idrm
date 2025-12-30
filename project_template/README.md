# {{PROJECT_NAME}}

基于 IDRM 规范的 Go-Zero 微服务项目。

## 快速开始

### 1. 初始化项目

```bash
# 交互式模式
./scripts/init-project.sh -i

# 或直接指定参数
./scripts/init-project.sh -n my-app -m github.com/your-org/my-app

# 包含所有可选服务
./scripts/init-project.sh -n my-app -a
```

### 2. 配置环境

编辑 `api/etc/api.yaml`，配置数据库连接等信息。

### 3. 安装依赖

```bash
go mod tidy
```

### 4. 运行服务

```bash
# 运行 API 服务
make run

# 或直接运行
go run api/api.go
```

## 项目结构

```
{{PROJECT_NAME}}/
├── api/                    # API 服务
│   ├── doc/               # API 定义 (.api 文件)
│   ├── etc/               # 配置文件
│   └── internal/          # 内部代码
│       ├── config/        # 配置结构
│       ├── handler/       # Handler 层
│       ├── logic/         # Logic 层
│       ├── svc/           # 服务上下文
│       └── types/         # 类型定义
├── model/                  # Model 层
├── pkg/                    # 公共包
├── rpc/                    # RPC 服务 (可选)
├── job/                    # 定时任务 (可选)
├── consumer/               # 消息消费者 (可选)
├── migrations/             # 数据库迁移
├── specs/                  # 功能规格文档
└── sdd_doc/spec/          # 项目规范
```

## 开发规范

本项目遵循 IDRM 开发规范，详见：

- [CLAUDE.md](CLAUDE.md) - AI 协作配置
- [sdd_doc/spec/](sdd_doc/spec/) - 完整规范文档

### 核心原则

1. **5 阶段工作流**：Context → Specify → Design → Tasks → Implement
2. **分层架构**：Handler → Logic → Model
3. **双 ORM 模式**：GORM + SQLx
4. **代码规范**：函数 < 50 行，测试覆盖 > 80%

## 常用命令

```bash
make build      # 编译项目
make test       # 运行测试
make lint       # 代码检查
make run        # 运行服务
make gen-api    # 生成 API 代码
make help       # 查看所有命令
```

## 技术栈

- **语言**: Go 1.21+
- **框架**: Go-Zero v1.9+
- **数据库**: MySQL 8.0
- **缓存**: Redis 7.0

## 许可证

MIT License
