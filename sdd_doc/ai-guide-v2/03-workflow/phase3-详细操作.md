# Phase 3 详细操作指南

> Tasks - 任务拆分，化整为零

---

## 🎯 本阶段目标

将Phase 2的技术设计，拆分为可独立执行的小任务。

**时间投入**：
- 简单功能：10-15分钟
- 中等功能：20-30分钟
- 复杂功能：30-60分钟

**核心产物**：`tasks.md`

---

## 📋 详细操作步骤

### Step 1: 任务拆分原则

#### 1.1 单一职责原则

每个task只做一件事：
- ✅ Task 1: 创建Model接口
- ✅ Task 2: 实现GORM DAO
- ❌ Task 1: 创建Model接口和实现GORM（太大）

#### 1.2 小任务原则

**黄金法则**：每个task ≤ 50行代码

**评估方法**：
```
Task预估行数 = 
  函数签名 + 
  方法体 + 
  注释 + 
  测试代码（简化）
```

**示例**：
```
Task: 创建Model接口
- 接口定义：8个方法 × 2行 = 16行
- 注释：8行
- 预估：24行 ✅
```

#### 1.3 依赖顺序原则

**自下而上**：
```
Task 1: Model接口       （无依赖）
Task 2: GORM实现        （依赖Task 1）
Task 3: SQLx实现        （依赖Task 1）
Task 4: Factory         （依赖Task 2, 3）
Task 5: Logic           （依赖Task 4）
Task 6: Handler         （依赖Task 5）
```

---

### Step 2: 任务模板

#### 2.1 标准Task格式

```markdown
## Task {编号}: {任务名称}

**Status**: ⏸️ Not Started / 🚧 In Progress / ✅ Completed

**Priority**: P0 / P1 / P2

**Estimated Lines**: {预估行数}

**Depends On**: Task {依赖的task编号}

**Files**:
- {文件路径1}
- {文件路径2}

**Description**:
{详细描述，1-2句话}

**Acceptance Criteria**:
- [ ] {验收标准1}
- [ ] {验收标准2}

**Implementation Notes** (可选):
- {实现提示}
```

#### 2.2 实战示例

```markdown
## Task 1: 创建Category Model接口

**Status**: ⏸️ Not Started

**Priority**: P0

**Estimated Lines**: 30

**Depends On**: None

**Files**:
- model/resource_catalog/category/interface.go
- model/resource_catalog/category/types.go

**Description**:
定义Category的数据访问接口和数据结构

**Acceptance Criteria**:
- [ ] Model接口包含Insert/FindOne/List/Update/SoftDelete方法
- [ ] Category结构体定义完整
- [ ] 包含完整的中文注释

**Implementation Notes**:
- 参考：model/*/*/interface.go
- 使用context.Context作为第一个参数
```

---

### Step 3: 完整任务列表示例

#### 3.1 Category管理功能Tasks

```markdown
# Tasks: Category Management

## Overview
总计：15个tasks
预估总行数：~450行
预估时间：4-6小时

---

## Model层（Tasks 1-4）

### Task 1: 创建Model接口和类型
**Status**: ⏸️ Not Started
**Lines**: 30
**Files**: interface.go, types.go
**Criteria**:
- [ ] Model接口定义8个方法
- [ ] Category结构体完整
- [ ] 错误定义（errors.go）

### Task 2: 实现GORM DAO
**Status**: ⏸️ Not Started
**Lines**: 120
**Depends**: Task 1
**Files**: gorm_dao.go
**Criteria**:
- [ ] 实现所有Model接口方法
- [ ] 包含软删除处理
- [ ] 完整中文注释

### Task 3: 实现SQLx Model
**Status**: ⏸️ Not Started
**Lines**: 100
**Depends**: Task 1
**Files**: sqlx_model.go
**Criteria**:
- [ ] 实现关键CRUD方法
- [ ] 使用sqlx.DB
- [ ] 性能优化（prepared statements）

### Task 4: 实现Factory
**Status**: ⏸️ Not Started
**Lines**: 40
**Depends**: Task 2, Task 3
**Files**: factory.go
**Criteria**:
- [ ] 支持配置切换ORM
- [ ] 单例模式
- [ ] 线程安全

---

## Logic层（Tasks 5-9）

### Task 5: 创建分类Logic
**Status**: ⏸️ Not Started
**Lines**: 45
**Depends**: Task 4
**Files**: createcategorylogic.go
**Criteria**:
- [ ] 名称唯一性检查
- [ ] 层级限制检查（≤3）
- [ ] 完整错误处理

### Task 6: 查询列表Logic
**Status**: ⏸️ Not Started
**Lines**: 35
**Depends**: Task 4
**Files**: listcategorylogic.go
**Criteria**:
- [ ] 支持分页
- [ ] 支持parent_id筛选
- [ ] 返回总数

### Task 7: 查询详情Logic
**Status**: ⏸️ Not Started
**Lines**: 25
**Depends**: Task 4
**Files**: getcategorylogic.go
**Criteria**:
- [ ] 根据ID查询
- [ ] 不存在返回404错误
- [ ] 包含子分类数量

### Task 8: 更新分类Logic
**Status**: ⏸️ Not Started
**Lines**: 40
**Depends**: Task 4
**Files**: updatecategorylogic.go
**Criteria**:
- [ ] 部分更新支持
- [ ] 名称唯一性检查
- [ ] 乐观锁（可选）

### Task 9: 删除分类Logic
**Status**: ⏸️ Not Started
**Lines**: 35
**Depends**: Task 4
**Files**: deletecategorylogic.go
**Criteria**:
- [ ] 软删除
- [ ] 检查是否有子分类
- [ ] 级联处理策略

---

## Handler层（Tasks 10-14）

### Task 10: 创建分类Handler
**Status**: ⏸️ Not Started
**Lines**: 30
**Depends**: Task 5
**Files**: createcategoryhandler.go
**Criteria**:
- [ ] 参数绑定和验证
- [ ] 调用Logic.Create
- [ ] 返回201和分类ID

### Task 11-14: 其他CRUD Handlers
（类似Task 10）

---

## 测试（Task 15）

### Task 15: 单元测试
**Status**: ⏸️ Not Started
**Lines**: 150
**Depends**: All above
**Files**: *_test.go
**Criteria**:
- [ ] Model层测试覆盖>80%
- [ ] Logic层表驱动测试
- [ ] Handler层Mock测试
```

---

### Step 4: 任务依赖图

#### 4.1 可视化依赖

```markdown
### 任务依赖关系

\`\`\`mermaid
graph TD
    T1[Task 1: Model接口]
    T2[Task 2: GORM]
    T3[Task 3: SQLx]
    T4[Task 4: Factory]
    T5[Task 5: Create Logic]
    T6[Task 6: List Logic]
    T7[Task 7: Get Logic]
    T8[Task 8: Update Logic]
    T9[Task 9: Delete Logic]
    T10[Task 10: Create Handler]
    T11[Task 11: List Handler]
    T12[Task 12: Get Handler]
    T13[Task 13: Update Handler]
    T14[Task 14: Delete Handler]
    T15[Task 15: Tests]
    
    T1 --> T2
    T1 --> T3
    T2 --> T4
    T3 --> T4
    T4 --> T5
    T4 --> T6
    T4 --> T7
    T4 --> T8
    T4 --> T9
    T5 --> T10
    T6 --> T11
    T7 --> T12
    T8 --> T13
    T9 --> T14
    T10 --> T15
    T11 --> T15
    T12 --> T15
    T13 --> T15
    T14 --> T15
\`\`\`
```

#### 4.2 执行顺序建议

**并行执行**（如果多人）：
- 一人：Task 1-4（Model层）
- 二人：Task 5-9（Logic层，等Task 4完成）
- 三人：Task 10-14（Handler层，等Task 5-9完成）

**串行执行**（单人）：
1. Task 1-4（Model层，1-2小时）
2. Task 5-9（Logic层，1-2小时）
3. Task 10-14（Handler层，1小时）
4. Task 15（测试，1-2小时）

---

### Step 5: 任务拆分技巧

#### 5.1 如何判断task太大？

**信号**：
- 预估>50行
- 涉及多个文件
- 包含多个功能点

**解决**：进一步拆分
```
Task: 实现Category CRUD
↓ 拆分为
Task 1: 实现Insert
Task 2: 实现FindOne
Task 3: 实现List
Task 4: 实现Update
Task 5: 实现Delete
```

#### 5.2 如何判断task太小？

**信号**：
- <10行
- 过于琐碎
- 执行成本高于收益

**解决**：合并任务
```
Task 1: 定义接口
Task 2: 定义类型
↓ 合并为
Task: 定义Model接口和类型
```

#### 5.3 AI辅助拆分

**Cursor**：
```
@design.md
@spec/coding-standards/go-style-guide.md

Phase 3: 请生成Tasks

要求：
1. 每个task <50行
2. 明确依赖关系
3. 自下而上（Model → Logic → Handler）
4. 包含验收标准
```

**Kiro.dev**：
```
在Implementation阶段
Kiro自动将Design拆分为Tasks
可调整粒度
```

---

## ✅ Gate 3 质量检查

### 自检清单

**任务粒度**：
- [ ] 每个task ≤50行
- [ ] 职责单一
- [ ] 可独立完成

**任务完整性**：
- [ ] 覆盖所有设计内容
- [ ] Model/Logic/Handler都有
- [ ] 包含测试任务

**依赖关系**：
- [ ] 依赖关系明确
- [ ] 无循环依赖
- [ ] 自下而上顺序

**验收标准**：
- [ ] 每个task都有验收标准
- [ ] 标准清晰可测试
- [ ] 覆盖功能和质量

**预估合理**：
- [ ] 行数预估合理
- [ ] 工时预估合理
- [ ] 总量与需求匹配

---

## 💡 实战技巧

### 技巧1：先骨架后血肉

第一轮：创建所有文件和接口
```
Task 1: Model接口
Task 2: Logic接口  
Task 3: Handler签名
```

第二轮：实现核心逻辑
```
Task 4: Model实现
Task 5: Logic实现
Task 6: Handler实现
```

第三轮：完善测试
```
Task 7: 单元测试
```

### 技巧2：按优先级标记

**P0**：核心功能，必须完成
**P1**：重要功能，应该完成
**P2**：次要功能，可选完成

### 技巧3：时间盒

给每个task设定时间上限：
```
Task 1: Model接口 - 最多30分钟
Task 2: GORM实现 - 最多1.5小时
```

超时就Review是否拆分不够细。

### 技巧4：Checklist驱动

```markdown
## 今日任务
- [ ] ✅ Task 1: Model接口
- [ ] 🚧 Task 2: GORM实现
- [ ] ⏸️ Task 3: SQLx实现
```

---

## 🔧 常见问题

### Q1: 一定要拆分得这么细吗？
**A**: 
- 简单功能：可以粗一点
- 复杂功能：必须细
- 团队协作：必须细

### Q2: 拆分tasks花太多时间？
**A**: 
- 前期投入，后期节省
- 用AI辅助生成
- 第一次慢，第二次快

### Q3: 中途发现tasks不合理？
**A**: 
- 立即调整
- 这很正常
- 不断优化

### Q4: 如何追踪进度？
**A**:
- tasks.md实时更新状态
- Kiro.dev自动追踪
- GitHub Projects

---

## 📊 Phase 3输出

### tasks.md结构

```markdown
# Tasks: {功能名}

## Overview
- 总计：X个tasks
- 预估行数：~XXX
- 预估时间：X-X小时

## Task Dependencies
（依赖图）

## Model层
### Task 1-N: ...

## Logic层
### Task N+1: ...

## Handler层
### Task M: ...

## Testing
### Task Z: ...

## Progress Tracking
- [ ] Task 1
- [X] Task 2
...
```

---

**Phase 3做好了，执行清晰高效！** ✂️
