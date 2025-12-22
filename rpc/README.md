# RPC 层

gRPC 服务目录。

## 用途

用于存放 gRPC 服务代码，提供内部服务间通信或对外 RPC 接口。

## 结构示例

```
rpc/
├── resource_catalog/          # 资源目录RPC服务
│   ├── proto/                 # protobuf定义
│   ├── internal/              # 内部实现
│   └── resource_catalog.go    # 服务入口
├── data_view/                 # 数据视图RPC服务
└── ...
```

## 注意事项

- 使用 `goctl rpc` 生成代码
- Proto 文件放在 `proto/` 子目录
- 可复用 `model/` 层的数据访问接口
