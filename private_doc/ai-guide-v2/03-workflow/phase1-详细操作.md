# Phase 1 详细操作指南

> Specify - 需求规范化，AI开发的基石

---

## 🎯 本阶段目标

将模糊的功能需求，转化为清晰、可测试的规范文档。

**时间投入**：
- 简单功能：15-20分钟
- 中等功能：30-40分钟  
- 复杂功能：1-2小时

**核心产物**：`requirements.md`

---

## 📋 详细操作步骤

### Step 1: 编写用户故事（User Stories）

#### 1.1 User Story模板

```markdown
## User Story: {功能名称}

AS a {角色}
I WANT {功能}
SO THAT {价值/目标}
```

#### 1.2 实战示例

**示例1：创建分类**
```markdown
## User Story: 创建资源分类

AS a 数据管理员
I WANT 创建新的资源分类
SO THAT 我可以更好地组织数据资源
```

**示例2：查询分类**
```markdown
## User Story: 查询分类列表

AS a 数据管理员
I WANT 查询所有分类并支持分页
SO THAT 我可以快速浏览和管理大量分类
```

**示例3：删除分类**
```markdown
## User Story: 删除分类

AS a 数据管理员
I WANT 删除不再使用的分类
SO THAT 保持分类列表的整洁
```

#### 1.3 编写技巧

**DO ✅**：
- 从用户视角出发
- 价值明确
- 简洁明了

**DON'T ❌**：
- 不要写技术实现
- 不要太抽象
- 不要漏掉SO THAT

#### 1.4 AI辅助生成

**Cursor**：
```
我要添加{功能描述}

Phase 1: 请生成User Stories
要求：
- 至少3个story
- 从用户视角
- 价值明确
```

**Kiro.dev**：
```
在Requirements阶段
描述功能，Kiro自动生成User Stories
```

---

### Step 2: EARS Notation（核心！）

#### 2.1 EARS是什么？

**EARS** = **E**asy **A**pproach to **R**equirements **S**yntax

**格式**：
```
WHEN [触发条件/事件]
THE SYSTEM SHALL [期望行为]
```

**为什么用EARS？**
- ✅ 清晰：条件和行为明确
- ✅ 可测试：每条都能验证
- ✅ 完整：覆盖正常和异常

#### 2.2 EARS基础语法

**模板1：基本格式**
```
WHEN {用户操作} {输入条件}
THE SYSTEM SHALL {系统行为} {输出结果}
```

**模板2：异常情况**
```
WHEN {用户操作} {错误输入}
THE SYSTEM SHALL {错误处理} {错误信息}
```

**模板3：边界情况**
```
WHEN {特殊条件}
THE SYSTEM SHALL {特殊行为}
```

#### 2.3 EARS实战示例

**功能：创建分类**

**正常流程**：
```markdown
WHEN 用户提交有效的分类创建请求
THE SYSTEM SHALL 保存分类到数据库并返回201状态码和分类ID

WHEN 用户提交的分类数据包含name、code、parent_id
THE SYSTEM SHALL 验证name长度1-50字符，code符合编码规范
```

**异常流程**：
```markdown
WHEN 用户提交的分类名称为空
THE SYSTEM SHALL 返回400错误和"名称不能为空"的错误信息

WHEN 用户提交的分类名称与已存在分类重复
THE SYSTEM SHALL 返回409错误和"分类名称已存在"的错误信息

WHEN 用户提交的parent_id不存在
THE SYSTEM SHALL 返回404错误和"父分类不存在"的错误信息
```

**边界情况**：
```markdown
WHEN 用户创建的分类层级超过3层
THE SYSTEM SHALL 返回400错误和"分类层级不能超过3层"的错误信息

WHEN 用户提交的分类名称长度超过50字符
THE SYSTEM SHALL 返回400错误和"名称长度不能超过50字符"的错误信息
```

#### 2.4 EARS常见错误

**❌ 错误示例1：没有条件**
```
THE SYSTEM SHALL 保存分类
```
**问题**：什么时候保存？什么条件下？

**✅ 正确版本**：
```
WHEN 用户提交有效的分类创建请求
THE SYSTEM SHALL 保存分类到数据库
```

---

**❌ 错误示例2：行为不明确**
```
WHEN 用户提交无效数据
THE SYSTEM SHALL 报错
```
**问题**：怎么报错？什么错误？

**✅ 正确版本**：
```
WHEN 用户提交的名称为空
THE SYSTEM SHALL 返回400错误和"名称不能为空"的验证错误
```

---

**❌ 错误示例3：包含实现细节**
```
WHEN 用户点击保存按钮
THE SYSTEM SHALL 调用CreateCategory API并使用GORM插入数据库
```
**问题**：太技术化，这是实现细节

**✅ 正确版本**：
```
WHEN 用户提交分类创建请求
THE SYSTEM SHALL 保存分类并返回分类ID
```

#### 2.5 EARS覆盖清单

**必须覆盖**：
- [ ] 正常流程（Happy Path）
- [ ] 参数验证失败
- [ ] 业务规则违反
- [ ] 权限不足
- [ ] 资源不存在

**应该覆盖**：
- [ ] 并发冲突
- [ ] 网络超时
- [ ] 数据库错误

**可选覆盖**：
- [ ] 性能要求（响应时间）
- [ ] 容量限制（最大数据量）

#### 2.6 AI辅助生成EARS

**Cursor示例**：
```
功能：创建资源分类

Phase 1: 请生成Acceptance Criteria，使用EARS notation

要求：
1. 覆盖正常流程
2. 覆盖所有参数验证
3. 覆盖业务规则（如层级限制）
4. 覆盖异常情况
5. 每条EARS必须可测试
```

**Kiro.dev**：
```
在Requirements阶段
Kiro会自动使用EARS格式
你可以Review并补充遗漏的场景
```

---

### Step 3: 技术约束（Technical Constraints）

#### 3.1 IDRM标准约束

**必须声明**：
```markdown
## Technical Constraints

### 架构约束
- MUST follow layered architecture (Handler → Logic → Model)
- MUST implement dual ORM support (GORM + SQLx)

### 编码约束
- Functions MUST be < 50 lines
- MUST use Chinese comments for all public items
- Error wrapping MUST use %w

### 质量约束
- Test coverage MUST be > 80%
- MUST pass golangci-lint check
- MUST follow go-style-guide
```

#### 3.2 功能特定约束

**示例：分类管理**
```markdown
### 业务约束
- Category hierarchy MUST NOT exceed 3 levels
- Category name MUST be unique within parent
- Category code MUST match pattern: ^[a-zA-Z0-9_-]{2,20}$

### 性能约束
- List API response time SHOULD be < 200ms
- MUST support pagination (max 100 items per page)

### 数据约束
- name: 1-50 characters, non-empty
- code: 2-20 characters, alphanumeric + underscore/hyphen
- parent_id: nullable, must reference existing category
```

#### 3.3 AI辅助生成

```
@private_doc/spec/core/project-charter.md
@private_doc/spec/architecture/layered-architecture.md

基于功能：{功能描述}

请生成Technical Constraints
包括：
1. IDRM通用约束
2. 本功能特定的业务约束
3. 性能和数据约束
```

---

### Step 4: 数据模型（Data Model）

#### 4.1 表结构定义

```markdown
## Data Model

### categories表

\`\`\`sql
CREATE TABLE categories (
    id          BIGINT PRIMARY KEY AUTO_INCREMENT,
    name        VARCHAR(50) NOT NULL COMMENT '分类名称',
    code        VARCHAR(20) NOT NULL UNIQUE COMMENT '分类编码',
    parent_id   BIGINT NULL COMMENT '父分类ID',
    level       TINYINT NOT NULL DEFAULT 1 COMMENT '层级(1-3)',
    sort_order  INT NOT NULL DEFAULT 0 COMMENT '排序',
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  DATETIME NULL COMMENT '软删除时间',
    
    INDEX idx_parent_id (parent_id),
    INDEX idx_code (code),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源分类表';
\`\`\`
```

#### 4.2 字段说明

```markdown
### 字段定义

| 字段 | 类型 | 约束 | 说明 |
|------|------|------|------|
| id | BIGINT | PK, AUTO_INCREMENT | 主键 |
| name | VARCHAR(50) | NOT NULL | 分类名称，1-50字符 |
| code | VARCHAR(20) | NOT NULL, UNIQUE | 分类编码，全局唯一 |
| parent_id | BIGINT | NULLABLE, FK | 父分类ID，NULL表示顶级 |
| level | TINYINT | NOT NULL, 1-3 | 层级，1=顶级，最多3层 |
| sort_order | INT | NOT NULL | 同级排序 |
| created_at | DATETIME | NOT NULL | 创建时间 |
| updated_at | DATETIME | NOT NULL | 更新时间 |
| deleted_at | DATETIME | NULLABLE | 软删除时间 |
```

#### 4.3 Go结构体

```go
// Category 资源分类
type Category struct {
    ID        int64      `gorm:"column:id;primaryKey" db:"id" json:"id"`
    Name      string     `gorm:"column:name;size:50;not null" db:"name" json:"name"`
    Code      string     `gorm:"column:code;size:20;not null;uniqueIndex" db:"code" json:"code"`
    ParentID  *int64     `gorm:"column:parent_id" db:"parent_id" json:"parent_id,omitempty"`
    Level     int8       `gorm:"column:level;not null;default:1" db:"level" json:"level"`
    SortOrder int        `gorm:"column:sort_order;not null;default:0" db:"sort_order" json:"sort_order"`
    CreatedAt time.Time  `gorm:"column:created_at;not null" db:"created_at" json:"created_at"`
    UpdatedAt time.Time  `gorm:"column:updated_at;not null" db:"updated_at" json:"updated_at"`
    DeletedAt *time.Time `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at,omitempty"`
}
```

---

### Step 5: 未确定项（Open Questions）

#### 5.1 列出所有不确定的问题

```markdown
## Open Questions

### Q1: 级联删除策略
**问题**：删除父分类时，子分类如何处理？
**选项**：
- A: 禁止删除有子分类的父分类
- B: 级联删除所有子分类
- C: 子分类提升一级

**建议**：选项A，保守安全
**待确认**：产品经理

### Q2: 编码规则  
**问题**：分类编码是手动输入还是自动生成？
**建议**：手动输入，给用户更大灵活性
**待确认**：产品经理

### Q3: 并发控制
**问题**：是否需要乐观锁防止并发修改？
**建议**：初期不需要，后续按需添加
**待确认**：技术负责人
```

#### 5.2 AI辅助提问

```
@requirements.md

分析这个需求规范，找出：
1. 有歧义的地方
2. 缺失的细节
3. 可能的边界情况

输出Open Questions清单
```

---

## ✅ Gate 1 质量检查

### 自检清单

**User Stories**：
- [ ] 至少3个user story
- [ ] 使用AS/I WANT/SO THAT格式
- [ ] 价值明确

**EARS Notation**：
- [ ] 所有关键流程都有EARS
- [ ] 正常流程完整
- [ ] 异常流程覆盖充分
- [ ] 每条EARS可测试

**Technical Constraints**：
- [ ] 包含IDRM通用约束
- [ ] 包含功能特定约束
- [ ] 约束清晰具体

**Data Model**：
- [ ] 表结构完整
- [ ] 字段定义清楚
- [ ] Go结构体正确

**Open Questions**：
- [ ] 列出所有不确定项
- [ ] 每项都有建议答案

### Peer Review

```markdown
## Review Checklist for Reviewer

- [ ] User stories是否完整？
- [ ] EARS是否覆盖主要场景？
- [ ] Technical constraints是否遵循规范？
- [ ] Data model是否合理？
- [ ] Open questions是否已澄清？
```

---

## 💡 实战技巧

### 技巧1：从简单开始

第一次写EARS可能觉得繁琐，从最核心的流程开始：
1. 先写Happy Path
2. 再补充参数验证
3. 最后补充边界情况

### 技巧2：参考现有代码

```bash
# 查看类似功能的requirements
find specs/ -name "requirements.md" -exec grep -l "category" {} \;

# 学习EARS如何编写
```

### 技巧3：用AI验证

```
@requirements.md

Review这个requirements文档：
1. EARS是否完整？
2. 有没有遗漏的场景？
3. 技术约束是否符合规范？

给出改进建议
```

### 技巧4：可视化

```markdown
## 状态机图（可选）

\`\`\`mermaid
stateDiagram-v2
    [*] --> Draft: 创建
    Draft --> Active: 发布
    Active --> Inactive: 停用
    Inactive --> Active: 启用
    Active --> [*]: 删除
\`\`\`
```

---

## 🔧 常见问题

### Q1: EARS写得太长怎么办？
**A**: 正常情况建议拆分功能。一个功能的requirements应控制在200行内。

### Q2: 不知道写哪些EARS？
**A**: 
1. 至少覆盖Happy Path
2. 每个输入参数都要验证
3. 每个业务规则都要表达

### Q3: Technical Constraints记不住？
**A**: 
```
@spec/core/project-charter.md
列出IDRM的Technical Constraints模板
```

### Q4: Phase 1需要多详细？
**A**: 
- 简单功能：30-50行
- 中等功能：100-150行
- 复杂功能：200-300行

---

## 📊 Phase 1输出

### requirements.md结构

```markdown
# Feature: {功能名}

## Overview
简述（2-3行）

## User Stories
至少3个

## Acceptance Criteria (EARS)
分类组织，每类5-10条

## Technical Constraints
IDRM通用 + 功能特定

## Data Model
表结构 + 字段说明 + Go结构体

## Open Questions
待澄清问题列表

## Next Steps
→ Phase 2: Design
```

---

**Phase 1做好了，后面一路绿灯！** 📝
