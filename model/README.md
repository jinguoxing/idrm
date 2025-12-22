# Model 层说明

## 📋 概述

Model 层采用**双 ORM 支持**架构，同时支持 `go-zero sqlx` 和 `gorm` 两种 ORM 实现。

## 🎯 设计理念

- **接口抽象**：统一的模型接口，业务层不感知底层 ORM
- **自动降级**：优先使用 gorm，不可用时自动降级到 sqlx
- **事务支持**：两种 ORM 都支持事务操作
- **无缝切换**：通过工厂方法自动选择，无需配置

## 📁 目录结构

```
model/
├── README.md                    # 本文档
├── resource_catalog/           # 资源目录模块
│   ├── interface.go            # 模型接口定义
│   ├── types.go                # 数据结构（sqlx和gorm共用）
│   ├── vars.go                 # 常量定义
│   ├── factory.go              # ORM工厂（自动选择）
│   ├── sqlx/                   # sqlx实现
│   │   └── category_model.go
│   └── gorm/                   # gorm实现
│       └── category_dao.go
└── data_view/                  # 数据视图模块（同上结构）
```

## 🚀 使用方法

### 1. 初始化模型

```go
// api/internal/svc/servicecontext.go
func NewServiceContext(c config.Config) *ServiceContext {
    // 初始化数据库连接
    sqlConn := sqlx.NewMySQL(c.Mysql.DataSource)
    gormDB := initGorm(c.Mysql.DataSource)
    
    return &ServiceContext{
        Config: c,
        // 自动选择ORM（优先gorm）
        CategoryModel: resource_catalog.NewCategoryModel(sqlConn.RawDB(), gormDB),
    }
}
```

### 2. 基础CRUD操作

```go
// 查询
category, err := svcCtx.CategoryModel.FindOne(ctx, id)

// 插入
newCategory := &resource_catalog.Category{
    Name: "测试类别",
    Code: "TEST001",
}
result, err := svcCtx.CategoryModel.Insert(ctx, newCategory)

// 更新
category.Name = "新名称"
err := svcCtx.CategoryModel.Update(ctx, category)

// 删除
err := svcCtx.CategoryModel.Delete(ctx, id)

// 列表查询
categories, total, err := svcCtx.CategoryModel.List(ctx, page, pageSize)
```

### 3. 事务操作

#### 方式一：使用 Trans 方法（推荐）

```go
err := svcCtx.CategoryModel.Trans(ctx, func(ctx context.Context, model resource_catalog.CategoryModel) error {
    // 在事务中执行多个操作
    category1, err := model.Insert(ctx, &resource_catalog.Category{...})
    if err != nil {
        return err // 自动回滚
    }
    
    category2, err := model.Insert(ctx, &resource_catalog.Category{...})
    if err != nil {
        return err // 自动回滚
    }
    
    return nil // 自动提交
})
```

#### 方式二：使用 WithTx 方法

```go
// 获取事务
tx := db.Begin() // gorm事务 或 sqlxConn.Begin() // sqlx事务

// 创建事务模型
txModel := svcCtx.CategoryModel.WithTx(tx)

// 在事务中操作
_, err := txModel.Insert(ctx, category1)
if err != nil {
    tx.Rollback()
    return err
}

_, err = txModel.Insert(ctx, category2)
if err != nil {
    tx.Rollback()
    return err
}

tx.Commit()
```

## 📊 ORM 选择逻辑

```
┌─────────────────────┐
│  NewCategoryModel   │
└──────────┬──────────┘
           │
           ▼
    ┌──────────────┐
    │ gormDB != nil?│
    └──────┬───────┘
           │
     Yes ┌─┴─┐ No
         │   │
         ▼   ▼
    ┌────────┐  ┌──────────────┐
    │  gorm  │  │sqlConn != nil?│
    └────────┘  └──────┬───────┘
                       │
                  Yes ┌─┴─┐ No
                      │   │
                      ▼   ▼
                 ┌────────┐ ┌──────┐
                 │  sqlx  │ │ panic │
                 └────────┘ └───────┘
```

**优先级**：gorm > sqlx

## 🔧 添加新模型

### 1. 创建目录结构

```bash
mkdir -p model/{module}/{sqlx,gorm}
```

### 2. 创建核心文件

```bash
touch model/{module}/interface.go
touch model/{module}/types.go
touch model/{module}/vars.go
touch model/{module}/factory.go
```

### 3. 定义接口

```go
// model/{module}/interface.go
package {module}

type {Model}Model interface {
    Insert(ctx context.Context, data *{Model}) (*{Model}, error)
    FindOne(ctx context.Context, id int64) (*{Model}, error)
    // ... 其他方法
    
    // 事务支持
    WithTx(tx interface{}) {Model}Model
    Trans(ctx context.Context, fn func(ctx context.Context, model {Model}Model) error) error
}
```

### 4. 实现 ORM

- **gorm**: 实现 `model/{module}/gorm/{table}_dao.go`
- **sqlx**: 实现 `model/{module}/sqlx/{table}_model.go`

### 5. 创建工厂

```go
// model/{module}/factory.go
func New{Model}Model(sqlConn *sql.DB, gormDB *gorm.DB) {Model}Model {
    if gormDB != nil {
        return gorm.New{Model}Dao(gormDB)
    }
    if sqlConn != nil {
        return sqlx.New{Model}Model(sqlConn)
    }
    panic("no database connection available")
}
```

## ⚠️ 注意事项

### 1. 接口一致性
- 两种 ORM 实现必须符合相同接口
- 方法签名完全一致
- 返回的数据结构要一致
- 错误处理要统一

### 2. 数据结构
- `types.go` 中的结构体同时支持 `db` 和 `gorm` tag
- 使用 `TableName()` 方法指定表名

```go
type Category struct {
    Id   int64  `db:"id" gorm:"column:id;primaryKey"`
    Name string `db:"name" gorm:"column:name"`
}

func (Category) TableName() string {
    return "category"
}
```

### 3. 事务处理
- 使用 `Trans()` 方法更安全，自动提交/回滚
- 使用 `WithTx()` 更灵活，但需手动管理事务
- 两种方法不要混用

### 4. 性能考虑
- **gorm**: 功能强大，支持复杂查询，但稍慢
- **sqlx**: 性能更好，接近原生 SQL，但功能较少
- 根据场景自动选择

## 🎨 最佳实践

### 1. 统一使用接口

```go
// ✅ 正确：使用接口
type ServiceContext struct {
    CategoryModel resource_catalog.CategoryModel
}

// ❌ 错误：直接使用实现
type ServiceContext struct {
    CategoryModel *gorm.CategoryDao
}
```

### 2. 事务中避免嵌套查询

```go
// ✅ 正确
err := model.Trans(ctx, func(ctx context.Context, txModel CategoryModel) error {
    // 使用 txModel 操作
    return txModel.Insert(ctx, data)
})

// ❌ 错误
err := model.Trans(ctx, func(ctx context.Context, txModel CategoryModel) error {
    // 不要使用原 model
    return model.Insert(ctx, data)
})
```

### 3. 错误处理

```go
category, err := model.FindOne(ctx, id)
if err != nil {
    if err.Error() == "category not found" {
        // 处理不存在的情况
    }
    return err
}
```

## 📚 参考资料

- [go-zero 文档](https://go-zero.dev/)
- [gorm 文档](https://gorm.io/)
- [sqlx 文档](https://github.com/jmoiron/sqlx)

## 🆘 常见问题

**Q: 如何判断当前使用的是哪个 ORM？**

A: 查看启动日志，会输出 "Using GORM" 或 "Using SQLx"。

**Q: 可以强制指定使用某个 ORM 吗？**

A: 可以，在 `NewCategoryModel` 时只传入一个连接即可。

**Q: 如何处理复杂查询？**

A: 建议在各自的实现中添加专用方法，然后在接口中声明可选方法。

**面向对你Q: 两种 ORM 的查询结果一致吗？**

A: 是的，接口保证了返回的数据结构完全一致。
