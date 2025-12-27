# Phase 4 详细操作指南

> Implement - 实施验证，落地成果

---

## 🎯 本阶段目标

按照Phase 3的任务列表，逐个实施并验证。

**时间投入**：取决于任务数量和复杂度

**核心产物**：
- 完整代码
- 单元测试
- 通过所有质量门禁

---

## 📋 详细操作步骤

### Step 1: 环境准备

#### 1.1 创建分支

```bash
# 从main创建feature分支
git checkout main
git pull origin main
git checkout -b feature/category-management

# 验证
git branch
```

#### 1.2 确认环境

```bash
# 编译通过
go build ./...

# 测试通过
go test ./...

# 数据库连接
# 启动服务验证
```

---

### Step 2: 逐个执行Task

#### 2.1 Task执行模板

**对于每个Task**：
1. 标记状态为🚧 In Progress
2. 创建/修改文件
3. 编写代码
4. Self Review
5. 运行测试
6. 标记状态为✅ Completed
7. Commit

#### 2.2 Task 1示例：创建Model接口

**Step 1: 创建文件**
```bash
mkdir -p model/resource_catalog/category
cd model/resource_catalog/category
touch interface.go types.go errors.go
```

**Step 2: AI生成代码（Cursor）**
```
@private_doc/spec/architecture/layered-architecture.md
@private_doc/spec/architecture/dual-orm-pattern.md
@tasks.md

执行Task 1: 创建Category Model接口

要求：
1. Model接口包含8个方法（Insert/FindOne/List/Update等）
2. Category结构体完整定义
3. 错误定义
4. 完整中文注释
5. 遵循编码规范
```

**Step 3: Review代码**
```
检查：
- [ ] 接口方法签名正确？
- [ ] 使用context.Context？
- [ ] 返回error？
- [ ] 中文注释？
- [ ] 编码规范？
```

**Step 4: 编译检查**
```bash
go build ./model/resource_catalog/category/
```

**Step 5: Commit**
```bash
git add model/resource_catalog/category/
git commit -m "feat: add category model interface and types"
```

---

### Step 3: 编写测试

#### 3.1 测试优先还是后写？

**建议**：边实现边测试
```
Task 2: 实现GORM DAO
├─ 实现Insert方法
├─ 测试Insert方法  ← 立即测试
├─ 实现FindOne方法
├─ 测试FindOne方法 ← 立即测试
...
```

#### 3.2 表驱动测试模板

```go
func TestGormDao_Insert(t *testing.T) {
    // Setup
    db := setupTestDB(t)
    dao := NewGormDao(db)
    ctx := context.Background()
    
    tests := []struct {
        name    string
        input   *Category
        wantErr bool
        errType error
    }{
        {
            name: "正常插入",
            input: &Category{
                Name: "测试分类",
                Code: "test_category",
            },
            wantErr: false,
        },
        {
            name: "名称为空",
            input: &Category{
                Code: "test",
            },
            wantErr: true,
        },
        {
            name: "编码重复",
            input: &Category{
                Name: "重复",
                Code: "duplicate", // 假设已存在
            },
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Execute
            err := dao.Insert(ctx, tt.input)
            
            // Assert
            if tt.wantErr {
                assert.Error(t, err)
                if tt.errType != nil {
                    assert.ErrorIs(t, err, tt.errType)
                }
            } else {
                assert.NoError(t, err)
                assert.NotZero(t, tt.input.ID)
            }
        })
    }
}
```

#### 3.3 Mock测试（Logic/Handler层）

```go
// mockgen生成Mock
//go:generate mockgen -source=interface.go -destination=mock_category.go -package=category

func TestCreateCategoryLogic_Create(t *testing.T) {
    // Setup
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()
    
    mockModel := NewMockModel(ctrl)
    logic := NewCreateCategoryLogic(mockModel)
    
    // Mock行为
    mockModel.EXPECT().
        ExistsByName(gomock.Any(), "测试分类", gomock.Any()).
        Return(false, nil)
    
    mockModel.EXPECT().
        Insert(gomock.Any(), gomock.Any()).
        Return(nil)
    
    // Execute
    result, err := logic.Create(context.Background(), &CreateCategoryReq{
        Name: "测试分类",
        Code: "test",
    })
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
}
```

---

### Step 4: 质量检查

#### 4.1 编译检查

```bash
# 编译整个项目
go build ./...

# 如果失败，修复后重新编译
```

#### 4.2 测试检查

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./model/resource_catalog/category/

# 带覆盖率
go test -cover ./model/resource_catalog/category/

# 生成覆盖率报告
go test -coverprofile=coverage.out ./model/resource_catalog/category/
go tool cover -html=coverage.out -o coverage.html
```

**覆盖率要求**: >80%

#### 4.3 Lint检查

```bash
# 运行golangci-lint
golangci-lint run

# 运行特定目录
golangci-lint run ./model/resource_catalog/category/

# 修复自动可修复的问题
golangci-lint run --fix
```

---

### Step 5: Self Review

#### 5.1 Code Review Checklist

```markdown
## Self Review Checklist

### 功能正确性
- [ ] 所有Task的验收标准都满足
- [ ] 边界情况处理正确
- [ ] 错误处理完整

### 架构合规
- [ ] 遵循分层架构
- [ ] Handler不包含业务逻辑
- [ ] Logic不直接访问DB
- [ ] 依赖方向正确

### 代码质量
- [ ] 函数<50行
- [ ] 变量命名清晰
- [ ] 无magic number
- [ ] 无重复代码

### 注释文档
- [ ] 所有公开函数有中文注释
- [ ] 复杂逻辑有解释
- [ ] TODO标记清晰

### 测试覆盖
- [ ] 单元测试>80%
- [ ] 关键路径都测试
- [ ] 异常情况都测试

### 性能考虑
- [ ] 无N+1查询
- [ ] 合理使用索引
- [ ] 避免过度查询
```

#### 5.2 AI辅助Review

```
@model/resource_catalog/category/gorm_dao.go
@spec/coding-standards/code-review-checklist.md

Review这个文件：
1. 是否符合编码规范？
2. 有什么安全问题？
3. 性能有问题吗？
4. 可以改进的地方？

输出Review报告
```

---

### Step 6: Peer Review

#### 6.1 创建Pull Request

```bash
# 推送分支
git push origin feature/category-management

# 在GitHub创建PR
# 标题：feat: add category management
# 描述：参考PR模板
```

#### 6.2 PR描述模板

```markdown
## 功能描述
添加资源分类管理功能

## 变更内容
- Model层：Category的GORM和SQLx实现
- Logic层：CRUD业务逻辑
- Handler层：5个REST API接口
- 测试：单元测试覆盖率85%

## 测试情况
- [X] 单元测试通过
- [X] 集成测试通过  
- [X] 手动测试通过

## 检查清单
- [X] go build通过
- [X] go test通过
- [X] golangci-lint通过
- [X] 测试覆盖率>80%
- [X] Self Review完成

## 相关文档
- Requirements: specs/category/requirements.md
- Design: specs/category/design.md
- Tasks: specs/category/tasks.md

## 截图/演示
（可选）API测试截图

## Review重点
请重点关注：
1. Model层的双ORM实现是否合理
2. Logic层的业务逻辑是否正确
3. 错误处理是否完整
```

#### 6.3 响应Review意见

```markdown
## Review意见处理

### ✅ 已处理
1. 修复了XXX问题
2. 优化了YYY逻辑
3. 补充了ZZZ测试

### 💬 讨论中
1. 关于AAA的建议，我认为...

### ❓待澄清
1. BBB的意见不太理解，能详细说明吗？
```

---

### Step 7: 集成验证

#### 7.1 本地验证

```bash
# 启动服务
go run api/api.go

# 测试API
curl -X POST http://localhost:8888/api/v1/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"测试分类","code":"test"}'

# 查询列表
curl http://localhost:8888/api/v1/categories

# 查询详情
curl http://localhost:8888/api/v1/categories/1

# 更新
curl -X PUT http://localhost:8888/api/v1/categories/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"更新后的名称"}'

# 删除
curl -X DELETE http://localhost:8888/api/v1/categories/1
```

#### 7.2 使用Postman/Insomnia

创建测试集合：
```
Collection: Category Management

Requests:
├─ Create Category
├─ List Categories
├─ Get Category
├─ Update Category
└─ Delete Category

Environment:
- base_url: http://localhost:8888
- api_version: v1
```

---

## ✅ Gate 4 质量检查

### 自动化检查

```bash
# 1. 编译
go build ./...
if [ $? -ne 0 ]; then
    echo "❌ Build failed"
    exit 1
fi

# 2. 测试
go test -cover ./... | tee test-results.txt
coverage=$(grep "coverage" test-results.txt | awk '{print $5}' | tr -d '%')
if [ $coverage -lt 80 ]; then
    echo "❌ Coverage $coverage% < 80%"
    exit 1
fi

# 3. Lint
golangci-lint run
if [ $? -ne 0 ]; then
    echo "❌ Lint failed"
    exit 1
fi

echo "✅ All checks passed!"
```

### 手动检查

- [ ] Self Review完成
- [ ] Peer Review通过
- [ ] 功能验证通过
- [ ] 性能测试通过（如需要）
- [ ] 文档已更新

---

## 💡 实战技巧

### 技巧1：小步提交

每完成一个Task就commit：
```bash
git commit -m "feat: add category model interface"
git commit -m "feat: implement gorm dao for category"
git commit -m "feat: add category create logic"
```

好处：
- 易于Review
- 易于回滚
- 清晰的历史

### 技巧2：测试驱动

```
1. 写测试（红）
2. 写实现（绿）
3. 重构（重构）
4. 重复
```

### 技巧3：并行开发

如果任务独立，可并行：
```bash
# Terminal 1
cd model/resource_catalog/category
# 开发Model层

# Terminal 2
# 等Model接口定义好后
cd api/internal/logic/resource_catalog/category
# 并行开发Logic层
```

### 技巧4：问题列表

遇到问题立即记录：
```markdown
## Issues

1. [ ] GORM预加载性能问题 - 待优化
2. [X] 并发写入冲突 - 已用乐观锁解决
3. [ ] 测试数据清理不彻底 - 待修复
```

---

## 🔧 常见问题

### Q1: 测试写不出来怎么办？
**A**:
```
@logic/createcategorylogic.go
@spec/coding-standards/testing-standards.md

为这个Logic生成表驱动测试
要求：
1. 覆盖正常流程
2. 覆盖所有异常
3. 使用Mock
```

### Q2: Lint报错太多？
**A**:
1. 先修复明显的
2. 用AI批量修复：`@文件 "修复所有lint错误"`
3. 实在不行，临时禁用（不推荐）

### Q3: 覆盖率不够80%？
**A**:
1. 查看覆盖率报告：`go tool cover -html=coverage.out`
2. 补充未覆盖的分支
3. 用AI生成测试

### Q4: Review意见太多，改不完？
**A**:
1. 优先级：P0 > P1 > P2
2. 分批修复
3. 讨论不合理的意见

---

## 📊 Phase 4输出

### 完整代码

```
model/resource_catalog/category/
├── interface.go      ✅
├── types.go          ✅
├── errors.go         ✅
├── gorm_dao.go       ✅
├── sqlx_model.go     ✅
├── factory.go        ✅
└── *_test.go         ✅ (>80% coverage)

api/internal/logic/resource_catalog/category/
├── *.go              ✅
└── *_test.go         ✅

api/internal/handler/resource_catalog/category/
├── *.go              ✅
└── *_test.go         ✅
```

### 质量指标

```
✅  Build: PASS
✅  Tests: PASS (Coverage: 85%)
✅  Lint:  PASS
✅  Review: APPROVED
```

---

**Phase 4完成了，功能上线！** 🚀
