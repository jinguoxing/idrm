# API 模块化目录结构说明

## 📁 新的目录结构

按照功能模块组织，每个模块下按功能点（feature）再细分：

```
api/
├── doc/                                    # API 定义文件
│   ├── api.api                            # 主入口（导入所有模块）
│   ├── README.md                          # 使用说明
│   └── resource_catalog/                  # 资源目录模块
│       └── category.api                   # 类别功能 API
│
├── internal/
│   ├── handler/
│   │   ├── routes.go                      # 统一路由注册（goctl 自动生成）
│   │   └── resource_catalog/              # 资源目录模块
│   │       └── category/                  # 类别功能
│   │           ├── getcategoryhandler.go
│   │           ├── createcategoryhandler.go
│   │           └── listcategoryhandler.go
│   │
│   └── logic/
│       └── resource_catalog/              # 资源目录模块
│           └── category/                  # 类别功能
│               ├── getcategorylogic.go
│               ├── createcategorylogic.go
│               └── listcategorylogic.go
```

## 🎯 设计原则

### 1. 模块化组织
- 第一层：业务模块（如 `resource_catalog`、`data_view`）
- 第二层：功能点（如 `category`、`directory`）
- 第三层：CRUD 操作文件

### 2. 无模块级 routes.go
- ✅ **只保留顶层的 `handler/routes.go`**（由 goctl 自动生成和维护）
- ❌ **不需要** `handler/resource_catalog/routes.go`
- ❌ **不需要** 各个功能目录下的 routes.go

### 3. API 定义对应
每个功能的 API 定义文件中使用 `group` 指定生成路径：

```api
// doc/resource_catalog/category.api
@server(
    group: resource_catalog/category  // 生成到 resource_catalog/category/ 目录
    prefix: /api/v1/catalog
)
```

## 📝 添加新功能示例

### 示例：添加 "目录 (Directory)" 功能

#### 1. 创建 API 定义文件

`api/doc/resource_catalog/directory.api`：

```api
syntax = "v1"

// 目录类型定义
type (
    DirectoryReq {
        Id int64 `path:"id"`
    }
    
    DirectoryResp {
        Id   int64  `json:"id"`
        Name string `json:"name"`
        // ...其他字段
    }
)

// 资源目录 - 目录服务
@server(
    group: resource_catalog/directory  // 注意这里
    prefix: /api/v1/catalog
)
service Api {
    @doc "获取目录详情"
    @handler GetDirectory
    get /directories/:id (DirectoryReq) returns (DirectoryResp)
}
```

#### 2. 在主文件中导入

`api/doc/api.api`：

```api
import "resource_catalog/category.api"
import "resource_catalog/directory.api"  // 新增
```

#### 3. 生成代码

```bash
cd api
goctl api go -api doc/api.api -dir .
```

#### 4. 生成结果

会自动创建：
```
handler/resource_catalog/directory/
└── getdirectoryhandler.go

logic/resource_catalog/directory/
└── getdirectorylogic.go
```

## 🔄 目录结构示例（完整）

假设有多个模块和功能：

```
api/internal/
├── handler/
│   ├── routes.go                          # 唯一的路由文件
│   ├── resource_catalog/
│   │   ├── category/                      # 类别功能
│   │   │   ├── getcategoryhandler.go
│   │   │   ├── createcategoryhandler.go
│   │   │   └── listcategoryhandler.go
│   │   └── directory/                     # 目录功能
│   │       ├── getdirectoryhandler.go
│   │       └── createdirectoryhandler.go
│   └── data_view/
│       └── query/                         # 查询功能
│           └── executeQueryhandler.go
│
└── logic/
    ├── resource_catalog/
    │   ├── category/
    │   │   ├── getcategorylogic.go
    │   │   ├── createcategorylogic.go
    │   │   └── listcategorylogic.go
    │   └── directory/
    │       ├── getdirectorylogic.go
    │       └── createdirectorylogic.go
    └── data_view/
        └── query/
            └── executequerylogic.go
```

## ✅ 优势

1. **清晰的层级**：模块 → 功能 → 操作
2. **易于扩展**：新增功能只需添加新的功能目录
3. **避免混乱**：不会出现大量文件平铺在同一目录
4. **符合规范**：遵循 go-zero 的 group 机制
5. **统一管理**：所有路由在一个 routes.go 中注册

## ⚠️ 注意事项

1. **group 路径格式**：使用斜杠分隔，如 `resource_catalog/category`
2. **不要手动创建目录**：让 goctl 自动生成
3. **不要手动修改 routes.go**：由 goctl 维护
4. **API 文件组织**：`doc/` 下的目录结构不必和 `internal/` 完全一致
