# Phase 1 详细操作指南

> Specify - 需求规范化，AI开发的基石

---

## 🎯 本阶段目标

将模糊的功能需求，转化为清晰的**业务需求规范**。

**核心原则** (对齐GitHub Spec Kit)：
> "Focus on the what and why, not the tech stack"

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

### Step 3: Business Rules（业务规则）

> **新增**：对齐Spec Kit，聚焦业务逻辑

#### 3.1 什么是Business Rules？

**定义**：业务层面的规则和约束，**不涉及技术实现**。

**包含**：
- ✅ 业务流程规则（如：审批流程）
- ✅ 数据约束（如：唯一性、层级限制）
- ✅ 业务逻辑限制（如：删除前检查依赖）
- ✅ 业务值域（如：状态枚举）

**不包含**（这些属于Phase 2）：
- ❌ 技术架构（Handler/Logic/Model）
- ❌ ORM选择（GORM/SQLx）
- ❌ 编码规范（函数行数、注释）
- ❌ 数据库表结构

#### 3.2 Business Rules示例

**分类管理功能**：
```markdown
## Business Rules

### 层级规则
- 分类最多支持3层（顶级→二级→三级）
- 超过3层时系统拒绝创建

### 唯一性规则
- 分类名称在同一父分类下必须唯一
- 分类编码全局唯一

### 删除规则
- 删除分类前必须检查是否有子分类
- 存在子分类时禁止删除
- 删除使用软删除方式

### 编码规则
- 编码长度：2-20个字符
- 允许字符：字母、数字、下划线、连字符
- 格式示例：data-catalog、user_management

### 排序规则
- 同级分类支持自定义排序
- 默认按创建时间排序
```

#### 3.3 Data Considerations（数据考量）

描述需要持久化的数据，**但不定义表结构**：

```markdown
## Data Considerations

需要存储的信息：
- 分类基本信息（名称、编码、描述）
- 层级关系（父分类引用）
- 排序信息（同级排序值）
- 元数据（创建时间、更新时间）
- 软删除标记

数据关系：
- 分类之间存在父子关系（树形结构）
- 一个分类可以有多个子分类
- 一个分类只能有一个父分类
```

**注意**：具体的表结构、字段类型、Go结构体都在**Phase 2: Design**中定义。

#### 3.4 AI辅助生成

```
@spec/workflow/phase1-specify.md

功能：{功能描述}

Phase 1: 生成Business Rules

要求：
1. 聚焦业务规则，不涉及技术实现
2. 包含数据约束（唯一性、范围等）
3. 包含业务流程限制
4. 简要说明需要持久化的数据（Data Considerations）
```

---

### Step 4: 未确定项（Open Questions）

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

**Business Rules**：
- [ ] 业务规则完整
- [ ] 不包含技术实现细节
- [ ] 数据约束清晰

**Data Considerations**：
- [ ] 列出需要持久化的数据
- [ ] 描述数据关系
- [ ] **不定义表结构**（留给Phase 2）

**Open Questions**：
- [ ] 列出所有不确定项
- [ ] 每项都有建议答案

### Peer Review

```markdown
## Review Checklist for Reviewer

- [ ] User stories是否完整？
- [ ] EARS是否覆盖主要场景？
- [ ] Business Rules是否明确？
- [ ] **没有包含技术实现细节？**
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

### Q3: Business Rules和Technical Constraints如何区分？
**A**: 
- Business Rules：业务层面（如层级限制、唯一性）→ Phase 1
- Technical Constraints：技术层面（如架构、ORM）→ Phase 2

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

## Business Rules
业务规则和约束

## Data Considerations
需要持久化的数据描述（不是表结构）

## Open Questions
待澄清问题列表

## Next Steps
→ Phase 2: Design
```

---

**Phase 1做好了，后面一路绿灯！** 📝
