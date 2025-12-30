# Model Layer

Model 层是数据访问层，负责与数据库交互。采用双 ORM 模式（GORM + SQLx）。

## 目录结构

每个表对应一个子目录：

```
model/{module}/{table}/
├── interface.go    # Model 接口定义
├── types.go        # 数据结构
├── vars.go         # 常量和错误定义
├── factory.go      # ORM 工厂函数
├── gorm_dao.go     # GORM 实现
└── sqlx_model.go   # SQLx 实现（可选）
```

## 接口定义示例

```go
type Model interface {
    Insert(ctx context.Context, data *Entity) (*Entity, error)
    FindOne(ctx context.Context, id int64) (*Entity, error)
    Update(ctx context.Context, data *Entity) error
    Delete(ctx context.Context, id int64) error
    WithTx(tx interface{}) Model
    Trans(ctx context.Context, fn func(ctx context.Context, model Model) error) error
}
```

## ORM 选择策略

- **GORM**：复杂查询、关联加载、事务管理
- **SQLx**：简单查询、性能敏感场景
