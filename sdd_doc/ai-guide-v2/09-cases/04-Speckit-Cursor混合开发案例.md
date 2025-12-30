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
- 📋 **模板化**：自动生成标准化文档结构
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

## 🛠️ 工具准备

### 1. 安装 Specify CLI

```bash
# 使用 uv 安装（推荐）
uv tool install specify-cli --from git+https://github.com/github/spec-kit.git

# 验证安装
specify check

# 查看帮助
specify --help
```

### 2. 初始化 Spec Kit

```bash
# 在项目根目录初始化（选择 Cursor Agent）
cd /path/to/idrm
specify init . --ai cursor-agent

# 这会创建：
# .specify/          # Spec Kit 配置和模板
# .github/           # AI Agent 的 prompt 文件
```

### 3. 配置 Spec Kit

初始化后，Spec Kit 会自动配置 Cursor Agent 可用的 slash commands。

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
@sdd_doc/spec/core/workflow.md
@sdd_doc/spec/architecture/layered-architecture.md
@sdd_doc/spec/constitution.md

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

要求：
1. User Stories 使用 AS/I WANT/SO THAT 格式
2. Acceptance Criteria 使用 EARS notation (WHEN... THE SYSTEM SHALL...)
3. 覆盖正常、异常、边界场景
4. 定义 Business Rules（唯一性、长度、关联、删除）
5. 定义 Data Considerations
6. **不包含技术实现细节**
```

**/speckit.specify** 会生成 **spec.md** 文件，包含完整的需求规范。

### Step 2: 使用 /speckit.clarify 澄清问题

```
/speckit.clarify

请检查需求规范是否有遗漏或模糊的地方，如果有请自动补充。
```

### Step 3: 提交 Phase 1

```bash
git add .speckit/
git commit -m "docs: complete tag management requirements (Phase 1)

Generated with GitHub Spec Kit
Enhanced with Cursor Agent"
```

---

## 🎨 Phase 2: Design (40分钟)

### Step 1: 使用 /speckit.plan 创建技术方案

**Cursor Agent Prompt**：

```
/speckit.plan 使用以下技术栈：

- Go-Zero 微服务框架
- 遵循分层架构（Handler→Logic→Model）
- 数据库设计：`tags` 表和 `resource_tags` 关联表，包含完整索引设计
- API 接口：使用 go-zero .api 格式定义
- ORM：选择 GORM（复杂查询）
- 序列图：用 Mermaid 描述"为资源打标签"的流程

参考：
@sdd_doc/spec/architecture/layered-architecture.md
@sdd_doc/spec/architecture/dual-orm-pattern.md
```

**/speckit.plan** 会生成 **plan.md** 文件，包含完整的技术设计。

### Step 2: 生成 API 和 DDL 文件

```
基于 plan.md，请生成：
1. `api/doc/resource_catalog/tag.api` - go-zero API 定义
2. `migrations/resource_catalog/tags.sql` - DDL 文件

遵循项目的文件路径规范。
```

### Step 3: 提交 Phase 2

```bash
git add .speckit/ api/doc/ migrations/
git commit -m "docs: complete tag management design (Phase 2)

Generated with GitHub Spec Kit
Enhanced with Cursor Agent"
```

---

## 📋 Phase 3: Tasks (20分钟)

### 使用 /speckit.tasks 拆分任务

**Cursor Agent Prompt**:

```
/speckit.tasks

要求：
1. 基于 plan.md 拆分任务
2. 每个 task < 50行代码
3. 明确依赖关系
4. 按顺序：Model → Logic → Handler → Test
```

**/speckit.tasks** 会生成 **tasks.md** 文件，包含任务列表。

---

## 💻 Phase 4: Implement (4-6小时)

### 方法：/speckit.implement + Cursor Agent 实现

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
@specs/features/tag-management/design.md
@sdd_doc/spec/coding-standards/go-style-guide.md

请实现 Model 层：

文件：
- model/resource_catalog/tag/interface.go
- model/resource_catalog/tag/types.go  
- model/resource_catalog/tag/gorm_dao.go
- model/resource_catalog/tag/factory.go

要求：
- 遵循 design 的接口定义
- 完整的中文注释
- 错误处理
- 每个函数<50行
```

### Step 3: 实现 Logic 层

**批量生成所有 Logic**：

```
@model/resource_catalog/tag/interface.go
@.speckit/plan.md

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
# 运行测试
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
@.speckit/plan.md

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
1. specify init . --ai cursor-agent  → 初始化项目结构
2. /speckit.constitution             → 建立项目原则
3. /speckit.specify                  → 创建需求规范
4. /speckit.clarify                  → 澄清问题
5. /speckit.plan                     → 创建技术方案
6. /speckit.tasks                    → 拆分任务
7. /speckit.implement                → 实现代码
8. /speckit.checklist                → 质量检查
9. Cursor Agent                      → 修正问题
```

---

## 💡 组合使用技巧

### 1. 使用 Spec Kit Slash Commands + Cursor 填充

Spec Kit 提供的 slash commands 可以直接在 Cursor Agent 对话中使用，自动生成规范化的文档结构。

### 2. 使用 Cursor 生成 + /speckit.checklist 验证

```
# Cursor 快速生成
# → /speckit.checklist 验证
# → Cursor 修正
# → /speckit.checklist 验证
# → 通过
```

这样保证生成的内容符合规范。

### 3. 使用 /speckit.analyze 分析代码质量

```
/speckit.analyze

分析当前代码的质量指标，给出改进建议。
```

---

## 🎯 工具职责分工

### Spec Kit 负责

- ✅ 规范化文档结构 (Slash Commands)
- ✅ 引导开发流程 (SDD Workflow)
- ✅ 验证规范依从性 (Checklist)
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
- 框架生成快（Spec Kit Slash Commands）
- 内容填充快（Cursor）
- 流程引导（SDD Workflow）
- 验证自动化（Checklist）

---

## 📚 Spec Kit 命令参考

### Specify CLI 命令
| 命令 | 用途 |
| :--- | :--- |
| `specify init . --ai cursor-agent` | 初始化项目（Cursor 支持） |
| `specify check` | 检查环境 |

### Slash Commands (Cursor Agent 对话中使用)
| 命令 | 用途 |
| :--- | :--- |
| `/speckit.constitution` | 创建项目原则 |
| `/speckit.specify` | 创建需求规范 |
| `/speckit.clarify` | 澄清规范问题 |
| `/speckit.plan` | 创建技术方案 |
| `/speckit.tasks` | 拆分任务 |
| `/speckit.implement` | 执行实现 |
| `/speckit.checklist` | 验证规范依从性 |
| `/speckit.analyze` | 分析代码质量 |

---

## 🎯 总结

### Spec Kit + Cursor 的完美配合

1. **Spec Kit** 提供结构和规范
   - 标准化文档框架 (Slash Commands)
   - 开发流程引导 (SDD Workflow)
   - 自动验证机制

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

**Spec Kit + Cursor = 规范化 + 高效率！** 🚀
