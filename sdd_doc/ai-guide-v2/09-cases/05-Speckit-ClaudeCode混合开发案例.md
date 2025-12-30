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
- **Spec Kit (Specify CLI)**：GitHub 官方规范驱动开发工具包，负责"做什么"（What）和"标准是什么"（Rules）。它提供结构、模板和验证。
- **Claude Code**：负责"怎么做"（How）。它作为智能Agent，不仅能写代码，还能**直接执行终端命令**。

### 核心优势：闭环自动化
与IDE插件不同，Claude Code 运行在终端中，可以形成完美的自动化闭环：
1. Claude 读取 Spec Kit 规范
2. Claude 生成代码
3. Claude 使用 `/speckit.checklist` 检查自身工作
4. 发现问题自动修复，直到通过验证

---

## 🛠️ 工具准备

### 1. 安装 Specify CLI

```bash
# 使用 uv 安装（推荐）
uv tool install specify-cli --from git+https://github.com/github/spec-kit.git

# 验证安装
specify check

# 确认 Claude Code 登录
claude login
```

### 2. 初始化项目

```bash
cd /path/to/idrm

# 初始化 Spec Kit（选择 Claude 作为 AI 助手）
specify init . --ai claude
```

这将在项目中创建 `.speckit/` 目录和相关配置文件。

---

## 📝 Phase 0: Context (10分钟)

### Step 1: 创建Feature并让Claude理解上下文

我们直接通过 Claude CLI 启动任务：

```bash
$ claude
```

**Prompt 1 (Context)**:

```text
/init
我要在这个项目中开发"数据标签管理"功能。

请先做以下准备：
1. 阅读 @CLAUDE.md 和 @sdd_doc/spec/core/workflow.md 理解工作流
2. 了解项目架构和编码规范
3. 准备好开发环境
```

**Claude 的执行过程**：
1. 读取指定文档
2. 理解项目规范
3. 向用户汇报理解的内容

---

## 📋 Phase 1: Specify (15分钟)

### Step 1: 使用 /speckit.specify 创建需求规范

**Prompt 2 (Requirements)**:

```text
现在进入 Phase 1: Specify。

请使用 /speckit.specify 命令创建需求规范：

/speckit.specify 数据标签管理功能，包含：
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

**Claude 的执行过程**：
- 执行 `/speckit.specify` 创建 `spec.md`
- 自动填充需求内容
- 生成结构化的规范文档

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

/speckit.plan 使用以下技术栈：
- Go-Zero 微服务框架
- 遵循 @sdd_doc/spec/architecture/layered-architecture.md (Handler→Logic→Model)
- 数据库设计：`tags` 表和 `resource_tags` 关联表，包含完整索引设计
- API 接口：使用 go-zero .api 格式定义
- 序列图：用 Mermaid 描述"为资源打标签"的流程
```

**关键点**：Claude Code 能很好地处理 Mermaid 语法和 SQL Schema 设计。`/speckit.plan` 会生成结构化的 `plan.md` 文件。

### Step 2: 生成 API 和 DDL 文件

**Prompt 5**:
```text
基于 plan.md，请生成：
1. `api/doc/resource_catalog/tag.api` - go-zero API 定义
2. `migrations/resource_catalog/tags.sql` - DDL 文件

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

要求：
- 每个任务代码行数 < 50行
- 明确依赖关系 (Model → Logic → Handler)
- 包含详细验收标准
```

---

## 💻 Phase 4: Implement (3小时)

这里是 Claude Code 发挥最大威力的地方。我们可以让它**逐个任务**执行，并**自动测试**。

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
3. 创建并编写 Model 层代码 (interface, types, gorm_dao)
   - 位置：`model/resource_catalog/tag/`
   - 确保包含中文注释
4. 运行 `go build ./model/resource_catalog/tag/...` 确保编译通过
```

### Step 3: Logic 层开发 (循环模式)

**Prompt 9 (Logic Implementation)**:

```text
现在实现 Logic 层。请按以下顺序逐个文件实现，并为每个文件生成单元测试：

1. `createtaglogic.go`
2. `assigntaglogic.go`
3. `querybytaglogic.go`

对于每个文件：
1. 编写 Logic 代码（<50行，错误处理）
2. 编写对应的 `_test.go` 文件（Table-driven test）
3. **运行 `go test` 验证该文件**。如果测试失败，自动修复代码
```

**Claude 的强大之处**：
它会在终端中实际运行 `go test`。如果失败，它会读取 panic 信息或错误输出，修改代码，再次运行，直到测试通过。这是真正的 TDD（Test-Driven Development）闭环。

### Step 4: Handler 与 API

**Prompt 10 (API)**:

```text
最后实现 Handler 层和 API 定义。

1. 更新 `api/doc/api.api` 添加 Tag 相关接口
2. 运行 `goctl api go ...` 生成代码
3. 填充 Handler 逻辑（调用 Logic）
4. 运行 `go run api.go` 启动服务，并使用 `curl` 进行一次集成测试（创建一个标签），确保存活
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
| **Specify** | 手动编写检测 | AI辅助编写，插件检测 | **AI编写+AI自动运行检测+AI自动修复** (最快) |
| **Impl (Coding)** | 手写 | AI生成，手动运行测试 | **AI生成+AI自动运行测试+AI自动Debug** |
| **Verify** | 手动提交 | 手动操作 | **CLI一键完成** |
| **Context Switch** | 高 (文档<->IDE) | 中 (都在IDE) | **低 (全在终端对话)** |

## 🎓 最佳实践总结

1. **"Trust but Verify" Loop**：
   始终要求 Claude 在生成内容后**立即**运行验证命令（`/speckit.checklist` 或 `go test`）。不要等到最后再验。

2. **Explicit Context**：
   在 Prompt 中显式指明 `@document`，虽然 Claude Context 很大，但这能提高准确度。

3. **Step-by-Step**：
   虽然 Claude 很强，但将 Model、Logic、Handler 分开 Prompt 效果最好，避免 token 输出截断或逻辑混淆。

4. **Self-Correction**：
   利用 Claude Code 读取终端输出的能力。如果命令报错，直接告诉它 "Fix it based on the error output"，它通常能自行解决。

---

## 📚 Spec Kit 命令参考

### Specify CLI 命令
| 命令 | 用途 |
| :--- | :--- |
| `specify init . --ai claude` | 初始化项目 |
| `specify check` | 检查环境 |

### Slash Commands (AI 对话中使用)
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

**结论**：对于熟悉 CLI 的开发者，**Spec Kit + Claude Code** 是目前自动化程度最高的组合。它将"规范检查"这一耗时环节变成了 AI 自动修正的目标函数，极大提升了交付质量。

**官方文档**：[github/spec-kit](https://github.com/github/spec-kit)
