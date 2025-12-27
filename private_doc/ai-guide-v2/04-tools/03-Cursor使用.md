# Cursor使用指南

> IDE集成，快速开发

---

## 🎯 为什么选Cursor？

- ⭐⭐⭐ 开发速度最快
- ⭐⭐⭐ IDE深度集成
- ⭐⭐⭐ 学习成本最低

**最适合**：日常快速开发、代码重构

---

## ⚙️ 配置

### 1. 安装Cursor

下载：https://cursor.sh/

### 2. 配置.cursorrules

项目根目录创建`.cursorrules`：

```markdown
# IDRM Project Rules

## 5-Phase Workflow
Context → Specify → Design → Tasks → Implement

## Architecture
Handler → Logic → Model

## Standards
- Chinese comments
- Functions < 50 lines
- Error wrapping %w

## Specs
private_doc/spec/
```

---

## 🚀 核心功能

### 1. @-mentions（最重要）

**引用规范**：
```
@private_doc/spec/core/workflow.md
@private_doc/spec/architecture/layered-architecture.md

请遵循5阶段工作流和分层架构
```

**引用代码**：
```
@api/internal/handler/category/createcategoryhandler.go

参考这个handler，创建UpdateCategoryHandler
```

### 2. Cmd+K（快速编辑）

选中代码 → `Cmd+K` → 输入指令：
```
添加中文注释
重构这个函数，拆分为更小的函数
添加错误处理
```

### 3. Cmd+L（对话模式）

适合复杂任务：
```
Phase 1: 生成requirements
Phase 2: 生成design
Phase 4: 生成代码
```

### 4. Composer（多文件编辑）

同时编辑多个文件：
1. `Cmd+I`打开Composer
2. 添加多个文件
3. 一次性生成/修改

---

## 💡 最佳实践

### 1. 逐阶段开发

```
# Step 1
@spec/core/workflow.md
添加xxx功能
Phase 1: 生成Specify

# Step 2  
确认Specify
Phase 2: 生成Design

# Step 3
确认Design
Phase 4: 生成代码
```

### 2. 充分引用规范

每次对话都引用相关规范：
```
@spec/architecture/layered-architecture.md
@spec/coding-standards/go-style-guide.md
@spec/coding-standards/error-handling.md
```

### 3. 明确要求

```
要求：
1. 完整的中文注释
2. 函数<50行
3. 错误处理用%w
4. 遵循分层架构
```

---

## 📖 使用示例

### 示例1：添加API

```
@spec/core/workflow.md
@spec/architecture/layered-architecture.md

添加接口：GET /api/v1/categories/:id

Phase 1: 生成Specify（EARS格式）

[Cursor生成requirements]

确认
Phase 2: 生成Design

[Cursor生成design]

确认
Phase 4: 生成代码
- Handler: getcategoryhandler.go
- Logic: getcategorylogic.go
- 中文注释
- 错误处理
```

### 示例2：重构代码

选中函数 → `Cmd+K`：
```
重构这个函数：
1. 拆分为更小的函数
2. 添加中文注释
3. 优化错误处理

参考规范：
@spec/coding-standards/go-style-guide.md
```

---

## ⚡ 快捷键

- `Cmd+K` - 快速编辑
- `Cmd+L` - 对话模式
- `Cmd+I` - Composer
- `Cmd+Shift+P` - 命令面板

---

## 🔧 技巧

1. **保存上下文** - 同一个Chat持续对话
2. **分步执行** - 不要一次要求太多
3. **及时Review** - 应用前先Review代码
4. **版本控制** - 频繁commit
5. **善用@** - 引用比描述更准确

---

详细指南：`private_doc/Cursor使用指南.md`

**Cursor，日常开发首选！** ⚡
