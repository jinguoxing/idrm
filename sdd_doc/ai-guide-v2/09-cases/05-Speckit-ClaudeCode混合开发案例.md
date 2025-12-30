# 实战案例：Spec Kit + Claude Code 混合开发

> 规范驱动 + 命令行AI - 极速开发工作流

---

## 📋 案例概述

**功能**：数据标签（Tag）管理  
**复杂度**：中等（约400行）  
**工具**：[GitHub Spec Kit](https://github.com/github/spec-kit) + Claude Code (CLI)  
**耗时**：0.8个工作日  
**适用场景**：自动化程度高、喜欢命令行操作的开发者

---

## 💡 为什么选择 Spec Kit + Claude Code？

### 强强联合
- **Spec Kit + IDRM 模板**：提供结构、规范和验证（做什么 + 标准）
- **Claude Code**：负责执行（怎么做）。它不仅能写代码，还能**直接执行终端命令**

### 核心优势：闭环自动化
1. Claude 读取 `.specify/` 规范和模板
2. Claude 生成符合规范的代码
3. Claude 使用 `/speckit.checklist` 检查自身工作
4. 发现问题自动修复，直到通过验证

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

## 📝 Phase 0: Context (10分钟)

### Step 1: 让 Claude 理解项目上下文

```bash
$ claude
```

**Prompt 1 (Context)**:

```text
/init
我要在这个项目中开发"数据标签管理"功能。

请先阅读以下文档理解项目规范：
@CLAUDE.md
@.specify/memory/constitution.md
@sdd_doc/spec/core/workflow.md
@sdd_doc/spec/architecture/layered-architecture.md

总结关键规范要求。
```

**Claude 的执行过程**：
1. 读取 IDRM 宪章和规范文档
2. 理解分层架构和编码规范
3. 向用户汇报理解的内容

---

## 📋 Phase 1: Specify (15分钟)

### Step 1: 使用 /speckit.specify 创建需求规范

**Prompt 2 (Requirements)**:

```text
现在进入 Phase 1: Specify。

请使用 /speckit.specify 命令创建需求规范：

/speckit.specify 数据标签管理功能

参考模板：@.specify/templates/spec-template.md
参考引导：@.github/prompts/specify.prompt.md

功能需求：
- 创建/删除标签
- 为资源打标签
- 按标签查询资源
- 标签统计

要求：
- User Stories: 使用 AS/I WANT/SO THAT 格式
- Acceptance Criteria: 使用 EARS 格式 (WHEN... THE SYSTEM SHALL...)
- Business Rules: 标签名唯一且不为空，颜色为HEX格式
- Data Considerations: 软删除，级联解除关联
```

**输出文件**：`specs/tag-management/spec.md`

### Step 2: 验证规范

**Prompt 3**:
```text
请使用 /speckit.clarify 检查规范是否有遗漏或模糊的地方。
如果有问题，请自动修正。
```

---

## 🎨 Phase 2: Design (25分钟)

### Step 1: 使用 /speckit.plan 创建技术方案

**Prompt 4 (Design)**:

```text
进入 Phase 2: Design。

请使用 /speckit.plan 命令创建技术方案：

/speckit.plan

参考模板：@.specify/templates/plan-template.md
参考引导：@.github/prompts/plan.prompt.md
参考规范：
@sdd_doc/spec/architecture/layered-architecture.md
@sdd_doc/spec/architecture/dual-orm-pattern.md

技术要求：
- Go-Zero 微服务框架
- 遵循 Handler→Logic→Model 分层
- 数据库：`tags` 表和 `resource_tags` 关联表
- API：go-zero .api 格式
- ORM：GORM
- 序列图：Mermaid 格式
```

**输出文件**：`specs/tag-management/plan.md`

### Step 2: 生成 API 和 DDL 文件

**Prompt 5**:
```text
基于 plan.md，请生成：
1. `api/doc/resource_catalog/tag.api` - 参考 @.specify/templates/api-template.api
2. `migrations/resource_catalog/tags.sql` - 参考 @.specify/templates/schema-template.sql

然后 git commit 提交变更。
```

---

## 📋 Phase 3: Tasks (10分钟)

### Step 1: 使用 /speckit.tasks 拆分任务

**Prompt 6 (Tasks)**:

```text
进入 Phase 3: Tasks。

请使用 /speckit.tasks 命令拆分开发任务：

/speckit.tasks

参考模板：@.specify/templates/tasks-template.md
参考引导：@.github/prompts/tasks.prompt.md

要求：
- 每个任务代码行数 < 50行
- 明确依赖关系 (Model → Logic → Handler)
- 包含详细验收标准
```

**输出文件**：`specs/tag-management/tasks.md`

---

## 💻 Phase 4: Implement (3小时)

这是 Claude Code 发挥最大威力的地方。

### Step 1: 使用 /speckit.implement 执行实现

**Prompt 7 (Implementation)**:

```text
开始 Phase 4。

请使用 /speckit.implement 按照 tasks.md 中的任务列表逐个实现：

/speckit.implement

对于每个任务：
1. 生成代码框架（goctl api go / goctl model mysql ddl）
2. 编写业务逻辑
3. 编写单元测试
4. 运行 `go test` 验证
5. 如果测试失败，自动修复代码
```

### Step 2: Model 层开发

**Prompt 8 (Model Implementation)**:

```text
请实现 Model 层任务。

1. 运行 `goctl api go -api api/doc/resource_catalog/tag.api -dir api/ --style=goZero`
2. 运行 `goctl model mysql ddl -src migrations/resource_catalog/tags.sql -dir model/resource_catalog/tag/ --style=goZero`
3. 创建 Model 层代码 (interface, types, gorm_dao)
   - 位置：`model/resource_catalog/tag/`
   - 确保包含中文注释
4. 运行 `go build ./model/resource_catalog/tag/...` 确保编译通过
```

### Step 3: Logic 层开发 (循环模式)

**Prompt 9 (Logic Implementation)**:

```text
现在实现 Logic 层。请按以下顺序逐个文件实现：

1. `createtaglogic.go`
2. `assigntaglogic.go`
3. `querybytaglogic.go`

对于每个文件：
1. 编写 Logic 代码（<50行，错误处理）
2. 编写对应的 `_test.go` 文件（Table-driven test）
3. **运行 `go test` 验证**。如果测试失败，自动修复代码
```

**Claude 的强大之处**：它会在终端中实际运行 `go test`。如果失败，它会读取错误输出，修改代码，再次运行，直到测试通过。这是真正的 TDD 闭环。

### Step 4: Handler 与 API

**Prompt 10 (API)**:

```text
最后实现 Handler 层和 API 定义。

1. 更新 `api/doc/api.api` 添加 Tag 相关接口
2. 运行 `goctl api go ...` 生成代码
3. 填充 Handler 逻辑（调用 Logic）
4. 运行 `go run api.go` 启动服务，并使用 `curl` 进行一次集成测试
```

### Step 5: 最终质量门禁

**Prompt 11 (Final Check)**:

```text
开发完成。请运行项目级的质量检查：

1. 使用 /speckit.checklist 检查规范依从性
2. `golangci-lint run ./...` (检查代码质量)
3. `go test ./...` (确保无回归)

如果发现任何 Lint 错误或测试失败，请修复它们。
```

---

## 📊 效率对比

| 开发阶段 | 传统方式 | Spec Kit + Cursor | **Spec Kit + Claude Code** |
| :--- | :--- | :--- | :--- |
| **Specify** | 手动编写 | AI辅助编写 | **AI编写+AI自动检测+AI自动修复** |
| **Impl** | 手写 | AI生成，手动测试 | **AI生成+AI自动测试+AI自动Debug** |
| **Verify** | 手动 | 手动操作 | **CLI一键完成** |
| **Context Switch** | 高 | 中 | **低 (全在终端)** |

---

## 🎓 最佳实践总结

1. **"Trust but Verify" Loop**：始终要求 Claude 在生成内容后**立即**运行验证命令

2. **Explicit Context**：在 Prompt 中显式指明 `@.specify/templates/...`

3. **Step-by-Step**：将 Model、Logic、Handler 分开 Prompt，避免输出截断

4. **Self-Correction**：利用 Claude Code 读取终端输出的能力自动修复

---

## 📚 Spec Kit 命令参考

### Slash Commands (AI 对话中使用)
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

**结论**：**Spec Kit + IDRM 模板 + Claude Code** 是目前自动化程度最高的组合。它将规范检查变成了 AI 自动修正的目标函数，极大提升了交付质量。

**官方文档**：[github/spec-kit](https://github.com/github/spec-kit)
