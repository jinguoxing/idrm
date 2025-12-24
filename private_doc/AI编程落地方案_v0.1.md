# AI编程落地实施方案

基于 GitHub Spec-Kit + Claude Code CLI

## 📋 概述

### 什么是 Spec-Kit？

GitHub Spec-Kit 是一个**规范驱动开发（Spec-Driven Development）**工具包，通过4个门控阶段来确保AI生成的代码质量：

```
Specify → Plan → Tasks → Implement
(定义)   (规划)  (任务)   (实施)
```

核心理念：**规范是可执行的产物**，而非事后文档。

### 为什么需要 Spec-Kit？

**传统AI编程问题**：
- ❌ "Vibe Coding" - 基于模糊提示生成代码
- ❌ AI经常生成与意图不符的代码
- ❌ 需要大量调试和修正
- ❌ 缺乏一致性和上下文持久性

**Spec-Kit解决方案**：
- ✅ 规范先行 - 先定义"做什么"和"为什么"
- ✅ 分阶段验证 - 每个阶段需要确认后才进入下一阶段  
- ✅ 持久化上下文 - AI始终了解项目全貌
- ✅ 任务拆分 - 小而可测试的增量开发

---

## 🎯 在 IDRM 项目中的应用

### 目标

1. **建立规范驱动开发流程**
2. **为AI提供完整的项目上下文**
3. **标准化开发工作流**
4. **提高代码质量和一致性**

### 适用场景

```
✅ 新功能开发（如data_view模块）
✅ 复杂重构（如刚完成的model层重构）
✅ 模块扩展（如添加新的RPC服务）
✅ 架构调整（如引入新中间件）
```

---

## 🏗️ 实施方案

### 阶段一：基础设施搭建

#### 1.1 创建目录结构

```bash
idrm/
├── .github/
│   └── spec/                      # 规范驱动开发目录
│       ├── PRINCIPLES.md          # 项目原则
│       ├── ARCHITECTURE.md        # 架构规范
│       ├── TECH_STACK.md          # 技术栈说明
│       ├── CODING_STYLE.md        # 代码风格指南
│       ├── PATTERNS.md            # 设计模式
│       └── AI_GUIDELINES.md       # AI编程指南
├── .speckit/                      # Spec-Kit 配置
│   ├── prompts/                   # AI提示词模板
│   │   ├── specify.md
│   │   ├── plan.md
│   │   ├── tasks.md
│   │   └── implement.md
│   └── templates/                 # 文档模板
│       ├── spec_template.md
│       ├── plan_template.md
│       └── task_template.md
└── specs/                         # 具体功能规范
    ├── features/
    │   ├── data_view.md
    │   └── data_understanding.md
    └── refactors/
        └── model_layer_2024.md
```

#### 1.2 安装Specify CLI

```bash
# 使用 npx (推荐)
npx @github/specify init

# 或全局安装
npm install -g @github/specify
specify init
```

---

### 阶段二：编写核心规范文档

#### 2.1 项目原则 (PRINCIPLES.md)

```markdown
# IDRM项目原则

## 核心价值观
- 代码质量优于开发速度
- 可维护性优于便利性
- 明确性优于简洁性

## 技术原则
1. **双ORM支持** - GORM优先，SQLx降级
2. **分层架构** - API/Logic/Model严格分离
3. **接口抽象** - 面向接口编程
4. **完整可观测性** - Logging + Tracing + Metrics

## AI编程原则
1. **规范先行** - 先写spec，后生成代码
2. **小步迭代** - 每次只做一件事
3. **持续验证** - 每个阶段都要review
4. **文档同步** - 代码和文档同步更新
```

#### 2.2 架构规范 (ARCHITECTURE.md)

```markdown
# IDRM架构规范

## 分层架构

\`\`\`
┌─────────────────────────────────┐
│     API Layer (Go-Zero)         │
│  Handler → Logic → Model        │
└─────────────────────────────────┘
             ↓
┌─────────────────────────────────┐
│      Model Layer (Dual ORM)     │
│  Interface → Factory → Impl     │
└─────────────────────────────────┘
             ↓
┌─────────────────────────────────┐
│     Infrastructure Layer        │
│  MySQL / Redis / Kafka          │
└─────────────────────────────────┘
\`\`\`

## Model层结构

每个表独立目录：
\`\`\`
model/resource_catalog/category/
├── interface.go    # Model接口
├── types.go        # 数据结构
├── factory.go      # ORM工厂
├── gorm_dao.go     # GORM实现
└── sqlx_model.go   # SQLx实现
\`\`\`

## 中间件栈

顺序：Recovery → RequestID → Trace → CORS → Logger
```

#### 2.3 代码风格指南 (CODING_STYLE.md)

```markdown
# IDRM代码风格指南

## 命名规范

### 文件命名
- API Handler: `{resource}{action}handler.go`
- Logic: `{resource}{action}logic.go`  
- Model DAO: `{table}_dao.go`
- Model  Model: `{table}_model.go`

### 包命名
- 全小写
- 按功能模块组织
- 表级别独立包 (如 `category`)

## 代码组织

### 导入顺序
1. 标准库
2. 第三方库
3. 项目内包

### 函数顺序
1. 构造函数
2. 公开方法
3. 私有方法
4. init函数

## 注释规范
- 所有公开类型/函数必须有注释
- 使用中文注释
- 复杂逻辑必须解释
```

---

### 阶段三：配置 Claude Code CLI

#### 3.1 创建 AI 指导文件

**.speckit/AI_GUIDELINES.md**:

```markdown
# AI编程指南

## 你的身份
你是IDRM项目的AI编程助手，熟悉Go-Zero框架和双ORM架构。

## 项目上下文

### 技术栈
- **框架**: Go-Zero (API), Gorm/SQLx (ORM)
- **数据库**: MySQL 8.0
- **中间件**: Redis, Kafka
- **可观测性**: Jaeger (Tracing), ELK (Logging)

### 项目结构
\`\`\`
api/     - API服务 (HTTP)
job/     - 定时任务
consumer/ - 消息消费
model/   - 数据模型
pkg/     - 公共包
\`\`\`

### 关键约束
1. **必须**使用双ORM模式
2. **必须**实现接口抽象
3. **必须**遵循分层架构
4. **必须**添加完整注释

## 开发流程

### 🔷 Specify阶段
1. 理解需求
2. 定义用户故事
3. 列出技术要点
4. 标注不确定项

### 🔷 Plan阶段
1. 设计技术方案
2. 选择技术栈
3. 定义数据模型
4. 规划文件结构

### 🔷 Tasks阶段
1. 拆分可测试任务
2. 按依赖排序
3. 每个任务<50行代码
4. 明确验收条件

### 🔷 Implement阶段
1. 一次执行一个任务
2. 生成完整代码
3. 添加单元测试
4. 更新文档

## 代码模板

### 新增Model
\`\`\`go
package {table}

// Model 定义接口
type Model interface {
    Insert(ctx context.Context, data *{Table}) (*{Table}, error)
    FindOne(ctx context.Context, id int64) (*{Table}, error)    
    // ... 其他方法
    WithTx(tx interface{}) Model
    Trans(ctx context.Context, fn func(ctx context.Context, model Model) error) error
}
\`\`\`

### 新增Logic
\`\`\`go
type {Action}{Resource}Logic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func New{Action}{Resource}Logic(ctx context.Context, svcCtx *svc.ServiceContext) *{Action}{Resource}Logic {
    return \u0026{Action}{Resource}Logic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *{Action}{Resource}Logic) {Action}{Resource}(req *types.{Req}) (resp *types.{Resp}, err error) {
    // 实现逻辑
    return
}
\`\`\`
```

#### 3.2 创建斜杠命令配置

**.speckit/prompts/commands.md**:

```markdown
# 可用斜杠命令

## /speckit.specify
生成功能规范

输入：功能描述
输出：结构化规范（包含用户故事、技术要点、不确定项）

## /speckit.plan  
生成技术方案

输入：功能规范
输出：技术方案（架构、文件列表、数据模型）

## /speckit.tasks
生成任务列表

输入：技术方案
输出：可执行任务清单（带序号、描述、验收）

## /speckit.implement
执行单个任务

输入：任务编号
输出：完整实现代码
```

---

### 阶段四：实践工作流

#### 4.1 示例：添加 Directory 表

##### Step 1: Specify

**命令**:
```bash
specify /speckit.specify "在resource_catalog模块添加directory表，支持目录的CRUD操作"
```

**生成的spec**:
```markdown
# Feature Spec: Resource Catalog - Directory

## 用户故事
作为系统管理员
我想要管理资源目录
以便组织和分类系统资源

## 功能需求
1. 创建目录
2. 查询目录列表
3. 更新目录信息
4. 删除目录
5. 按父目录查询子目录

## 技术要点
- 表结构：id, name, code, parent_id, level, status
- 支持树形结构
- 双ORM实现
- 完整CRUD接口

## 不确定项
- [ ] 是否需要缓存？
- [ ] 删除时是否级联？
- [ ] level最大深度限制？
```

##### Step 2: Plan

**命令**:
```bash
specify /speckit.plan
```

**生成的计划**:
```markdown
# Technical Plan: Directory Model

## 架构
遵循现有Model层结构，table-based directory pattern

## 文件列表
1. model/resource_catalog/directory/interface.go
2. model/resource_catalog/directory/types.go
3. model/resource_catalog/directory/vars.go
4. model/resource_catalog/directory/factory.go
5. model/resource_catalog/directory/gorm_dao.go
6. model/resource_catalog/directory/sqlx_model.go

## 数据模型
\`\`\`go
type Directory struct {
    Id        int64     `db:"id" gorm:"column:id;primaryKey"`
    Name      string    `db:"name" gorm:"column:name"`
    Code      string    `db:"code" gorm:"column:code;uniqueIndex"`
    ParentId  int64     `db:"parent_id" gorm:"column:parent_id;index"`
    Level     int       `db:"level" gorm:"column:level"`
    Status    int       `db:"status" gorm:"column:status"`
    CreatedAt time.Time `db:"created_at" gorm:"column:created_at"`
    UpdatedAt time.Time `db:"updated_at" gorm:"column:updated_at"`
}
\`\`\`

## Dependencies
- Gorm
- SQLx
- Go-Zero
```

##### Step 3: Tasks

**命令**:
```bash
specify /speckit.tasks
```

**生成的任务**:
```markdown
# Task Breakdown

## Task 1: Create directory and base files
- Create model/resource_catalog/directory/
- Create interface.go with Model interface
- Create types.go with Directory struct
- Create vars.go with constants

## Task 2: Implement factory
- Create factory.go
- Implement NewModel function
- Add factory registration logic

## Task 3: Implement GORM DAO
- Create gorm_dao.go
- Implement all Model interface methods
- Add init() for factory registration

## Task 4: Implement SQLx Model
- Create sqlx_model.go
- Implement all Model interface methods  
- Add init() for factory registration

## Task 5: Update ServiceContext
- Add DirectoryModel field
- Initialize in NewServiceContext
- Add blank import

## Task 6: Add unit tests
- Test GORM implementation
- Test SQLx implementation
- Test factory selection
```

##### Step 4: Implement

**命令**:
```bash
specify /speckit.implement 1
```

AI会生成Task 1的完整代码，然后逐个执行后续任务。

---

## 📊 实施路线图

### 第1周：基础设施
- [ ] 创建.github/spec/目录
- [ ] 编写核心规范文档
- [ ] 安装Specify CLI
- [ ] 配置AI指导文件

### 第2周：模板和工作流
- [ ] 创建规范模板
- [ ] 创建任务模板
- [ ] 测试完整工作流
- [ ] 编写使用教程

### 第3周：实际应用
- [ ] 用spec-kit开发新功能
- [ ] 记录问题和优化点
- [ ] 改进模板和指南
- [ ] 团队培训

---

## 🎯 成功指标

### 短期（1个月）
- ✅ 所有新功能都使用spec-first approach
- ✅ AI生成代码通过率 >80%
- ✅ 规范文档覆盖所有模块

### 长期（3个月）
- ✅ 开发效率提升 30%
- ✅ Bug率降低 25%
- ✅ 代码一致性显著提高
- ✅ 文档始终保持最新

---

## 🔧 工具集成

### Claude Code CLI
```bash
# 基本用法
claude "按照spec实现directory model"

# 使用上下文
claude --context=".github/spec/**/*.md" "实现新功能"

# 分阶段执行
claude /speckit.specify
claude /speckit.plan  
claude /speckit.tasks
claude /speckit.implement 1
```

### Git工作流
```bash
# 创建spec分支
git checkout -b spec/directory-model

# 提交spec
git add specs/features/directory.md
git commit -m "spec: add directory model specification"

# 提交plan
git add specs/features/directory-plan.md
git commit -m "plan: directory model technical plan"

# 提交实现
git add model/resource_catalog/directory/
git commit -m "feat: implement directory model"
```

---

## 💡 最佳实践

### DO ✅
1. **始终从spec开始** - 即使是小改动
2. **小步迭代** - 一次只做一个任务
3. **及时验证** - 每个阶段都要review
4. **保持文档更新** - spec是活文档

### DON'T ❌
1. **跳过spec直接coding** - 会导致返工
2. **任务太大** - 超过50行就要拆分
3. **忽略不确定项** - 必须先澄清
4. **文档与代码脱节** - 同步更新

---

## 📚 参考资源

- [GitHub Spec-Kit](https://github.com/github/spec-kit)
- [Specify CLI文档](https://github.com/github/specify-cli)
- [Claude Code CLI](https://docs.anthropic.com/claude/docs/code-cli)
- IDRM项目内部文档：
  - `/docs/architecture/分层架构.md`
  - `/model/README.md`
  - `/api/doc/STRUCTURE.md`

---

## 🎉 预期收益

### 开发体验改善
- 🚀 更快的功能开发
- 🎯 更准确的AI代码生成
- 📝 自动生成的文档
- 🔄 一致的代码风格

### 代码质量提升
- ✨ 更清晰的架构
- 🐛 更少的bug
- 📖 更好的可维护性
- 🔍 更易于code review

### 团队协作优化
- 👥 统一的开发流程
- 📋 清晰的任务分解
- 💬 更好的沟通（通过spec）
- 🎓 新人快速上手
