# API设计指南

> **文档版本**: v2.0  
> **最后更新**: 2025-12-30  
> **状态**: ✅ 已完善

---

## API 服务入口规范

### 初始化顺序

API 服务入口文件 (`api.go`) 必须按以下顺序初始化：

```go
func main() {
    // 1. 加载配置
    var c config.Config
    conf.MustLoad(*configFile, &c)
    
    // 2. 初始化遥测（日志、链路追踪、审计）
    telemetry.Init(c.Telemetry)
    
    // 3. 初始化验证器
    validator.Init()
    
    // 4. 创建服务器
    server := rest.MustNewServer(c.RestConf)
    
    // 5. 注册中间件（按顺序）
    server.Use(middleware.Recovery())   // 1. Panic 恢复
    server.Use(middleware.RequestID())  // 2. 请求ID
    server.Use(middleware.Trace())      // 3. 链路追踪
    server.Use(middleware.CORS())       // 4. 跨域
    server.Use(middleware.Logger())     // 5. 日志
    
    // 6. 初始化服务上下文
    ctx := svc.NewServiceContext(c)
    
    // 7. 注册路由
    handler.RegisterHandlers(server, ctx)
    
    // 8. 启动服务
    server.Start()
}
```

---

## 配置文件规范

### 环境变量支持

敏感配置必须支持环境变量，使用 `${VAR:default}` 语法：

| 配置项 | 环境变量 | 默认值 |
|--------|----------|--------|
| `DB.*.Host` | `DB_HOST` | `127.0.0.1` |
| `DB.*.Port` | `DB_PORT` | `3306` |
| `DB.*.Username` | `DB_USERNAME` | `root` |
| `DB.*.Password` | `DB_PASSWORD` | 无 |
| `Auth.AccessSecret` | `AUTH_ACCESS_SECRET` | 无 |

### api.yaml 模板

```yaml
Name: idrm-api
Host: 0.0.0.0
Port: ${API_PORT:8888}

DB:
  ResourceCatalog:
    Host: ${DB_HOST:127.0.0.1}
    Port: ${DB_PORT:3306}
    Database: idrm_resource_catalog
    Username: ${DB_USERNAME:root}
    Password: ${DB_PASSWORD}
    Charset: utf8mb4
    MaxIdleConns: 10
    MaxOpenConns: 100
    ConnMaxLifetime: 3600
    LogLevel: ${DB_LOG_LEVEL:warn}

Auth:
  AccessSecret: ${AUTH_ACCESS_SECRET}
  AccessExpire: 7200
```

---

## base.api 通用类型

所有模块必须引用 `base.api` 中的通用类型：

**文件位置**: `api/doc/base.api`

```api
import "base.api"
```

### 分页类型

| 类型 | 用途 |
|------|------|
| `PageBaseInfo` | 基础分页（offset + limit） |
| `PageInfo` | 完整分页（+ direction + sort） |
| `KeywordInfo` | 关键字搜索 |
| `PageInfoWithKeyword` | 分页 + 关键字 |

### 通用请求

| 类型 | 用途 |
|------|------|
| `IdReq` | 单个ID路径参数 |
| `IdsReq` | 批量ID请求 |

### 通用响应

| 类型 | 用途 |
|------|------|
| `PageResp` | 分页响应 |
| `EmptyResp` | 空响应 |

---

## RESTful 原则

### 资源导向

```
GET    /api/v1/catalog/categories        # 列表
POST   /api/v1/catalog/categories        # 创建
GET    /api/v1/catalog/categories/:id    # 详情
PUT    /api/v1/catalog/categories/:id    # 更新
DELETE /api/v1/catalog/categories/:id    # 删除
```

### HTTP动词

| 动词 | 用途 |
|------|------|
| GET | 查询 |
| POST | 创建 |
| PUT | 更新（全量） |
| PATCH | 更新（部分） |
| DELETE | 删除 |

---

## 统一响应格式

### 成功响应

```json
{
  "code": 0,
  "msg": "success",
  "data": {}
}
```

### 分页响应

```json
{
  "code": 0,
  "data": {
    "entries": [],
    "total_count": 100
  }
}
```

### 错误响应

```json
{
  "code": "idrm.common.validation_error",
  "description": "参数验证失败",
  "solution": "请检查请求参数",
  "cause": "name字段不能为空"
}
```

---

## 错误码设计

| 错误码 | 说明 |
|--------|------|
| `0` | 成功 |
| `idrm.common.validation_error` | 参数验证错误 |
| `idrm.common.not_found` | 资源不存在 |
| `idrm.common.unauthorized` | 未授权 |
| `idrm.common.forbidden` | 禁止访问 |
| `idrm.common.internal_error` | 内部错误 |

---

## 版本管理

使用URL路径: `/api/v1/`, `/api/v2/`

---

**参考**: [分层架构](./layered-architecture.md) | [Constitution](../constitution.md)
