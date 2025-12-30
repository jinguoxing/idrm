# Phase 2 详细操作指南

> Design - 架构设计，代码的蓝图

---

## 🎯 本阶段目标

将Phase 1的需求规范，转化为具体的技术设计方案。

**时间投入**：
- 简单功能：20-30分钟
- 中等功能：40-60分钟
- 复杂功能：1-2小时

**核心产物**：`design.md`

---

## 📋 详细操作步骤

### Step 1: 架构设计（Architecture Overview）

#### 1.1 确定分层架构

**IDRM标准三层**：
```
Handler → Logic → Model
```

**具体到功能**：
```markdown
## Architecture

### 分层结构

\`\`\`
┌─────────────────────────────────────┐
│  Handler层 (HTTP接口)               │
│  - 参数解析和验证                   │
│  - 调用Logic                        │
│  - 响应封装                         │
└─────────────────────────────────────┘
              ↓
┌─────────────────────────────────────┐
│  Logic层 (业务逻辑)                 │
│  - 业务规则实现                     │
│  - 数据转换                         │
│  - 调用Model                        │
└─────────────────────────────────────┘
              ↓
┌─────────────────────────────────────┐
│  Model层 (数据访问)                 │
│  - 接口定义                         │
│  - GORM实现 / SQLx实现              │
│  - Factory                          │
└─────────────────────────────────────┘
\`\`\`
```

#### 1.2 选择ORM策略

**决策树**：
```
功能复杂度？
├─ 简单CRUD → SQLx
├─ 复杂查询 → GORM
└─ 混合场景 → 两者都实现
```

**示例决策**：
```markdown
### ORM选择

**Category CRUD**:
- 创建/查询/更新/删除 → SQLx（性能优先）
- 树形查询（父子关系） → GORM（支持预加载）

**实现策略**: 双ORM都实现，通过Factory切换
```

---

### Step 2: 文件结构规划（File Structure）

#### 2.1 完整文件清单

```markdown
## File Structure

### Model层
\`\`\`
model/resource_catalog/category/
├── interface.go          # Model接口定义
├── types.go              # 数据结构定义
├── factory.go            # ORM工厂
├── gorm_dao.go           # GORM实现
├── sqlx_model.go         # SQLx实现
└── category_test.go      # Model测试
\`\`\`

### Logic层
\`\`\`
api/internal/logic/resource_catalog/category/
├── createcategorylogic.go   # 创建分类逻辑
├── listcategorylogic.go     # 列表查询逻辑
├── getcategorylogic.go      # 详情查询逻辑
├── updatecategorylogic.go   # 更新分类逻辑
├── deletecategorylogic.go   # 删除分类逻辑
└── *_test.go                # Logic测试
\`\`\`

### Handler层
\`\`\`
api/internal/handler/resource_catalog/category/
├── createcategoryhandler.go  # POST /api/v1/categories
├── listcategoryhandler.go    # GET /api/v1/categories
├── getcategoryhandler.go     # GET /api/v1/categories/:id
├── updatecategoryhandler.go  # PUT /api/v1/categories/:id
├── deletecategoryhandler.go  # DELETE /api/v1/categories/:id
├── routes.go                 # 路由注册
└── *_test.go                 # Handler测试
\`\`\`
```

#### 2.2 命名规范

**文件命名**：
- Handler: `{操作}{资源}handler.go` (小写)
- Logic: `{操作}{资源}logic.go` (小写)
- Model: `{orm}_dao.go` 或 `{orm}_model.go`

**示例**：
- ✅ `createcategoryhandler.go`
- ✅ `listcategorylogic.go`
- ✅ `gorm_dao.go`
- ❌ `CategoryHandler.go` (大写)
- ❌ `create_category_handler.go` (下划线)

---

### Step 3: 接口定义（Interface Definitions）

#### 3.1 Model接口

```go
// interface.go
package category

import "context"

// Model 分类数据访问接口
type Model interface {
    // Insert 插入新分类
    Insert(ctx context.Context, data *Category) error
    
    // FindOne 根据ID查询分类
    FindOne(ctx context.Context, id int64) (*Category, error)
    
    // List 分页查询分类列表
    List(ctx context.Context, parentID *int64, page, size int) ([]*Category, int64, error)
    
    // Update 更新分类
    Update(ctx context.Context, data *Category) error
    
    // SoftDelete 软删除分类
    SoftDelete(ctx context.Context, id int64) error
    
    // FindByCode 根据编码查询
    FindByCode(ctx context.Context, code string) (*Category, error)
    
    // ExistsByName 检查名称是否存在
    ExistsByName(ctx context.Context, name string, excludeID int64) (bool, error)
    
    // GetChildren 获取子分类
    GetChildren(ctx context.Context, parentID int64) ([]*Category, error)
}
```

#### 3.2 Logic接口（可选）

```go
// Logic 分类业务逻辑接口
type Logic interface {
    // Create 创建分类
    Create(ctx context.Context, req *CreateCategoryReq) (*Category, error)
    
    // List 查询分类列表
    List(ctx context.Context, req *ListCategoryReq) (*ListCategoryResp, error)
    
    // Get 查询分类详情
    Get(ctx context.Context, id int64) (*Category, error)
    
    // Update 更新分类
    Update(ctx context.Context, id int64, req *UpdateCategoryReq) error
    
    // Delete 删除分类
    Delete(ctx context.Context, id int64) error
}
```

#### 3.3 请求/响应结构

```go
// types.go

// CreateCategoryReq 创建分类请求
type CreateCategoryReq struct {
    Name      string  `json:"name" validate:"required,min=1,max=50"`
    Code      string  `json:"code" validate:"required,min=2,max=20,alphanum_"`
    ParentID  *int64  `json:"parent_id,omitempty"`
    SortOrder int     `json:"sort_order"`
}

// UpdateCategoryReq 更新分类请求
type UpdateCategoryReq struct {
    Name      *string `json:"name,omitempty" validate:"omitempty,min=1,max=50"`
    SortOrder *int    `json:"sort_order,omitempty"`
}

// ListCategoryReq 查询分类列表请求
type ListCategoryReq struct {
    ParentID *int64 `form:"parent_id"`
    Page     int    `form:"page" validate:"min=1"`
    Size     int    `form:"size" validate:"min=1,max=100"`
}

// ListCategoryResp 查询分类列表响应
type ListCategoryResp struct {
    Total int64       `json:"total"`
    Items []*Category `json:"items"`
}
```

---

### Step 4: 序列图设计（Sequence Diagrams）

#### 4.1 创建分类流程

```markdown
### 创建分类序列图

\`\`\`mermaid
sequenceDiagram
    participant C as Client
    participant H as Handler
    participant L as Logic
    participant M as Model
    participant DB as Database
    
    C->>H: POST /api/v1/categories
    activate H
    H->>H: 解析JSON
    H->>H: 参数验证
    
    H->>L: Create(ctx, req)
    activate L
    L->>L: 业务规则检查
    L->>M: ExistsByName(ctx, name)
    M->>DB: SELECT FROM categories
    DB-->>M: result
    M-->>L: false (不存在)
    
    alt 名称已存在
        M-->>L: true
        L-->>H: error: 名称已存在
        H-->>C: 409 Conflict
    else 名称不存在
        L->>M: Insert(ctx, category)
        M->>DB: INSERT INTO categories
        DB-->>M: id
        M-->>L: nil
        L-->>H: category
        deactivate L
        H-->>C: 201 Created
        deactivate H
    end
\`\`\`
```

#### 4.2 查询列表流程

```markdown
### 查询列表序列图

\`\`\`mermaid
sequenceDiagram
    participant C as Client
    participant H as Handler
    participant L as Logic
    participant M as Model
    participant DB as Database
    
    C->>H: GET /api/v1/categories?page=1&size=20
    H->>H: 解析Query参数
    H->>H: 参数验证
    
    H->>L: List(ctx, req)
    L->>M: List(ctx, nil, 1, 20)
    M->>DB: SELECT * FROM categories LIMIT 20
    DB-->>M: rows
    M->>DB: SELECT COUNT(*) FROM categories
    DB-->>M: total
    M-->>L: items, total
    L-->>H: resp
    H-->>C: 200 OK
\`\`\`
```

---

### Step 5: Technical Constraints（技术约束）

> **新增**：从Phase 1移至Phase 2，聚焦技术实现约束

#### 5.1 IDRM通用约束（必须包含）

每个Design都必须声明：

```markdown
## Technical Constraints

### IDRM通用约束
- MUST follow layered architecture (Handler → Logic → Model)
- MUST implement dual ORM support (GORM + SQLx)
- Functions MUST be < 50 lines
- MUST use Chinese comments for all public items
- Error wrapping MUST use %w
- Test coverage MUST be > 80%
- MUST pass golangci-lint check
- MUST follow go-style-guide
```

#### 5.2 功能特定约束

基于具体功能添加技术约束：

**示例：分类管理**
```markdown
### 功能特定约束

#### ORM选择
- 使用GORM实现树形查询（支持预加载）
- 使用SQLx实现简单CRUD（性能优化）
- 通过Factory pattern支持运行时切换

#### 性能要求
- List API响应时间 SHOULD be < 200ms
- MUST支持分页（每页最多100条）
- 树形查询SHOULD使用索引优化

#### 安全约束
- 父分类ID必须验证存在性
- 软删除不能被普通查询返回
- 层级深度必须在数据库层面限制
```

#### 5.3 AI辅助生成

```
@requirements.md (Phase 1的业务需求)
@spec/architecture/layered-architecture.md
@spec/coding-standards/go-style-guide.md

Phase 2: 生成Technical Constraints

要求：
1. 包含所有IDRM通用约束
2. 基于Business Rules确定技术约束
3. ORM选择有明确理由
4. 性能和安全要求量化
```

---

### Step 6: Data Model（数据模型）

> **新增**：从Phase 1移至Phase 2，定义具体的数据库设计

#### 6.1 表结构设计

```markdown
## Data Model

### categories表

\`\`\`sql
CREATE TABLE categories (
    id          BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '主键ID',
    name        VARCHAR(50) NOT NULL COMMENT '分类名称',
    code        VARCHAR(20) NOT NULL UNIQUE COMMENT '分类编码',
    parent_id   BIGINT NULL COMMENT '父分类ID',
    level       TINYINT NOT NULL DEFAULT 1 COMMENT '层级(1-3)',
    sort_order  INT NOT NULL DEFAULT 0 COMMENT '排序值',
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at  DATETIME NULL COMMENT '软删除时间',
    
    INDEX idx_parent_id (parent_id),
    INDEX idx_code (code),
    INDEX idx_deleted_at (deleted_at),
    INDEX idx_level (level),
    
    CONSTRAINT fk_parent FOREIGN KEY (parent_id) REFERENCES categories(id),
    CONSTRAINT chk_level CHECK (level BETWEEN 1 AND 3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源分类表';
\`\`\`
```

#### 6.2 字段详细说明

```markdown
### 字段定义

| 字段 | 类型 | 约束 | 说明 | 业务规则来源 |
|------|------|------|------|-------------|
| id | BIGINT | PK, AUTO_INCREMENT | 主键 | - |
| name | VARCHAR(50) | NOT NULL | 分类名称 | Phase 1: 1-50字符 |
| code | VARCHAR(20) | NOT NULL, UNIQUE | 分类编码 | Phase 1: 全局唯一 |
| parent_id | BIGINT | NULLABLE, FK | 父分类ID | Phase 1: 树形结构 |
| level | TINYINT | 1-3, CHECK | 层级 | Phase 1: 最多3层 |
| sort_order | INT | NOT NULL | 排序 | Phase 1: 自定义排序 |
| created_at | DATETIME | NOT NULL | 创建时间 | - |
| updated_at | DATETIME | NOT NULL | 更新时间 | - |
| deleted_at | DATETIME | NULLABLE | 软删除 | Phase 1: 软删除规则 |

**索引设计**：
- `idx_parent_id`: 支持按父分类查询
- `idx_code`: 支持编码唯一性快速检查
- `idx_deleted_at`: 支持排除已删除记录
- `idx_level`: 支持按层级过滤
```

#### 6.3 Go结构体定义

```go
// types.go

// Category 资源分类实体
type Category struct {
    ID        int64      `gorm:"column:id;primaryKey" db:"id" json:"id"`
    Name      string     `gorm:"column:name;size:50;not null" db:"name" json:"name" validate:"required,min=1,max=50"`
    Code      string     `gorm:"column:code;size:20;not null;uniqueIndex" db:"code" json:"code" validate:"required,min=2,max=20"`
    ParentID  *int64     `gorm:"column:parent_id" db:"parent_id" json:"parent_id,omitempty"`
    Level     int8       `gorm:"column:level;not null;default:1" db:"level" json:"level"`
    SortOrder int        `gorm:"column:sort_order;not null;default:0" db:"sort_order" json:"sort_order"`
    CreatedAt time.Time  `gorm:"column:created_at;not null" db:"created_at" json:"created_at"`
    UpdatedAt time.Time  `gorm:"column:updated_at;not null" db:"updated_at" json:"updated_at"`
    DeletedAt *time.Time `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at,omitempty"`
}

// TableName 指定表名
func (Category) TableName() string {
    return "categories"
}
```

#### 6.4 业务规则到数据库映射

说明Phase 1的Business Rules如何转化为数据库设计：

```markdown
### Business Rules → Database Design

| Business Rule (Phase 1) | Database Implementation (Phase 2) |
|-------------------------|-----------------------------------|
| 分类最多3层 | CHECK约束: level BETWEEN 1 AND 3 |
| 名称唯一（同父级下） | 应用层实现（Model层检查） |
| 编码全局唯一 | UNIQUE约束 on code |
| 删除前检查子分类 | 应用层实现（Logic层检查） |
| 软删除 | deleted_at字段 + 索引 |
| 父子关系 | parent_id外键 + 索引 |
```

---

### Step 7: 错误处理设计

#### 5.1 错误定义

```go
// errors.go
package category

import "errors"

var (
    // ErrCategoryNotFound 分类不存在
    ErrCategoryNotFound = errors.New("category not found")
    
    // ErrCategoryNameExists 分类名称已存在
    ErrCategoryNameExists = errors.New("category name already exists")
    
    // ErrCategoryCodeExists 分类编码已存在
    ErrCategoryCodeExists = errors.New("category code already exists")
    
    // ErrMaxLevelExceeded 超过最大层级
    ErrMaxLevelExceeded = errors.New("max category level (3) exceeded")
    
    // ErrCategoryHasChildren 分类有子分类
    ErrCategoryHasChildren = errors.New("category has children, cannot delete")
)
```

#### 5.2 分层错误处理

```markdown
### 错误处理策略

**Model层**:
- 数据库错误 → 包装返回
- 记录不存在 → 返回ErrCategoryNotFound

**Logic层**:
- 业务规则违反 → 返回业务错误
- Model错误 → 使用%w包装传递

**Handler层**:
- 错误 → HTTP状态码映射
- 统一错误响应格式
```

```go
// Handler层错误映射
func (h *Handler) handleError(c *gin.Context, err error) {
    switch {
    case errors.Is(err, category.ErrCategoryNotFound):
        ginx.ResNotFoundJson(c, err)
    case errors.Is(err, category.ErrCategoryNameExists):
        ginx.ResConflictJson(c, err)
    case errors.Is(err, category.ErrMaxLevelExceeded):
        ginx.ResBadRequestJson(c, err)
    default:
        ginx.ResInternalErrorJson(c, err)
    }
}
```

---

### Step 8: 测试策略

#### 6.1 测试金字塔

```markdown
### 测试分层

\`\`\`
         ┌──────────┐
         │  E2E     │  少量
         │  测试    │
         ├──────────┤
         │  集成    │  适量
         │  测试    │
         ├──────────┤
         │  单元    │  大量
         │  测试    │
         └──────────┘
\`\`\`

**单元测试**:
- Model层：每个方法
- Logic层：每个业务方法
- Handler层：每个接口

**集成测试**:
- Model + 真实DB
- Logic + Mock Model

**E2E测试**:
- 完整API流程
```

#### 6.2 Mock策略

```markdown
### Mock使用

**Model层**: 
- 无需Mock（直接测试DB或用内存DB）

**Logic层**:
- Mock Model接口

**Handler层**:
- Mock Logic接口
```

---

## ✅ Gate 2 质量检查

### 自检清单

**架构设计**：
- [ ] 符合分层架构（Handler → Logic → Model）
- [ ] ORM选择合理（GORM vs SQLx）
- [ ] 依赖方向正确

**文件结构**：
- [ ] 文件清单完整
- [ ] 命名规范统一
- [ ] 目录结构清晰

**接口定义**：
- [ ] Model接口完整
- [ ] 请求/响应结构清晰
- [ ] 参数验证规则明确

**序列图**（在约束之前）：
- [ ] 主要流程都有序列图
- [ ] 步骤清晰
- [ ] 包含异常处理

**Technical Constraints**：
- [ ] 包含IDRM通用约束
- [ ] 功能特定约束明确
- [ ] ORM选择有理由
- [ ] 性能要求量化

**Data Model**：
- [ ] 表结构完整
- [ ] 字段定义清晰
- [ ] 索引设计合理
- [ ] Go结构体正确
- [ ] Business Rules映射说明

**错误处理**：
- [ ] 错误定义完整
- [ ] 分层处理策略清晰
- [ ] HTTP状态码映射合理

**测试策略**：
- [ ] 测试分层明确
- [ ] Mock策略清晰
- [ ] 覆盖率目标明确（>80%）

---

## 💡 实战技巧

### 技巧1：参考现有实现

```bash
# 查找类似功能的design
find specs/ -name "design.md" | xargs grep -l "CRUD"

# 学习文件结构
ls -R model/*/category/
ls -R api/internal/logic/*/category/
```

### 技巧2：AI辅助设计

**生成文件结构**：
```
@requirements.md
@spec/architecture/layered-architecture.md

Phase 2: 请生成Design

要求：
1. 完整的文件清单
2. Model接口定义
3. 主要流程的序列图
```

**Review设计**：
```
@design.md
@spec/architecture/layered-architecture.md

Review这个设计：
1. 是否符合分层架构？
2. 接口设计是否合理？
3. 有没有遗漏的文件？
```

### 技巧3：可视化优先

先画图，再写代码：
1. 架构图 - 整体结构
2. 序列图 - 主要流程
3. 状态图 - 状态变化（如需要）

### 技巧4：接口先行

定义好接口，实现就简单了：
1. Model接口 - 数据访问能力
2. Logic接口 - 业务能力
3. Handler - 遵循接口调用

---

## 🔧 常见问题

### Q1: Design需要多详细？
**A**: 
- 接口定义：必须详细
- 文件清单：必须完整
- 序列图：主流程即可
- 实现细节：不需要

### Q2: 一定要画序列图吗？
**A**: 
- 简单CRUD：可选
- 复杂流程：强烈建议
- 多系统交互：必须

### Q3: ORM如何选择？
**A**:
- 简单CRUD → SQLx
- 复杂查询/关联 → GORM
- 不确定 → 先GORM，性能瓶颈时加SQLx

### Q4: 接口要定义多细？
**A**:
- 方法签名：必须完整
- 参数类型：必须清晰
- 返回值：必须明确
- 实现逻辑：不需要

---

## 📊 Phase 2输出

### design.md结构

```markdown
# Design: {功能名}

## Architecture Overview
分层架构图

## ORM Strategy
GORM vs SQLx选择

## File Structure
完整文件清单

## Interfaces

### Model Interface
Model接口定义

### Logic Interface
Logic接口定义（可选）

### Request/Response Types
请求响应结构

## Sequence Diagrams
主要流程序列图（在约束之前）

## Technical Constraints

### IDRM通用约束
标准技术约束

### 功能特定约束
ORM选择、性能、安全等

## Data Model

### 表结构
SQL DDL定义

### 字段说明
详细字段定义表

### Go结构体
实体类型定义

### Business Rules映射
业务规则到数据库的映射关系

## Error Handling
错误定义和处理策略

## Testing Strategy
测试策略和Mock方案

## Next Steps
→ Phase 3: Tasks
```

---

**Phase 2做好了，实施起来事半功倍！** 🎨
