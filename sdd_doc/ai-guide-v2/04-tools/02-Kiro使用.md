# Kiro.dev使用指南

> 结构化开发，团队协作

---

## 🎯 为什么选Kiro？

- ⭐⭐⭐ 结构化specs输出
- ⭐⭐⭐ 3-phase自动生成
- ⭐⭐⭐ 任务可视化追踪

**最适合**：大型功能开发、团队协作

---

## ⚙️ 配置

### 1. 配置Steering

创建`.kiro/steering/idrm-rules.md`：

```markdown
# IDRM Steering

## Workflow
5-Phase: Context → Specify (EARS) → Design → Tasks → Implement

## Architecture
Layered: Handler → Logic → Model

## Standards
- Functions < 50 lines
- Chinese comments
- Error wrapping %w

## Quality
- Test >80%
- Lint clean

## Specs
参考：sdd_doc/spec/
```

---

## 🚀 使用流程

### 1. 创建Spec

```
Kiro面板 → + → New Spec
名称：Feature Name
描述：功能描述
```

### 2. Requirements阶段

Kiro自动引导生成：
- User Stories
- Acceptance Criteria (EARS)
- Technical Constraints
- Data Model

### 3. Design阶段

Kiro自动生成：
- Architecture Overview
- File Structure
- Interface Definitions
- Sequence Diagrams

### 4. Implementation阶段

Kiro生成：
- Tasks列表
- 每个task可点击执行
- 自动追踪进度

---

## 📖 实战示例

### 添加Category管理功能

**Step 1: 创建Spec**
```
名称：Resource Category Management
```

**Step 2: Requirements**

Kiro会问：
```
这个功能是做什么的？
```

回答：
```
资源分类管理，包括：
- 创建、查询、更新、删除分类
- 支持父子关系（最多3层）
- 遵循IDRM规范
```

Kiro生成完整的requirements.md

**Step 3: Design**

Kiro自动生成design.md，包含：
- 分层架构设计
- 文件清单
- 接口定义

**Step 4: Tasks**

Kiro拆分为可执行任务：
- Task 1: Model接口
- Task 2: GORM实现
- Task 3: SQLx实现
- ...

点击Task执行，Kiro生成代码。

---

## 💡 最佳实践

### 1. 充分利用Steering

在Steering中引用所有规范，让Kiro理解项目标准。

### 2. 详细描述功能

Requirements阶段提供足够细节：
- 功能列表
- 数据结构
- 约束条件

### 3. Review生成内容

每个阶段都Review：
- Requirements完整吗？
- Design符合规范吗？
- Tasks合理吗？

### 4. 导出Specs

```
Kiro → Export → Markdown
```

保存到项目中，供团队查看。

---

## 🔄 Kiro vs Cursor

| 维度 | Kiro | Cursor |
|------|------|--------|
| 结构化 | ⭐⭐⭐ | ⭐⭐ |
| 速度 | ⭐⭐ | ⭐⭐⭐ |
| 团队协作 | ⭐⭐⭐ | ⭐ |
| 学习成本 | 中 | 低 |

**建议**：
- 大功能 → Kiro
- 小功能 → Cursor

---

## 🔧 常见问题

### Q: Kiro生成的代码质量如何？
**A**: 需要Review，但结构化程度高

### Q: 如何与团队共享？
**A**: 导出specs，提交到Git

### Q: 可以修改生成内容吗？
**A**: 可以手动编辑导出的markdown

---

详细指南：`sdd_doc/Kirodev使用指南.md`

**Kiro，大型功能首选！** 🏗️
