# Claude Skills和Sub-agents配置

本目录包含IDRM项目的Claude Code Skills和Sub-agents配置。

## 📁 目录结构

```
.claude/
├── config.json              # 主配置文件
├── skills/                  # Skills定义
│   ├── idrm-specify.json   # Phase 1: 生成业务需求
│   ├── idrm-design.json    # Phase 2: 生成技术设计
│   ├── idrm-tasks.json     # Phase 3: 拆分任务
│   ├── idrm-implement.json # Phase 4: 实施代码
│   └── idrm-review.json    # Code Review
└── agents/                  # Sub-agents定义
    ├── spec-agent.json           # 需求分析专家
    ├── design-agent.json         # 解决方案架构师
    ├── implementation-agent.json # Go开发专家
    └── review-agent.json         # 质量保证专家
```

## 🎯 Skills说明

### 1. idrm-specify (Phase 1)
- **功能**: 生成业务需求规范
- **输入**: 功能描述
- **输出**: `specs/features/{feature}/requirements.md`
- **特点**: 聚焦What和Why，不涉及技术实现

### 2. idrm-design (Phase 2)
- **功能**: 生成技术设计方案
- **输入**: Phase 1的requirements.md
- **输出**: `specs/features/{feature}/design.md`
- **特点**: 完整技术方案，包含架构、约束、数据模型

### 3. idrm-tasks (Phase 3)
- **功能**: 拆分为可执行任务
- **输入**: Phase 2的design.md
- **输出**: `specs/features/{feature}/tasks.md`
- **特点**: 每个task <50行，明确依赖

### 4. idrm-implement (Phase 4)
- **功能**: 实施代码
- **输入**: Task ID + tasks.md
- **输出**: 代码文件 + 测试文件
- **特点**: 遵循IDRM编码规范，自动化质量检查

### 5. idrm-review
- **功能**: Code Review
- **输入**: 要review的文件列表
- **输出**: review-report.md
- **特点**: 全面检查架构、代码质量、测试

## 🤖 Sub-agents说明

### 1. spec-agent (需求分析专家)
- **职责**: 生成业务需求规范
- **使用skill**: idrm-specify
- **特点**: 用非技术语言描述需求

### 2. design-agent (解决方案架构师)
- **职责**: 设计技术方案
- **使用skill**: idrm-design
- **特点**: 技术决策有明确理由

### 3. implementation-agent (Go开发专家)
- **职责**: 任务拆分和代码实施
- **使用skill**: idrm-tasks, idrm-implement
- **特点**: 严格遵循编码规范

### 4. review-agent (质量保证专家)
- **职责**: Code Review
- **使用skill**: idrm-review
- **特点**: 分级反馈(MUST/SHOULD/COULD)

## 🚀 使用方式

### 方式1: 完整工作流

```bash
# 自动执行5阶段工作流
claude --config .claude/config.json \\
       --workflow full-feature \\
       "Feature: 用户标签管理"
```

### 方式2: 单个Skill

```bash
# 只生成Requirements
claude --skill idrm-specify "Feature: 数据导出"

# 只生成Design
claude --skill idrm-design \\
       --input specs/features/data-export/requirements.md

# 只Review代码
claude --skill idrm-review \\
       --files "api/internal/handler/**/*.go"
```

### 方式3: 使用Sub-agent

```bash
# 使用Spec Agent
claude --agent spec-agent "我要添加分类管理功能"

# 使用Design Agent
claude --agent design-agent \\
       --input specs/features/category/requirements.md

# 使用Review Agent
claude --agent review-agent \\
       --files model/category/*.go
```

## 📊 质量门禁

### Gate 1 (Phase 1完成后)
- User Stories完整
- EARS notation正确
- Business Rules清晰
- 无技术实现细节

### Gate 2 (Phase 2完成后)
- 遵循分层架构
- 序列图在约束之前
- 完整Technical Constraints
- 完整Data Model
- Business Rules已映射

### Gate 3 (Phase 3完成后)
- 每个task <50行
- 依赖关系明确
- 自下而上顺序

### Gate 4 (Phase 4完成后)
- go build通过
- go test通过
- golangci-lint通过
- 覆盖率>80%

## 📚 参考文档

所有skills和agents都基于以下规范：

- `private_doc/spec/workflow/` - 5阶段工作流规范
- `private_doc/spec/architecture/` - 架构规范
- `private_doc/spec/coding-standards/` - 编码规范
- `private_doc/spec/quality/` - 质量规范
- `private_doc/ai-guide-v2/` - AI编程指南

## ⚙️ 配置说明

### config.json
主配置文件，定义：
- 项目信息
- Spec版本(v2.1.0)
- 默认工作流
- 质量门禁标准
- Spec文档引用

### Skills JSON结构
```json
{
  "name": "skill-name",
  "version": "2.1.0",
  "description": "...",
  "inputs": {...},
  "outputs": {...},
  "prompt_template": {...},
  "validation_rules": [...]
}
```

### Agents JSON结构
```json
{
  "name": "agent-name",
  "role": "...",
  "capabilities": {...},
  "personality": {...},
  "workflow": [...],
  "output_template": "..."
}
```

## 🔄 版本管理

- **当前版本**: v2.1.0
- **对齐标准**: GitHub Spec Kit + IDRM Spec v2.1.0
- **更新日志**: 见各文件的version字段

## 📝 注意事项

1. **Skills和Agents的关系**：Agents使用Skills执行任务
2. **配置修改**：修改后需重启Claude Code
3. **Spec更新**：Spec规范更新时，需同步更新这里的配置
4. **版本一致性**：所有配置文件version保持一致

---

**Created**: 2025-12-28  
**Version**: 2.1.0  
**Maintained by**: IDRM Team
