# 实战案例：使用Cursor完成完整开发

> 从需求到上线 - 纯Cursor工作流

---

## 📋 案例概述

**功能**：数据标签（Tag）管理  
**复杂度**：中等（约400行）  
**工具**：Cursor  
**耗时**：1个工作日  
**适用场景**：日常开发、中小型功能

---

## 💡 为什么选择Cursor？

**适合Cursor的场景**：
- ✅ 功能规模 < 500行
- ✅ 需要快速迭代
- ✅ 单人开发
- ✅ 已有明确需求

**Cursor的优势**：
- 🚀 即时反馈，快速编码
- 💬 对话式交互，自然流畅
- 🔍 上下文理解，代码关联性强
- 🎯 适合边思考边实现

---

## 🎯 需求背景

数据管理员希望为数据资源打标签，便于分类和检索。

**核心需求**：
- 创建/删除标签
- 为资源打标签
- 按标签查询资源
- 标签统计

---

## 📝 Phase 0: Context (15分钟)

### Step 1: 阅读规范

**在Cursor中打开规范文件**：

```
Cmd+P → 输入 "project-charter"
Cmd+P → 输入 "workflow"
Cmd+P → 输入 "layered-architecture"
```

浏览关键内容，理解：
- ✅ 5阶段工作流
- ✅ 分层架构（Handler→Logic→Model）
- ✅ 双ORM模式
- ✅ 编码规范

### Step 2: 环境检查

**在Cursor终端执行**：

```bash
# 拉取最新代码
git pull origin main

# 编译检查
go build ./...

# 确认数据库
docker ps | grep mysql
```

### Step 3: 创建工作目录

**使用Cursor Chat**：

```
@Workspace

请帮我创建标签管理功能的规范文档目录：
specs/features/tag-management/

需要创建以下文件：
- requirements.md
- design.md
- tasks.md

每个文件先创建头部框架即可
```

**创建分支**：

```bash
git checkout -b feature/tag-management
```

---

## 📋 Phase 1: Specify (30分钟)

### Step 1: 编写User Stories

**Cursor Prompt**：

```
@private_doc/spec/workflow/phase1-specify.md

我要添加数据标签管理功能。

Phase 1: 请帮我在 specs/features/tag-management/requirements.md 中编写User Stories

功能需求：
1. 创建/删除标签
2. 为资源（数据目录）打标签/取消标签
3. 按标签查询资源
4. 标签使用统计

要求：
- 使用 AS a/I WANT/SO THAT 格式
- 至少3个user story
- 从数据管理员视角
```

**Cursor会生成**：

```markdown
# Requirements: Tag Management

## User Stories

### Story 1: 创建标签
AS a 数据管理员
I WANT 创建新标签用于资源分类
SO THAT 我可以有效组织和管理数据资源

### Story 2: 为资源打标签
AS a 数据管理员
I WANT 为数据资源添加或移除标签
SO THAT 我可以快速组织和查找相关数据

### Story 3: 按标签查询
AS a 数据使用者
I WANT 按标签查询数据资源
SO THAT 我可以快速找到需要的数据
```

---

### Step 2: 编写EARS验收标准

**Cursor Prompt**：

```
@private_doc/spec/workflow/ears-notation-guide.md
@specs/features/tag-management/requirements.md

继续编写 Acceptance Criteria，使用EARS notation

要求：
1. 覆盖正常流程（Happy Path）
2. 覆盖所有参数验证（名称为空、过长、重复）
3. 覆盖异常情况（资源不存在、标签不存在、重复关联）
4. 覆盖删除场景（级联删除）
5. 每条EARS必须可测试
6. **不包含任何技术实现细节**

格式：
WHEN [条件]
THE SYSTEM SHALL [行为]
```

**Cursor会生成完整的EARS**，检查并调整。

---

### Step 3: 编写Business Rules

**Cursor Prompt**：

```
@specs/features/tag-management/requirements.md

继续编写 Business Rules 部分

要求：
- 只描述业务规则，不涉及技术实现
- 包含：唯一性规则、长度限制、关联关系、删除规则、统计规则
- 清晰明确，易于理解

参考模板：
### 标签规则
- ...

### 删除规则
- ...

### 统计规则
- ...
```

---

### Step 4: 编写Data Considerations

**Cursor Prompt**：

```
@specs/features/tag-management/requirements.md

继续编写 Data Considerations 部分

要求：
- 描述需要持久化的数据（不是表结构）
- 说明数据关系
- 明确级联删除需求

注意：不要定义数据库表结构，那是Phase 2的内容
```

---

### Step 5: Gate 1 检查

**Cursor Prompt**：

```
@private_doc/spec/quality/quality-gates.md
@specs/features/tag-management/requirements.md

请检查这个 requirements.md 是否通过 Gate 1

检查项：
- [ ] User stories完整 (AS/I WANT/SO THAT)
- [ ] 使用EARS notation
- [ ] Business rules明确
- [ ] Data considerations清晰
- [ ] **没有技术实现细节**

请逐项检查并给出改进建议
```

修正后，在文档末尾添加：

```markdown
## Gate 1 检查

- [x] User stories完整
- [x] EARS格式正确
- [x] Business rules明确
- [x] Data considerations清晰
- [x] 没有技术实现细节

✅ Pass
```

**提交Phase 1**：

```bash
git add specs/features/tag-management/requirements.md
git commit -m "docs: add tag management requirements (Phase 1)"
```

---

## 🎨 Phase 2: Design (40分钟)

### Step 1: 架构设计

**Cursor Prompt**：

```
@private_doc/spec/architecture/layered-architecture.md
@specs/features/tag-management/requirements.md

Phase 2: 请在 specs/features/tag-management/design.md 中创建技术设计

要求：
1. 遵循分层架构 (Handler→Logic→Model)
2. 定义完整的文件结构
3. 设计Model接口
4. 画出序列图
5. 定义数据库表结构（带索引和注释）
6. 说明ORM选择（GORM vs SQLx）
```

---

### Step 2: 文件结构规划

**Cursor Prompt**：

```
@specs/features/tag-management/design.md

请详细规划文件结构，包括：

Model层: model/resource_catalog/tag/
- interface.go (Model接口定义)
- types.go (数据结构)
- gorm_dao.go (GORM实现)
- factory.go (工厂方法)

Logic层: api/internal/logic/resource_catalog/tag/
- createtaglogic.go
- deletetaglogic.go
- assigntaglogic.go
- removetaglogic.go
- querybytaglogic.go
- tagstatslogic.go

Handler层: api/internal/handler/resource_catalog/tag/
- (对应的handler文件)
```

---

### Step 3: 接口定义

**Cursor Prompt**：

```
@private_doc/spec/architecture/dual-orm-pattern.md
@specs/features/tag-management/design.md

在Design文档中定义Model接口

要求：
1. 定义完整的Model interface
2. 包含所有CRUD方法
3. 包含资源-标签关联方法
4. 包含统计方法
5. 每个方法添加中文注释
6. 使用context.Context
```

---

### Step 4: 数据库设计

**Cursor Prompt**：

```
@specs/features/tag-management/design.md

在Design文档中定义数据库Schema

要求：
1. tags表：id, name(UNIQUE), color, created_at, updated_at
2. resource_tags表：id, resource_id, tag_id, created_at
3. 添加必要的索引
4. 添加中文注释
5. 指定字符集utf8mb4
6. 添加UNIQUE约束防止重复关联
```

---

### Step 5: 序列图设计

**Cursor Prompt**：

```
@specs/features/tag-management/design.md

使用Mermaid格式画出"为资源打标签"的序列图

要求：
1. Client → Handler → Logic → Model
2. 包含参数验证
3. 包含错误处理
4. 展示完整流程
```

---

### Step 6: Technical Constraints

**Cursor Prompt**：

```
@specs/features/tag-management/design.md

添加 Technical Constraints 部分

包含：
- 架构要求（分层架构）
- 函数大小限制（<50行）
- 注释要求（中文）
- 测试覆盖率（>80%）
- ORM选择及理由（为什么用GORM？）
```

---

### Step 7: Gate 2 检查

**Cursor Prompt**：

```
@private_doc/spec/quality/quality-gates.md
@specs/features/tag-management/design.md

请检查这个 design.md 是否通过 Gate 2

检查项：
- [ ] 符合分层架构
- [ ] 文件清单完整
- [ ] 接口定义清晰
- [ ] 序列图完整
- [ ] 数据库设计合理

请逐项检查并给出改进建议
```

**提交Phase 2**：

```bash
git add specs/features/tag-management/design.md
git commit -m "docs: add tag management design (Phase 2)"
```

---

## 📋 Phase 3: Tasks (20分钟)

### Step 1: 任务拆分

**Cursor Prompt**：

```
@specs/features/tag-management/design.md

Phase 3: 请在 specs/features/tag-management/tasks.md 中拆分任务

要求：
1. 每个task预估代码行数 < 50行
2. 明确依赖关系
3. 按顺序：Model → Logic → Handler → Test
4. 每个task包含：
   - 名称
   - 预估行数
   - 依赖任务
   - 涉及文件
   - 验收标准
   - 状态（⏸️ Not Started）
```

**Cursor会生成类似**：

```markdown
# Tasks: Tag Management

## Task 1: 创建Tag Model接口和类型
**Lines**: 40
**Status**: ⏸️ Not Started
**Depends**: -
**Files**:
- model/resource_catalog/tag/interface.go
- model/resource_catalog/tag/types.go

**Acceptance**:
- [ ] Model接口定义完整
- [ ] Tag和ResourceTag结构体正确
- [ ] 中文注释完整

## Task 2: 实现GORM DAO
**Lines**: 80
**Status**: ⏸️ Not Started
**Depends**: Task 1
...
```

---

### Step 2: Gate 3 检查

**Cursor Prompt**：

```
@specs/features/tag-management/tasks.md

检查任务拆分是否合理：

- 每个task < 80行？
- 依赖关系清晰？
- 验收标准明确？
- 是否可以独立实施？
```

**提交Phase 3**：

```bash
git add specs/features/tag-management/tasks.md
git commit -m "docs: add tag management tasks (Phase 3)"
```

---

## 💻 Phase 4: Implement (4-6小时)

### Task 1-2: Model层实现

**Step 1: 创建接口和类型**

**Cursor Prompt**：

```
@specs/features/tag-management/tasks.md
@specs/features/tag-management/design.md
@private_doc/spec/coding-standards/go-style-guide.md

执行 Task 1: 创建Tag Model接口和类型

请创建以下文件：
1. model/resource_catalog/tag/interface.go
2. model/resource_catalog/tag/types.go

要求：
- 定义完整的Model接口（所有CRUD和关联方法）
- 定义Tag和ResourceTag结构体
- 添加GORM和db标签
- 完整的中文注释
- 函数签名包含context.Context
```

**Cursor会生成代码**，检查后更新task状态：

```markdown
## Task 1: 创建Tag Model接口和类型
**Status**: ✅ Done
```

---

**Step 2: 实现GORM DAO**

**Cursor Prompt**：

```
@model/resource_catalog/tag/interface.go
@specs/features/tag-management/tasks.md
@private_doc/spec/architecture/dual-orm-pattern.md

执行 Task 2: 实现GORM DAO

请创建 model/resource_catalog/tag/gorm_dao.go

要求：
- 实现所有Model接口方法
- 使用GORM
- 正确处理事务（AssignTag, RemoveTag）
- 完整错误处理（使用%w包装）
- 中文注释
- 每个函数 < 50行
```

**验证实现**：

```bash
# 在Cursor终端运行
go build ./model/resource_catalog/tag/
```

---

**Step 3: 创建Factory**

**Cursor Prompt**：

```
@model/resource_catalog/tag/interface.go
@model/resource_catalog/tag/gorm_dao.go

执行 Task 3: 创建Factory

请创建 model/resource_catalog/tag/factory.go

要求：
- NewModel(db *gorm.DB) Model 工厂函数
- 返回GORM实现
```

**提交Model层**：

```bash
git add model/resource_catalog/tag/
git commit -m "feat: implement tag model layer"
```

---

### Task 4-9: Logic层实现

**批量生成Logic**

**Cursor Prompt**（为每个Logic文件执行一次）：

```
@model/resource_catalog/tag/interface.go
@specs/features/tag-management/design.md
@private_doc/spec/coding-standards/go-style-guide.md

执行 Task 4: 创建标签Logic

请创建 api/internal/logic/resource_catalog/tag/createtaglogic.go

要求：
1. 定义CreateTagLogic结构体
2. 实现CreateTag方法
3. 业务逻辑：
   - 验证参数（名称不为空，长度1-50）
   - 检查名称是否重复
   - 调用Model.Insert保存
   - 返回标签ID
4. 完整错误处理
5. 中文注释
6. 函数 < 50行

参考imports:
```go
import (
    "context"
    "fmt"
    
    "github.com/zeromicro/go-zero/core/logx"
    
    "af_idrm/api/internal/svc"
    "af_idrm/api/internal/types"
    "af_idrm/model/resource_catalog/tag"
)
```
```

**重复上述过程为每个Logic**：
- createtaglogic.go
- deletetaglogic.go
- assigntaglogic.go
- removetaglogic.go
- querybytaglogic.go
- tagstatslogic.go

**测试Logic层**：

**Cursor Prompt**：

```
@api/internal/logic/resource_catalog/tag/createtaglogic.go
@private_doc/spec/coding-standards/testing-standards.md

为这个Logic生成单元测试

要求：
- 文件名：createtaglogic_test.go
- 表驱动测试
- Mock Model接口
- 覆盖所有分支（正常、名称为空、名称重复）
- 使用testify/assert
```

**提交Logic层**：

```bash
go test ./api/internal/logic/resource_catalog/tag/
git add api/internal/logic/resource_catalog/tag/
git commit -m "feat: implement tag logic layer with tests"
```

---

### Task 10-15: Handler层实现

**批量生成Handler**

**Cursor Prompt**：

```
@api/internal/logic/resource_catalog/tag/createtaglogic.go
@private_doc/spec/architecture/api-design-guide.md

执行 Task 10: 创建标签Handler

请创建 api/internal/handler/resource_catalog/tag/createtaghandler.go

要求：
1. 定义CreateTagHandler函数
2. 实现逻辑：
   - 解析请求参数（r.ParseForm或ShouldBind）
   - 参数验证
   - 调用Logic.CreateTag
   - 统一响应格式
3. 函数 < 30行
4. 中文注释

响应格式：
```go
httpx.OkJson(w, map[string]interface{}{
    "code": 0,
    "message": "success",
    "data": result,
})
```

错误格式：
```go
httpx.Error(w, err)
```
```

**更新API路由**：

**Cursor Prompt**：

```
@api/internal/handler/resource_catalog/tag/

请帮我在 api/doc/api.api 中定义标签管理的所有API路由

包括：
- POST   /api/v1/tags           (创建标签)
- DELETE /api/v1/tags/:id       (删除标签)
- POST   /api/v1/resources/:id/tags  (为资源打标签)
- DELETE /api/v1/resources/:id/tags/:tagId  (移除标签)
- GET    /api/v1/tags/:id/resources (按标签查询资源)
- GET    /api/v1/tags/stats     (标签统计)

使用go-zero的api语法
```

**重新生成代码**：

```bash
# Cursor终端执行
goctl api go -api api/doc/api.api -dir api/ --style=goZero
```

**手动测试API**：

**在Cursor终端启动服务**：

```bash
cd api
go run api.go
```

**另开终端测试**：

```bash
# 创建标签
curl -X POST http://localhost:8888/api/v1/tags \
  -H "Content-Type: application/json" \
  -d '{"name":"重要数据","color":"#FF0000"}'

# 为资源打标签
curl -X POST http://localhost:8888/api/v1/resources/1/tags \
  -H "Content-Type: application/json" \
  -d '{"tag_id":1}'

# 查询
curl http://localhost:8888/api/v1/tags/1/resources
```

**提交Handler层**：

```bash
git add api/
git commit -m "feat: implement tag handlers and API routes"
```

---

### 补充测试

**Cursor Prompt**：

```
@api/internal/logic/resource_catalog/tag/
@private_doc/spec/coding-standards/testing-standards.md

请为所有Logic生成完整的单元测试

要求：
- 表驱动测试
- Mock Model接口
- 覆盖率 > 80%
- 测试正常和异常流程
```

**检查覆盖率**：

```bash
go test -cover ./api/internal/logic/resource_catalog/tag/
go test -cover ./model/resource_catalog/tag/
```

---

## ✅ Gate 4: 质量检查

### 自动化检查

**在Cursor终端依次执行**：

```bash
# 1. 编译检查
go build ./...
# 应该显示：Success

# 2. 测试检查
go test ./...
# 应该显示：All Pass

# 3. 覆盖率检查
go test -cover ./model/resource_catalog/tag/
go test -cover ./api/internal/logic/resource_catalog/tag/
# 应该 > 80%

# 4. Lint检查
golangci-lint run
# 应该：No issues
```

---

### Self Review

**Cursor Prompt**：

```
@api/internal/handler/resource_catalog/tag/
@api/internal/logic/resource_catalog/tag/
@model/resource_catalog/tag/

请Review我实现的标签管理功能：

检查项：
- [ ] 是否遵循分层架构？
- [ ] 函数是否 < 50行？
- [ ] 是否有完整的中文注释？
- [ ] 错误处理是否完整？
- [ ] 是否有hardcode的值？
- [ ] 命名是否符合规范？

请给出改进建议
```

按建议修改后：

```bash
git add .
git commit -m "refactor: code review improvements"
```

---

### 创建PR

**Cursor Prompt**：

```
帮我生成一个PR描述，包括：
1. 功能概述
2. 主要变更
3. API列表
4. 测试结果
5. 检查清单
```

**推送并创建PR**：

```bash
git push origin feature/tag-management

# 在GitHub创建Pull Request
# 标题：feat: add tag management feature
# 描述：使用Cursor生成的PR描述
```

---

## 📊 时间统计

| Phase | 时间 |
|-------|------|
| Phase 0: Context | 15min |
| Phase 1: Specify | 30min |
| Phase 2: Design | 40min |
| Phase 3: Tasks | 20min |
| Phase 4: Implement | 5h |
| **总计** | **~6.5h** |

---

## 💡 Cursor使用技巧

### 1. 善用@符号引用

```
@filename           # 引用单个文件
@folder/            # 引用整个文件夹
@Workspace          # 引用整个工作区
@Docs               # 引用文档网址
```

### 2. 分阶段提问

❌ **不好的prompt**：
```
帮我实现标签管理功能
```

✅ **好的prompt**：
```
@specs/features/tag-management/tasks.md
@model/resource_catalog/tag/interface.go

执行Task 2: 实现GORM DAO

要求：
- 实现所有Model接口方法
- 使用GORM
- 错误处理
- 中文注释
```

### 3. 持续Review

每完成一个Task就让Cursor Review：

```
@刚生成的文件

Review这个文件：
- 是否符合规范？
- 有没有潜在bug？
- 可以如何优化？
```

### 4. 利用Terminal

直接在Cursor终端运行命令，实时验证：

```bash
go build ./...
go test ./...
go run api.go
```

### 5. 善用Composer

对于复杂任务，使用Cursor Composer（Cmd+I）：
- 可以同时编辑多个文件
- 可以看到完整的diff
- 可以一次性Accept或Reject

---

## ⚠️ 常见陷阱

### 1. 过度依赖AI

❌ **错误做法**：
```
帮我实现整个标签管理功能，包括所有代码和测试
```

✅ **正确做法**：
```
先Phase 1定义需求 → Phase 2设计 → Phase 3拆分任务 → Phase 4逐个实现
每步都Review和调整
```

### 2. 忽略规范

❌ **错误做法**：
```
实现CreateTag功能
```

✅ **正确做法**：
```
@private_doc/spec/architecture/layered-architecture.md
@private_doc/spec/coding-standards/go-style-guide.md

实现CreateTag功能，要求遵循项目规范...
```

### 3. 跳过测试

❌ **错误做法**：
```
只实现功能代码，测试以后再说
```

✅ **正确做法**：
```
每完成一个Logic就写测试，保持覆盖率>80%
```

---

## 🎯 最佳实践总结

### DO ✅

1. **严格遵循5阶段**
   - 不要跳过Phase 1和Phase 2
   - 每个阶段都有明确产出

2. **充分引用规范**
   - 每个prompt都@相关规范文件
   - 让Cursor理解项目标准

3. **小步快跑**
   - 一次完成一个Task
   - 每个Task都测试验证

4. **主动Review**
   - 不要盲目接受AI生成的代码
   - 每次生成后都Review

5. **及时提交**
   - 每完成一层就commit
   - 保持git历史清晰

### DON'T ❌

1. **不要一次性生成所有代码**
   - 容易出错
   - 难以Review
   - 不符合规范

2. **不要忽略错误处理**
   - 每个error都要处理
   - 使用%w包装

3. **不要硬编码**
   - 配置放配置文件
   - 常量用const定义

4. **不要超过函数大小限制**
   - Handler ≤ 30行
   - Logic ≤ 50行

---

## 📝 总结

使用Cursor完成开发的关键：

1. **规范先行** - 始终引用项目规范
2. **分阶段执行** - 严格遵循5阶段工作流
3. **小步迭代** - 每次只完成一个Task
4. **持续验证** - 边写边测试
5. **主动Review** - 不盲目相信AI

---

**Cursor + 5阶段工作流 = 高效高质量开发！** 🚀
