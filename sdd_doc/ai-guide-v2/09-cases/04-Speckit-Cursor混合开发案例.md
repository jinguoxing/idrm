# 实战案例：Spec Kit + Cursor Agent 混合开发

> 规范驱动 + AI辅助 - 最佳组合工作流

---

## 📋 案例概述

**功能**：数据标签（Tag）管理  
**复杂度**：中等（约400行）  
**工具**：[GitHub Spec Kit](https://github.com/github/spec-kit) + Cursor Agent  
**耗时**：1个工作日  
**适用场景**：规范化团队开发、企业级项目

---

## 💡 为什么选择 Spec Kit + Cursor？

### Spec Kit 的优势
- 🎯 **规范驱动**：GitHub 官方 Spec-Driven Development 工具包
- 📋 **模板化**：IDRM 定制模板（EARS + 分层架构）
- ✅ **验证机制**：内置规范检查
- 🔗 **原生集成**：与 GitHub 和多种 AI 工具完美集成

### Cursor Agent 的优势
- 💬 **自然交互**：对话式开发体验
- 🔍 **上下文理解**：理解项目规范
- ⚡ **快速实现**：高效代码生成
- 🎨 **实时反馈**：即时预览效果

### 组合使用的价值
✅ **Spec Kit** 生成规范化文档 → **Cursor** 基于规范实现代码  
✅ 规范和实现完美对齐  
✅ 团队协作标准统一  
✅ 可追溯、可审计

---

## 🛠️ 项目结构

IDRM 项目已配置好 Spec Kit：

```
项目根目录/
├── .specify/                       # Spec Kit 配置
│   ├── memory/
│   │   └── constitution.md         # 项目原则（IDRM 宪章）
│   └── templates/
│       ├── spec-template.md        # 需求规范模板（EARS + User Stories）
│       ├── plan-template.md        # 技术设计模板（分层架构 + 双ORM）
│       ├── tasks-template.md       # 任务拆分模板（<50行任务）
│       ├── api-template.api        # go-zero API 模板
│       └── schema-template.sql     # DDL 模板
├── .github/
│   └── prompts/
│       ├── specify.prompt.md       # 需求阶段 AI 引导
│       ├── plan.prompt.md          # 设计阶段 AI 引导
│       └── tasks.prompt.md         # 任务阶段 AI 引导
└── specs/                          # 生成的规范文件
    └── {feature}/
        ├── spec.md
        ├── plan.md
        └── tasks.md
```

---

## 🎯 需求背景

数据管理员希望为数据资源打标签，便于分类和检索。

---

## 📝 Phase 0: Context (15分钟)

### Step 1: 阅读项目规范

在 Cursor 中打开项目：

```bash
cursor .
```

**打开 Cursor Agent (Cmd+L)**，阅读已有的项目规范：

```
请阅读以下项目规范文件：
@CLAUDE.md
@.specify/memory/constitution.md
@sdd_doc/spec/core/workflow.md
@sdd_doc/spec/architecture/layered-architecture.md

总结关键规范要求，包括：
- 架构原则（Handler→Logic→Model 分层）
- 编码规范（函数行数限制、注释要求）
- 质量要求（测试覆盖率）
```

> **注意**: 项目已有 `constitution.md`，无需使用 `/speckit.constitution` 创建。

---

## 📋 Phase 1: Specify (30分钟)

### Step 1: 使用 /speckit.specify 创建需求规范

**Cursor Agent Prompt**:

```
/speckit.specify 数据标签管理功能

功能需求：
1. 创建/删除标签
2. 为资源打标签/取消标签
3. 按标签查询资源
4. 标签统计

参考模板：@.specify/templates/spec-template.md

要求：
1. User Stories 使用 AS/I WANT/SO THAT 格式
2. Acceptance Criteria 使用 EARS notation (WHEN... THE SYSTEM SHALL...)
3. 覆盖正常、异常、边界场景
4. 定义 Business Rules（唯一性、长度、关联、删除）
5. 定义 Data Considerations
6. **不包含技术实现细节**
```

**/speckit.specify** 会生成 **specs/tag-management/spec.md** 文件。

### Step 2: 使用 /speckit.clarify 澄清问题

```
/speckit.clarify

请检查需求规范是否有遗漏或模糊的地方，如果有请自动补充。
```

### Step 3: 提交 Phase 1

```bash
git add specs/
git commit -m "docs: complete tag management requirements (Phase 1)

Generated with GitHub Spec Kit + IDRM templates"
```

---

## 🎨 Phase 2: Design (40分钟)

### Step 1: 使用 /speckit.plan 创建技术方案

**Cursor Agent Prompt**：

```
/speckit.plan 使用以下技术栈：

参考模板：@.specify/templates/plan-template.md
参考规范：
@sdd_doc/spec/architecture/layered-architecture.md
@sdd_doc/spec/architecture/dual-orm-pattern.md

技术要求：
- Go-Zero 微服务框架
- 遵循分层架构（Handler→Logic→Model）
- 数据库设计：`tags` 表和 `resource_tags` 关联表
- API 接口：使用 go-zero .api 格式定义
- ORM：选择 GORM（复杂查询）
- 序列图：用 Mermaid 描述"为资源打标签"的流程
```

**/speckit.plan** 会生成 **specs/tag-management/plan.md** 文件。

### Step 2: 生成 API 和 DDL 文件

```
基于 plan.md，请生成：
1. `api/doc/resource_catalog/tag.api` - 参考 @.specify/templates/api-template.api
2. `migrations/resource_catalog/tags.sql` - 参考 @.specify/templates/schema-template.sql

遵循项目的文件路径规范。
```

### Step 3: 提交 Phase 2

```bash
git add specs/ api/doc/ migrations/
git commit -m "docs: complete tag management design (Phase 2)

Generated with GitHub Spec Kit + IDRM templates"
```

---

## 📋 Phase 3: Tasks (20分钟)

### 使用 /speckit.tasks 拆分任务

**Cursor Agent Prompt**:

```
/speckit.tasks

参考模板：@.specify/templates/tasks-template.md

要求：
1. 基于 plan.md 拆分任务
2. 每个 task < 50行代码
3. 明确依赖关系
4. 按顺序：Model → Logic → Handler → Test
```

**/speckit.tasks** 会生成 **specs/tag-management/tasks.md** 文件。

---

## 💻 Phase 4: Implement (4-6小时)

### Step 1: 使用 /speckit.implement 开始实现

```
/speckit.implement

请按照 tasks.md 中的任务列表逐个实现。
首先生成代码框架：

1. 运行 `goctl api go -api api/doc/resource_catalog/tag.api -dir api/ --style=goZero`
2. 运行 `goctl model mysql ddl -src migrations/resource_catalog/tags.sql -dir model/resource_catalog/tag/ --style=goZero`
```

### Step 2: 实现 Model 层

**Cursor Composer (Cmd+I)**：

```
@specs/tag-management/plan.md
@sdd_doc/spec/coding-standards/go-style-guide.md

请实现 Model 层：

文件：
- model/resource_catalog/tag/interface.go
- model/resource_catalog/tag/types.go  
- model/resource_catalog/tag/gorm_dao.go
- model/resource_catalog/tag/factory.go

要求：
- 遵循 plan.md 的接口定义
- 完整的中文注释
- 错误处理
- 每个函数<50行
```

### Step 3: 实现 Logic 层

```
@model/resource_catalog/tag/interface.go
@specs/tag-management/plan.md

请实现 Logic 层，为每个功能创建 Logic 文件：
- createtaglogic.go
- deletetaglogic.go
- assigntaglogic.go
- removetaglogic.go
- querybytaglogic.go
- tagstatslogic.go

要求：
- 业务逻辑实现
- 调用 Model 接口
- 完整错误处理
- 函数<50行
- 中文注释
```

### Step 4: 生成测试

```
@api/internal/logic/resource_catalog/tag/*.go
@sdd_doc/spec/coding-standards/testing-standards.md

为所有 Logic 生成单元测试

要求：
- 表驱动测试
- Mock Model 接口
- 覆盖率>80%
- 测试正常和异常流程
```

### Step 5: 运行测试

```bash
go test -cover ./model/resource_catalog/tag/
go test -cover ./api/internal/logic/resource_catalog/tag/
```

---

## ✅ Gate 4: 质量检查

### 使用 /speckit.checklist 运行检查

```
/speckit.checklist

请检查实现是否符合规范：
- 所有需求都有对应实现
- 代码覆盖率 > 80%
- 所有函数 < 50 行
- 所有公共函数有注释
- 无硬编码值
```

### 使用 Cursor 进行 Code Review

```
@model/resource_catalog/tag/*.go
@api/internal/logic/resource_catalog/tag/*.go
@specs/tag-management/plan.md

请 Review 这些代码：

检查：
1. 是否符合 plan.md 的设计？
2. 是否遵循分层架构？
3. 函数是否<50行？
4. 错误处理是否完整？
5. 注释是否完整？

请给出具体改进建议
```

---

## 🔄 Spec Kit + Cursor 集成工作流总结

### 最佳实践流程

```
1. 阅读 @.specify/memory/constitution.md  → 理解项目原则
2. /speckit.specify                        → 创建需求规范（生成 spec.md）
3. /speckit.clarify                        → 澄清问题
4. /speckit.plan                           → 创建技术方案（生成 plan.md）
5. /speckit.tasks                          → 拆分任务（生成 tasks.md）
6. /speckit.implement                      → 实现代码
7. /speckit.checklist                      → 质量检查
8. Cursor Agent                            → 修正问题
```

---

## 💡 组合使用技巧

### 1. 使用 IDRM 模板 + Cursor 填充

Spec Kit 的模板已针对 IDRM 项目定制：
- `spec-template.md` → EARS notation + User Stories
- `plan-template.md` → 分层架构 + 双ORM + Mermaid
- `tasks-template.md` → <50行任务拆分

### 2. 使用 AI 引导文件

`.github/prompts/` 下的文件会引导 AI 使用正确的规范：
- `specify.prompt.md` → 引导生成 EARS 格式需求
- `plan.prompt.md` → 引导生成分层架构设计
- `tasks.prompt.md` → 引导任务拆分

### 3. 使用 Cursor 生成 + /speckit.checklist 验证

```
# Cursor 快速生成
# → /speckit.checklist 验证
# → Cursor 修正
# → /speckit.checklist 验证
# → 通过
```

---

## 🎯 工具职责分工

### Spec Kit + IDRM 模板负责

- ✅ 规范化文档结构（EARS, User Stories, 分层架构）
- ✅ 引导开发流程（SDD Workflow）
- ✅ 验证规范依从性（Checklist）
- ✅ 团队协作和审计

### Cursor Agent 负责

- ✅ 填充文档内容
- ✅ 生成代码实现
- ✅ 编写测试用例
- ✅ Code Review
- ✅ 快速迭代修改

---

## 📊 时间对比

| 方式 | Phase 1 | Phase 2 | Phase 3 | Phase 4 | 总计 |
|------|---------|---------|---------|---------|------|
| 纯手动 | 60min | 90min | 40min | 8h | **10.5h** |
| Cursor | 30min | 40min | 20min | 5h | **6.5h** |
| **Spec Kit+Cursor** | **20min** | **30min** | **15min** | **4h** | **5h** |

**优势**：
- IDRM 定制模板（EARS, 分层架构）
- AI 引导文件自动应用规范
- 流程引导（SDD Workflow）
- 验证自动化（Checklist）

---

## 📚 Spec Kit 命令参考

### Slash Commands (Cursor Agent 对话中使用)
| 命令 | 用途 | 输出文件 |
| :--- | :--- | :--- |
| `/speckit.specify` | 创建需求规范 | `spec.md` |
| `/speckit.clarify` | 澄清规范问题 | - |
| `/speckit.plan` | 创建技术方案 | `plan.md` |
| `/speckit.tasks` | 拆分任务 | `tasks.md` |
| `/speckit.implement` | 执行实现 | 代码文件 |
| `/speckit.checklist` | 验证规范依从性 | - |

### IDRM 模板文件
| 模板 | 位置 | 用途 |
| :--- | :--- | :--- |
| `spec-template.md` | `.specify/templates/` | EARS + User Stories |
| `plan-template.md` | `.specify/templates/` | 分层架构 + 双ORM |
| `tasks-template.md` | `.specify/templates/` | <50行任务拆分 |
| `api-template.api` | `.specify/templates/` | go-zero API 格式 |
| `schema-template.sql` | `.specify/templates/` | DDL 格式 |

---

## 🎯 总结

### Spec Kit + IDRM 模板 + Cursor 的完美配合

1. **Spec Kit + IDRM 模板** 提供结构和规范
   - EARS notation 标准化需求
   - 分层架构标准化设计
   - 任务拆分标准化实现

2. **Cursor** 提供内容和实现
   - 快速填充文档
   - 高效生成代码
   - 即时反馈调整

3. **两者结合** 实现最优工作流
   - 规范和实现对齐
   - 质量和效率并重
   - 可追溯可审计

### 适用场景

✅ **企业级项目** - 需要规范化管理  
✅ **团队协作** - 需要统一标准  
✅ **长期维护** - 需要文档追溯  
✅ **合规要求** - 需要审计证据

---

**官方文档**：[github/spec-kit](https://github.com/github/spec-kit)

**Spec Kit + IDRM 模板 + Cursor = 规范化 + 高效率！** 🚀
