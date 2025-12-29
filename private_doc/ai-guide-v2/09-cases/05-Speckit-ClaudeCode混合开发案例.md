# 实战案例：Speckit + Claude Code 混合开发

> 规范驱动 + 命令行AI - 极速开发工作流

---

## 📋 案例概述

**功能**：数据标签（Tag）管理  
**复杂度**：中等（约400行）  
**工具**：Speckit CLI + Claude Code (CLI)  
**耗时**：0.8个工作日  
**适用场景**：自动化程度高、喜欢命令行操作的开发者

---

## 💡 为什么选择 Speckit + Claude Code？

### 强强联合
- **Speckit**：负责"做什么"（What）和"标准是什么"（Rules）。它提供结构、模板和验证。
- **Claude Code**：负责"怎么做"（How）。它作为智能Agent，不仅能写代码，还能**直接执行终端命令**（包括Speckit命令）。

### 核心优势：闭环自动化
与IDE插件不同，Claude Code 运行在终端中，可以形成完美的自动化闭环：
1. Claude 读取 Speckit 规范
2. Claude 生成代码
3. Claude 运行 `speckit validate` 检查自身工作
4. 发现问题自动修复，直到通过验证

---

## 🛠️ 工具准备

### 1. 环境确认

```bash
# 确认 Speckit 安装
speckit --version

# 确认 Claude Code 登录
claude login
```

### 2. 初始化项目

```bash
cd /path/to/idrm
speckit init
```

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
1. 阅读 @CLAUDE.md 和 @private_doc/spec/core/workflow.md 理解工作流
2. 运行 `speckit create feature tag-management` 创建规范目录
3. 查看生成的文件结构
```

**Claude 的执行过程**：
1. 读取指定文档。
2. 执行 `speckit create feature tag-management`。
3. 也就是创建了 `specs/features/tag-management/` 目录。

---

## 📋 Phase 1: Specify (15分钟)

### Step 1: 生成并完善需求

**Prompt 2 (Requirements)**:

```text
现在进入 Phase 1: Specify。

1. 运行 `speckit generate requirements tag-management` 生成模板。
2. 读取生成的 `requirements.md`。
3. 根据以下简述完善该文件（直接编辑）：
   - 功能：创建/删除标签、为资源打标签、按标签查询、标签统计
   - User Stories: 使用 AS/I WANT/SO THAT 格式
   - Acceptance Criteria: 必须使用 EARS 格式 (WHEN... THE SYSTEM SHALL...)
   - Business Rules: 标签名唯一且不为空，颜色为HEX格式
   - Data Considerations: 软删除，级联解除关联

4. 编辑完成后，**务必运行** `speckit validate requirements tag-management` 验证你的工作。如果失败，请自动修正直到通过。
```

**Claude 的执行过程**：
- 生成模板。
- 编辑文件填充内容。
- **自动运行验证命令**。
- 如果 Speckit 报错（例如缺少 EARS 格式），Claude 会读取错误日志 -> 修改 markdown -> 再次验证，直到屏幕显示 `✓ Validation Passed`。

### Step 2: 提交

**Prompt 3**:
```text
验证通过了。请运行 `speckit mark complete requirements tag-management`，然后 git commit 提交变更。
```

---

## 🎨 Phase 2: Design (25分钟)

### Step 1: 架构设计与验证

**Prompt 4 (Design)**:

```text
进入 Phase 2: Design。

1. 运行 `speckit generate design tag-management --from-requirements`。
2. 读取 `requirements.md` 和生成的 `design.md`。
3. 完善 `design.md`，要求：
   - 遵循 @layered-architecture.md (Handler->Logic->Model)
   - 数据库设计：`tags` 表和 `resource_tags` 关联表，包含完整索引设计
   - Model接口：定义 Go interface
   - 序列图：用 Mermaid 描述"为资源打标签"的流程
   
4. 完成后，运行 `speckit validate design tag-management`。
5. 同样，如果验证失败，请自动修正。
```

**关键点**：Claude Code 能很好地处理 Mermaid 语法和 SQL Schema 设计。Speckit 的验证器会检查是否遗漏了必要的章节（如 database schema）。

### Step 2: 提交

**Prompt 5**:
```text
运行 `speckit mark complete design tag-management` 并提交。
```

---

## 📋 Phase 3: Tasks (10分钟)

### Step 1: 任务拆分

**Prompt 6 (Tasks)**:

```text
进入 Phase 3: Tasks。

请读取 `design.md`，并在 `specs/features/tag-management/tasks.md` 中拆分开发任务。
要求：
1. 每个任务代码行数 < 50行
2. 明确依赖关系 (Model -> Logic -> Handler)
3. 包含详细验收标准

完成后，无需Speckit验证，直接展示给我看。
```

---

## 💻 Phase 4: Implement (3小时)

这里是 Claude Code 发挥最大威力的地方。我们可以让它**逐个任务**执行，并**自动测试**。

### Step 1: Model 层开发

**Prompt 7 (Model Implementation)**:

```text
开始 Phase 4。请实现 Model 层任务。

1. 读取 `design.md` 和 `tasks.md`。
2. 创建并编写 Model 层代码 (interface, types, gorm_dao)。
   - 位置：`model/resource_catalog/tag/`
   - 确保包含中文注释
3. 编写完成后，**创建一个临时的 main_test.go** 或运行 `go build ./model/resource_catalog/tag/...` 确保编译通过。
4. 运行 `speckit track tag-management --task "Model Layer" --status completed`。
```

### Step 2: Logic 层开发 (循环模式)

我们可以使用 Claude 的 `--loop` 模式或者在对话中批量指令。

**Prompt 8 (Logic Implementation)**:

```text
现在实现 Logic 层。请按以下顺序逐个文件实现，并为每个文件生成单元测试：

1. `createtaglogic.go`
2. `assigntaglogic.go`
3. `querybytaglogic.go`

对于每个文件：
1. 编写 Logic 代码（<50行，错误处理）。
2. 编写对应的 `_test.go` 文件（Table-driven test, Mock generated via go-zero or manual interface mock）。
3. **运行 `go test` 验证该文件**。如果测试失败，自动修复代码。
```

**Claude 的强大之处**：
它会在终端中实际运行 `go test`。如果失败，它会读取 panic 信息或错误输出，修改代码，再次运行，直到测试通过。这是真正的 TDD（Test-Driven Development）闭环。

### Step 3: Handler 与 API

**Prompt 9 (API)**:

```text
最后实现 Handler 层和 API 定义。

1. 更新 `api/doc/api.api` 添加 Tag 相关接口。
2. 运行 `goctl api go ...` 生成代码。
3. 填充 Handler 逻辑（调用 Logic）。
4. 运行 `go run api.go` 启动服务，并使用 `curl` 进行一次集成测试（创建一个标签），确保存活。
```

### Step 4: 最终质量门禁

**Prompt 10 (Final Check)**:

```text
开发完成。请运行项目级的质量检查：

1. `speckit check tag-management` (检查规范依从性)
2. `golangci-lint run ./...` (检查代码质量)
3. `go test ./...` (确保无回归)

如果发现任何 Lint 错误或测试失败，请修复它们。
```

---

## 📊 效率对比

| 开发阶段 | 传统方式 | Speckit + Cursor | **Speckit + Claude Code** |
| :--- | :--- | :--- | :--- |
| **Specify** | 手动编写检测 | AI辅助编写，插件检测 | **AI编写+AI自动运行检测+AI自动修复** (最快) |
| **Impl (Coding)** | 手写 | AI生成，手动运行测试 | **AI生成+AI自动运行测试+AI自动Debug** |
| **Verify** | 手动提交 | 手动操作 | **CLI一键完成** |
| **Context Switch** | 高 (文档<->IDE) | 中 (都在IDE) | **低 (全在终端对话)** |

## 🎓 最佳实践总结

1. **"Trust but Verify" Loop**：
   始终要求 Claude 在生成内容后**立即**运行验证命令（`speckit validate` 或 `go test`）。不要等到最后再验。

2. **Explicit Context**：
   在 Prompt 中显式指明 `@document`，虽然 Claude Context 很大，但这能提高准确度。

3. **Step-by-Step**：
   虽然 Claude 很强，但将 Model、Logic、Handler 分开 Prompt 效果最好，避免 token 输出截断或逻辑混淆。

4. **Self-Correction**：
   利用 Claude Code 读取终端输出的能力。如果命令报错，直接告诉它 "Fix it based on the error output"，它通常能自行解决。

---

**结论**：对于熟悉 CLI 的开发者，**Speckit + Claude Code** 是目前自动化程度最高的组合。它将"规范检查"这一耗时环节变成了 AI 自动修正的目标函数，极大提升了交付质量。
