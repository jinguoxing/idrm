# API Types 命名规范

## 问题说明

在 go-zero 中，**所有 API 定义的类型都会生成到一个单一的 `internal/types/types.go` 文件中**。

当多个模块使用相同的类型名称时，会产生**重复定义冲突**。

## 解决方案

### ✅ 方案一：使用模块前缀（推荐）

在不同模块的 API 定义中，使用模块名作为类型前缀。

#### 示例：

**资源目录模块** (`resource_catalog/category.api`):
```api
type (
    CategoryReq {          // 不需要前缀（主模块）
        Id int64 `path:"id"`
    }
    
    CategoryResp {
        Id   int64  `json:"id"`
        Name string `json:"name"`
    }
)
```

**数据视图模块** (`data_view/category.api`):
```api
type (
    DataViewCategoryReq {  // 添加 DataView 前缀
        Id int64 `path:"id"`
    }
    
    DataViewCategoryResp {
        Id   int64  `json:"id"`
        Name string `json:"name"`
    }
)
```

#### 优点：
- ✅ 完全避免命名冲突
- ✅ 类型名称清晰表明所属模块
- ✅ 易于维护和理解

### ⚠️ 方案二：使用业务含义区分

如果同名类型在不同模块中代表不同的业务概念，使用业务名称：

```api
// resource_catalog 模块
type (
    CategoryReq {          // 目录类别
        Id int64 `path:"id"`
    }
)

// data_view 模块
type (
    ViewDefinitionReq {    // 视图定义（而不是叫 Category）
        Id int64 `path:"id"`
    }
)
```

### ❌ 不推荐的方案

#### 1. 手动拆分 types.go
**不可行**：goctl 会覆盖整个 types.go 文件

#### 2. 使用相同的类型名
**会报错**：Go 不允许重复定义

## 命名规范建议

### 模块前缀格式

```
{ModuleName}{TypeName}
```

示例：
- `resource_catalog` → `CategoryReq`（主模块可以不加前缀）
- `data_view` → `DataViewCategoryReq`
- `data_understanding` → `DataUnderstandingQueryReq`

### 常见类型后缀

- **Req**: 请求参数
- **Resp**: 响应结果
- **Info**: 信息对象
- **Item**: 列表项
- **Query**: 查询条件

### 完整示例

```api
// data_view/query.api
type (
    // 请求
    DataViewQueryReq {
        Id int64 `path:"id"`
    }
    
    DataViewExecuteQueryReq {
        Sql string `json:"sql"`
    }
    
    // 响应
    DataViewQueryResp {
        Id     int64  `json:"id"`
        Name   string `json:"name"`
        Status int    `json:"status"`
    }
    
    DataViewQueryResult {
        Columns []string        `json:"columns"`
        Rows    [][]interface{} `json:"rows"`
    }
)
```

## 最佳实践

### 1. 规划类型命名

在创建新模块前，先规划好类型命名规范：

```
resource_catalog/
  - CategoryReq, CategoryResp
  - DirectoryReq, DirectoryResp

data_view/
  - DataViewQueryReq, DataViewQueryResp
  - DataViewSchemaReq, DataViewSchemaResp

data_understanding/
  - DataUnderstandingAnalysisReq
  - DataUnderstandingReportResp
```

### 2. 保持一致性

同一模块内的所有类型使用统一前缀：

✅ **好的命名**：
```
DataViewQueryReq
DataViewQueryResp
DataViewExecuteReq
```

❌ **不好的命名**（不一致）：
```
DataViewQueryReq
QueryResp          // 缺少前缀
DViewExecuteReq    // 缩写不一致
```

### 3. 共享类型的处理

如果多个模块需要使用相同的类型定义，考虑：

#### 选项 A：创建独立的 common.api

```
doc/common/types.api
```

```api
syntax = "v1"

type (
    // 分页请求（所有模块共享）
    PageReq {
        Page     int `form:"page,optional,default=1"`
        PageSize int `form:"page_size,optional,default=10"`
    }
    
    // 通用响应
    BaseResp {
        Code int    `json:"code"`
        Msg  string `json:"msg"`
    }
)
```

然后在主文件中导入：
```api
import "common/types.api"
import "resource_catalog/category.api"
```

#### 选项 B：在各自模块中定义

```api
// 每个模块定义自己的分页类型
type (
    DataViewPageReq {
        Page     int `form:"page,optional,default=1"`
        PageSize int `form:"page_size,optional,default=10"`
    }
)
```

## 检查冲突的方法

### 1. 生成前检查

```bash
cd api
goctl api go -api doc/api.api -dir . 2>&1 | grep -i "duplicate"
```

### 2. 查看 types.go

生成后检查是否有重复定义：

```bash
grep "^type " internal/types/types.go | sort | uniq -d
```

### 3. 编译验证

```bash
go build ./...
```

## 总结

| 方案 | 优点 | 缺点 | 推荐度 |
|------|------|------|--------|
| **模块前缀** | 无冲突，清晰 | 类型名较长 | ⭐⭐⭐⭐⭐ |
| **业务区分** | 语义清晰 | 需仔细规划 | ⭐⭐⭐⭐ |
| **共享 common** | 复用性好 | 耦合度高 | ⭐⭐⭐ |

**推荐使用模块前缀方案**，既能避免冲突，又能保持代码清晰。
