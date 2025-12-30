# RPC Service

RPC 服务目录，用于定义和实现 gRPC 服务。

## 目录结构

```
rpc/
├── {service_name}/
│   ├── {service_name}.go        # 服务入口
│   ├── etc/
│   │   └── {service_name}.yaml  # 配置文件
│   ├── proto/
│   │   └── {service_name}.proto # Protobuf 定义
│   └── internal/
│       ├── config/
│       ├── logic/
│       ├── server/
│       └── svc/
```

## 创建 RPC 服务

```bash
# 使用 goctl 创建 RPC 服务
goctl rpc new {service_name}

# 从 proto 文件生成代码
goctl rpc protoc {service_name}.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

## 注意事项

- RPC 服务应遵循相同的分层架构
- 所有业务逻辑在 Logic 层实现
- Model 层通过 ServiceContext 注入
