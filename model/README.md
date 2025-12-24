# Model 层说明

## 📋 概述

Model 层采用**双 ORM 支持**架构，同时支持 `go-zero sqlx` 和 `gorm` 两种 ORM 实现。

## 🎯 设计理念

- **接口抽象**：统一的模型接口，业务层不感知底层 ORM
- **自动降级**：优先使用 gorm，不可用时自动降级到 sqlx
- **事务支持**：两种 ORM 都支持事务操作
- **无缝切换**：通过工厂方法自动选择，无需配置

## 📁 目录结构

### 新结构（按表分目录）

每个表都是独立的目录，包含该表的所有相关代码：

```
model/
├── README.md                          # 本文档
├── resource_catalog/                  # 资源目录模块
│   └── category/                      # 类别表（独立目录）
│       ├── interface.go               # Model接口定义
│       ├── types.go                   # Category数据结构
│       ├── vars.go                    # 常量和错误定义
│       ├── factory.go                 # ORM工厂（自动选择）
│       ├── gorm_dao.go                # GORM实现
│       └── sqlx_model.go              # SQLx实现
├── data_view/                         # 数据视图模块
│   └── query/                         # 查询表（同上结构）
│       ├── interface.go
│       ├── types.go
│       ├── factory.go
│       ├── gorm_dao.go
│       └── sqlx_model.go
└── data_understanding/                # 数据理解模块
    └── ...
```

### 优势

- ✅ **清晰的职责分离**：每个表是独立单元
- ✅ **易于定位**：按表名快速查找相关代码
- ✅ **独立扩展**：新增表不影响现有代码
- ✅ **符合Go惯例**：按功能模块组织

## 🚀 使用方法

### 1. 初始化模型

```go
// api/internal/svc/servicecontext.go
import (
    "idrm/model/resource_catalog/category"
    _ "idrm/model/resource_catalog/category" // 触发工厂注册
)

type ServiceContext struct {
    Config        config.Config
    CategoryModel category.Model  // 使用接口类型
}

func NewServiceContext(c config.Config) *ServiceContext {
    // 初始化数据库连接
    sqlConn, _ := sqlx.NewMysql(dsn).RawDB()
    gormDB, _ := db.InitGorm(c.DB.ResourceCatalog)
    
    return &ServiceContext{
        Config: c,
        // 自动选择ORM（优先gorm）
        CategoryModel: category.NewModel(sqlConn, gormDB),
    }
}
```

### 2. 基础CRUD操作

```go
// 查询
category, err := svcCtx.CategoryModel.FindOne(ctx, id)

// 插入
newCategory := &category.Category{
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
err := svcCtx.CategoryModel.Trans(ctx, func(ctx context.Context, model category.Model) error {
    // 在事务中执行多个操作
    category1, err := model.Insert(ctx, &category.Category{...})
    if err != nil {
        return err // 自动回滚
    }
    
    category2, err := model.Insert(ctx, &category.Category{...})
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

### 新结构（推荐）：按表分目录

当需要添加新表（如`directory`表）时：

#### 1. 创建表目录

```bash
mkdir -p model/resource_catalog/directory
```

#### 2. 创建核心文件

在`model/resource_catalog/directory/`目录下创建以下文件：

**interface.go** - 定义接口

```go
package directory

import "context"

type Model interface {
    Insert(ctx context.Context, data *Directory) (*Directory, error)
    FindOne(ctx context.Context, id int64) (*Directory, error)
    Update(ctx context.Context, data *Directory) error
    Delete(ctx context.Context, id int64) error
    // ... 其他方法
    
    // 事务支持
    WithTx(tx interface{}) Model
    Trans(ctx context.Context, fn func(ctx context.Context, model Model) error) error
}
```

**types.go** - 数据结构

```go
package directory

import "time"

type Directory struct {
    Id        int64     `db:"id" gorm:"column:id;primaryKey"`
    Name      string    `db:"name" gorm:"column:name;type:varchar(100)"`
    // ... 其他字段
    CreatedAt time.Time `db:"created_at" gorm:"column:created_at;autoCreateTime"`
    UpdatedAt time.Time `db:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}

func (Directory) TableName() string {
    return "directories"
}
```

**vars.go** - 常量和错误

```go
package directory

import "errors"

var (
    ErrNotFound = errors.New("directory not found")
    // ... 其他错误
)

const (
    StatusEnabled  = 1
    StatusDisabled = 0
)
```

**factory.go** - ORM工厂

```go
package directory

import (
    "database/sql"
    "github.com/zeromicro/go-zero/core/logx"
    "gorm.io/gorm"
)

type Factory func(interface{}) Model

var (
    gormFactory Factory
    sqlxFactory Factory
)

func RegisterGormFactory(factory Factory) {
    gormFactory = factory
}

func RegisterSqlxFactory(factory Factory) {
    sqlxFactory = factory
}

func NewModel(sqlConn *sql.DB, gormDB *gorm.DB) Model {
    if gormDB != nil && gormFactory != nil {
        logx.Info("Using GORM for DirectoryModel")
        return gormFactory(gormDB)
    }
    
    if sqlConn != nil && sqlxFactory != nil {
        logx.Info("Using SQLx for DirectoryModel (fallback)")
        return sqlxFactory(sqlConn)
    }
    
    panic("no database connection available for DirectoryModel")
}
```

#### 3. 实现 GORM（gorm_dao.go）

```go
package directory

import (
    "context"
    "gorm.io/gorm"
)

var _ Model = (*DirectoryDao)(nil)

type DirectoryDao struct {
    db *gorm.DB
}

func NewDirectoryDao(db *gorm.DB) Model {
    return &DirectoryDao{db: db}
}

// 实现Model接口的所有方法...

// init 注册gorm工厂
func init() {
    RegisterGormFactory(func(db interface{}) Model {
        if gormDB, ok := db.(*gorm.DB); ok {
            return NewDirectoryDao(gormDB)
        }
        panic("invalid database type for gorm factory")
    })
}
```

#### 4. 实现 SQLx（sqlx_model.go）

```go
package directory

import (
    "context"
    "database/sql"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ Model = (*DirectoryModel)(nil)

type DirectoryModel struct {
    conn sqlx.SqlConn
}

func NewDirectoryModel(conn *sql.DB) Model {
    return &DirectoryModel{
        conn: sqlx.NewSqlConnFromDB(conn),
    }
}

// 实现Model接口的所有方法...

// init 注册sqlx工厂
func init() {
    RegisterSqlxFactory(func(db interface{}) Model {
        if sqlDB, ok := db.(*sql.DB); ok {
            return NewDirectoryModel(sqlDB)
        }
        panic("invalid database type for sqlx factory")
    })
}
```

#### 5. 在ServiceContext中使用

```go
import (
    "idrm/model/resource_catalog/directory"
    _ "idrm/model/resource_catalog/directory" // 触发注册
)

type ServiceContext struct {
    DirectoryModel directory.Model
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        DirectoryModel: directory.NewModel(sqlConn, gormDB),
    }
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
- 每个表的types.go独立在表目录下

```go
// model/resource_catalog/category/types.go
package category

type Category struct {
    Id   int64  `db:"id" gorm:"column:id;primaryKey"`
    Name string `db:"name" gorm:"column:name"`
}

func (Category) TableName() string {
    return "categories"
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
import "idrm/model/resource_catalog/category"

type ServiceContext struct {
    CategoryModel category.Model  // 使用表目录的接口
}

// ❌ 错误：直接使用实现
type ServiceContext struct {
    CategoryModel *gorm.CategoryDao
}
```

### 2. 事务中避免嵌套查询

```go
// ✅ 正确
err := model.Trans(ctx, func(ctx context.Context, txModel category.Model) error {
    // 使用 txModel 操作
    return txModel.Insert(ctx, data)
})

// ❌ 错误
err := model.Trans(ctx, func(ctx context.Context, txModel category.Model) error {
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
