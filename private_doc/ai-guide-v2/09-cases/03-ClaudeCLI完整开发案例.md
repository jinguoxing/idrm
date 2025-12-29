# 实战案例：使用Claude Code CLI完成完整开发

> 批量处理和自动化 - CLI工作流

---

## 📋 案例概述

**功能**：数据标签（Tag）管理  
**复杂度**：中等（约400行）  
**工具**：Claude Code CLI  
**耗时**：1个工作日  
**适用场景**：批量处理、自动化、CI/CD集成

---

## 💡 为什么选择Claude Code CLI？

**适合CLI的场景**：
- ✅ 批量生成多个文件
- ✅ 文档和测试自动化生成
- ✅ CI/CD pipeline集成
- ✅ 代码审查和重构
- ✅ 规范检查和修复

**CLI的优势**：
- 🚀 批量处理，一次生成多个文件
- 🤖 可脚本化，易于自动化
- 📦 适合大规模重构
- 🔄 CI/CD友好
- 📊 输出可重定向和处理

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

### Step 1: 准备环境

**安装Claude CLI**（如果尚未安装）：

```bash
# 安装
npm install -g @anthropic-ai/claude-cli

# 配置API Key
export ANTHROPIC_API_KEY=your_api_key_here

# 验证安装
claude --version
```

### Step 2: 批量阅读规范

**使用CLI批量读取规范**：

```bash
# 读取核心规范并保存摘要
claude --files "private_doc/spec/core/*.md" \
  "请总结这些规范文档的关键点，特别是：
  1. 5阶段工作流
  2. 分层架构原则
  3. 编码规范要点
  
  输出为要点清单" > specs_summary.txt

# 查看摘要
cat specs_summary.txt
```

### Step 3: 环境检查

```bash
# 拉取最新代码
git pull origin main

# 编译检查
go build ./...

# 创建分支
git checkout -b feature/tag-management
```

---

## 📋 Phase 1: Specify (30分钟)

### Step 1: 批量生成requirements.md

**创建prompt文件**：

```bash
# 创建 prompt_phase1.txt
cat > prompt_phase1.txt << 'EOF'
我要添加数据标签管理功能。

Phase 1: 请生成完整的 requirements.md 文档

功能需求：
1. 创建/删除标签
2. 为资源（数据目录）打标签/取消标签
3. 按标签查询资源
4. 标签使用统计

要求：
1. 包含User Stories (AS a/I WANT/SO THAT格式)
2. 使用EARS notation编写Acceptance Criteria
3. 覆盖正常流程、参数验证、异常情况、边界条件
4. 定义Business Rules（唯一性、长度、关联关系、删除规则）
5. 定义Data Considerations（需要持久化的数据描述）
6. **不包含任何技术实现细节**（无Technical Constraints、无Data Model）

参考规范：
- private_doc/spec/workflow/phase1-specify.md
- private_doc/spec/workflow/ears-notation-guide.md
EOF
```

**执行生成**：

```bash
# 创建目录
mkdir -p specs/features/tag-management

# 生成requirements.md
claude --files "private_doc/spec/workflow/phase1-specify.md" \
       --files "private_doc/spec/workflow/ears-notation-guide.md" \
       < prompt_phase1.txt > specs/features/tag-management/requirements.md

# 查看生成的文档
cat specs/features/tag-management/requirements.md
```

---

### Step 2: Gate 1 检查

**使用CLI检查质量**：

```bash
# 创建检查prompt
cat > check_gate1.txt << 'EOF'
请检查这个 requirements.md 是否通过 Gate 1

检查项：
- [ ] User stories完整 (AS/I WANT/SO THAT)
- [ ] 使用EARS notation
- [ ] Business rules明确
- [ ] Data considerations清晰
- [ ] **没有技术实现细节**

请逐项检查并指出问题，如果有问题请给出修正建议。
如果全部通过，请在文档末尾添加Gate 1检查清单。
EOF

# 执行检查
claude --files "specs/features/tag-management/requirements.md" \
       --files "private_doc/spec/quality/quality-gates.md" \
       < check_gate1.txt

# 如果有问题，根据建议修正后重新生成
```

---

### Step 3: 提交Phase 1

```bash
git add specs/features/tag-management/requirements.md
git commit -m "docs: add tag management requirements (Phase 1)"
```

---

## 🎨 Phase 2: Design (40分钟)

### Step 1: 生成design.md

**创建prompt**：

```bash
cat > prompt_phase2.txt << 'EOF'
Phase 2: 基于requirements.md，请生成完整的 design.md 技术设计文档

要求：
1. 遵循分层架构 (Handler→Logic→Model)
2. 定义完整的文件结构（所有需要创建的文件）
3. 定义Model接口（Go interface）
4. 绘制序列图（Mermaid格式）
5. 定义数据库表结构（带索引和中文注释）
6. 说明Technical Constraints（架构、函数大小、注释、测试、ORM选择）
7. 说明ORM选择理由（为什么用GORM或SQLx）

参考规范：
- private_doc/spec/architecture/layered-architecture.md
- private_doc/spec/architecture/dual-orm-pattern.md
- private_doc/spec/workflow/phase2-design.md
EOF
```

**执行生成**：

```bash
claude --files "specs/features/tag-management/requirements.md" \
       --files "private_doc/spec/architecture/layered-architecture.md" \
       --files "private_doc/spec/architecture/dual-orm-pattern.md" \
       --files "private_doc/spec/workflow/phase2-design.md" \
       < prompt_phase2.txt > specs/features/tag-management/design.md
```

---

### Step 2: Gate 2 检查

```bash
cat > check_gate2.txt << 'EOF'
请检查这个 design.md 是否通过 Gate 2

检查项：
- [ ] 符合分层架构
- [ ] 文件清单完整
- [ ] 接口定义清晰
- [ ] 序列图完整
- [ ] 数据库设计合理

请逐项检查并给出建议。
EOF

claude --files "specs/features/tag-management/design.md" \
       --files "private_doc/spec/quality/quality-gates.md" \
       < check_gate2.txt
```

---

### Step 3: 提交Phase 2

```bash
git add specs/features/tag-management/design.md
git commit -m "docs: add tag management design (Phase 2)"
```

---

## 📋 Phase 3: Tasks (20分钟)

### Step 1: 生成tasks.md

```bash
cat > prompt_phase3.txt << 'EOF'
Phase 3: 基于design.md，请生成 tasks.md 任务拆分文档

要求：
1. 每个task预估代码行数 < 50行（handler < 30行）
2. 明确依赖关系
3. 按顺序：Model → Logic → Handler → Test
4. 每个task包含：名称、预估行数、依赖、文件列表、验收标准、状态
5. 总共约12-15个tasks

格式：
## Task 1: {名称}
**Lines**: {行数}
**Status**: ⏸️ Not Started
**Depends**: {依赖}
**Files**:
- {文件1}
- {文件2}
**Acceptance**:
- [ ] {标准1}
- [ ] {标准2}
EOF

claude --files "specs/features/tag-management/design.md" \
       < prompt_phase3.txt > specs/features/tag-management/tasks.md
```

---

### Step 2: 提交Phase 3

```bash
git add specs/features/tag-management/tasks.md
git commit -m "docs: add tag management tasks (Phase 3)"
```

---

## 💻 Phase 4: Implement (4-6小时)

### 方式1: 批量生成Model层

**生成所有Model文件**：

```bash
cat > gen_model.txt << 'EOF'
请生成Tag Model层的所有文件：

1. model/resource_catalog/tag/interface.go
   - 定义完整的Model接口
   - 包含所有CRUD和关联方法
   - 中文注释

2. model/resource_catalog/tag/types.go
   - 定义Tag和ResourceTag结构体
   - GORM和db标签
   - 中文注释

3. model/resource_catalog/tag/gorm_dao.go
   - 实现所有Model接口方法
   - 使用GORM
   - 事务处理
   - 错误处理(%w)
   - 每个方法<50行

4. model/resource_catalog/tag/factory.go
   - NewModel工厂函数

要求：
- 遵循编码规范
- 完整的中文注释
- 正确的错误处理
EOF

# 生成并分割输出
claude --files "specs/features/tag-management/design.md" \
       --files "private_doc/spec/coding-standards/go-style-guide.md" \
       < gen_model.txt > model_output.txt

# 手动分割或使用脚本提取各个文件内容
```

---

### 方式2: 逐个生成（推荐）

**生成interface.go**：

```bash
cat > gen_interface.txt << 'EOF'
生成 model/resource_catalog/tag/interface.go

只生成这一个文件，包含：
1. package声明
2. imports
3. Model接口定义（所有方法）
4. 完整的中文注释
EOF

claude --files "specs/features/tag-management/design.md" \
       < gen_interface.txt > model/resource_catalog/tag/interface.go
```

**生成types.go**：

```bash
cat > gen_types.txt << 'EOF'
生成 model/resource_catalog/tag/types.go

只生成这一个文件，包含：
1. Tag结构体（GORM和db标签）
2. ResourceTag结构体
3. 中文注释
EOF

claude --files "specs/features/tag-management/design.md" \
       < gen_types.txt > model/resource_catalog/tag/types.go
```

**类似方式生成其他文件...**

---

### 批量生成Logic层

**使用循环生成所有Logic**：

```bash
# 定义所有Logic文件
LOGICS=(
    "createtaglogic"
    "deletetaglogic"
    "assigntaglogic"
    "removetaglogic"
    "querybytaglogic"
    "tagstatslogic"
)

# 循环生成
for logic in "${LOGICS[@]}"; do
    cat > gen_${logic}.txt << EOF
生成 api/internal/logic/resource_catalog/tag/${logic}.go

要求：
1. 定义Logic结构体
2. 实现对应的方法
3. 业务逻辑实现
4. 调用Model接口
5. 完整错误处理
6. 中文注释
7. 函数<50行

参考：
- design.md中的接口定义
- 编码规范
EOF

    claude --files "specs/features/tag-management/design.md" \
           --files "model/resource_catalog/tag/interface.go" \
           --files "private_doc/spec/coding-standards/go-style-guide.md" \
           < gen_${logic}.txt > api/internal/logic/resource_catalog/tag/${logic}.go
           
    echo "Generated ${logic}.go"
done
```

---

### 批量生成测试

**生成所有Logic测试**：

```bash
cat > gen_tests.txt << 'EOF'
为所有Logic文件生成单元测试

要求：
1. 每个Logic一个测试文件（*_test.go）
2. 表驱动测试
3. Mock Model接口
4. 覆盖所有分支
5. 覆盖率>80%

Logic文件：
- createtaglogic.go
- deletetaglogic.go
- assigntaglogic.go
- removetaglogic.go
- querybytaglogic.go
- tagstatslogic.go
EOF

claude --files "api/internal/logic/resource_catalog/tag/*.go" \
       --files "private_doc/spec/coding-standards/testing-standards.md" \
       < gen_tests.txt > tests_output.txt

# 手动提取各个测试文件
```

---

### 验证和测试

```bash
# 编译检查
go build ./model/resource_catalog/tag/
go build ./api/internal/logic/resource_catalog/tag/

# 运行测试
go test ./model/resource_catalog/tag/
go test ./api/internal/logic/resource_catalog/tag/

# 检查覆盖率
go test -cover ./model/resource_catalog/tag/
go test -cover ./api/internal/logic/resource_catalog/tag/
```

---

### 批量Code Review

**使用CLI审查代码**：

```bash
cat > review_code.txt << 'EOF'
请Review以下代码，检查：

1. 是否遵循分层架构？
2. 函数是否<50行？
3. 是否有完整的中文注释？
4. 错误处理是否完整？
5. 是否有hardcode？
6. 命名是否符合规范？

请给出具体的改进建议和需要修改的代码位置。
EOF

claude --files "model/resource_catalog/tag/*.go" \
       --files "api/internal/logic/resource_catalog/tag/*.go" \
       --files "private_doc/spec/coding-standards/go-style-guide.md" \
       < review_code.txt > review_result.txt

# 查看review结果
cat review_result.txt
```

---

## ✅ Gate 4: 质量检查

### 自动化检查脚本

**创建质量检查脚本**：

```bash
cat > quality_check.sh << 'EOF'
#!/bin/bash

echo "=== Gate 4 Quality Check ==="

# 1. 编译检查
echo "1. Build check..."
if go build ./...; then
    echo "✅ Build passed"
else
    echo "❌ Build failed"
    exit 1
fi

# 2. 测试检查
echo "2. Test check..."
if go test ./...; then
    echo "✅ Tests passed"
else
    echo "❌ Tests failed"
    exit 1
fi

# 3. 覆盖率检查
echo "3. Coverage check..."
coverage=$(go test -cover ./model/resource_catalog/tag/ ./api/internal/logic/resource_catalog/tag/ | grep coverage | awk '{print $2}' | tr -d '%')
if [ "$coverage" -gt 80 ]; then
    echo "✅ Coverage: ${coverage}% (>80%)"
else
    echo "❌ Coverage: ${coverage}% (<80%)"
    exit 1
fi

# 4. Lint检查
echo "4. Lint check..."
if golangci-lint run; then
    echo "✅ Lint passed"
else
    echo "❌ Lint failed"
    exit 1
fi

echo ""
echo "=== All checks passed! ==="
EOF

chmod +x quality_check.sh
./quality_check.sh
```

---

## 🔄 CI/CD集成

### GitHub Actions工作流

**创建CI配置**：

```yaml
# .github/workflows/tag-management.yml
name: Tag Management Quality Check

on:
  pull_request:
    paths:
      - 'model/resource_catalog/tag/**'
      - 'api/internal/logic/resource_catalog/tag/**'
      - 'api/internal/handler/resource_catalog/tag/**'

jobs:
  quality-check:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Setup Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.21
      
      - name: Build
        run: go build ./...
      
      - name: Test
        run: go test -cover ./...
      
      - name: Lint
        run: golangci-lint run
      
      - name: Claude Review
        env:
          ANTHROPIC_API_KEY: ${{ secrets.ANTHROPIC_API_KEY }}
        run: |
          npm install -g @anthropic-ai/claude-cli
          claude --files "model/resource_catalog/tag/*.go" \
                 --files "api/internal/logic/resource_catalog/tag/*.go" \
                 "Review this code for compliance with coding standards" \
                 > review.txt
          cat review.txt
```

---

## 📊 批量处理高级技巧

### 1. 批量重构

**重命名函数**：

```bash
cat > refactor.txt << 'EOF'
请将所有Logic文件中的：
- FindByID 改为 GetByID
- List 改为 Query

输出修改后的完整文件内容。
EOF

claude --files "api/internal/logic/resource_catalog/tag/*.go" \
       < refactor.txt
```

---

### 2. 批量添加注释

```bash
cat > add_comments.txt << 'EOF'
为所有公开函数添加完整的中文注释

格式：
// FunctionName 功能描述
// 参数说明
// 返回值说明
EOF

claude --files "model/resource_catalog/tag/*.go" \
       < add_comments.txt
```

---

### 3. 批量生成文档

```bash
cat > gen_docs.txt << 'EOF'
基于代码生成API文档

包括：
1. 所有API端点
2. 请求参数
3. 响应格式
4. 错误码
5. 使用示例

输出为Markdown格式
EOF

claude --files "api/internal/handler/resource_catalog/tag/*.go" \
       < gen_docs.txt > docs/api/tag-management.md
```

---

### 4. 批量生成Mock

```bash
cat > gen_mocks.txt << 'EOF'
为Model接口生成Mock实现

要求：
1. 使用testify/mock
2. 实现所有接口方法
3. 可配置返回值
4. 可验证调用
EOF

claude --files "model/resource_catalog/tag/interface.go" \
       < gen_mocks.txt > model/resource_catalog/tag/mock_model.go
```

---

## 💡 CLI使用最佳实践

### 1. 使用配置文件

**创建 .claude-config.json**：

```json
{
  "context_files": [
    "private_doc/spec/core/*.md",
    "private_doc/spec/architecture/*.md",
    "private_doc/spec/coding-standards/*.md"
  ],
  "output_dir": "generated",
  "max_tokens": 4096
}
```

---

### 2. 创建Prompt模板库

**建立prompt模板目录**：

```bash
mkdir -p .claude/prompts

# Phase 1模板
cat > .claude/prompts/phase1_requirements.txt << 'EOF'
生成Phase 1 requirements.md
功能：{FEATURE_NAME}
需求：{REQUIREMENTS}
参考规范：phase1-specify.md, ears-notation-guide.md
EOF

# 使用模板
sed "s/{FEATURE_NAME}/Tag Management/" \
    .claude/prompts/phase1_requirements.txt | \
claude --files "private_doc/spec/workflow/*.md" \
       > specs/features/tag-management/requirements.md
```

---

### 3. 管道处理

**多步骤处理**：

```bash
# 生成 → 检查 → 修正
claude < gen_code.txt | \
claude "检查这段代码是否符合规范，如不符合请修正" | \
tee output.go
```

---

### 4. 批量处理脚本

```bash
#!/bin/bash
# batch_generate.sh

for feature in tag category resource; do
    echo "Processing $feature..."
    
    # 生成requirements
    claude --files "specs/features/$feature/design.md" \
           < gen_requirements.txt \
           > "specs/features/$feature/requirements.md"
    
    # 生成model
    claude --files "specs/features/$feature/design.md" \
           < gen_model.txt \
           > "model/$feature/generated.go"
done
```

---

## 📝 提交和PR

### 批量提交

```bash
# 添加所有生成的文件
git add specs/features/tag-management/
git add model/resource_catalog/tag/
git add api/internal/logic/resource_catalog/tag/
git add api/internal/handler/resource_catalog/tag/

# 生成commit message
claude "基于这些变更生成commit message，使用Conventional Commits格式" \
       --files "$(git diff --cached --name-only)" \
       > commit_msg.txt

# 提交
git commit -F commit_msg.txt
```

---

### 生成PR描述

```bash
cat > gen_pr.txt << 'EOF'
基于以下文件生成Pull Request描述

包括：
1. 功能概述
2. 主要变更
3. API列表
4. 测试结果
5. checklist

格式：Markdown
EOF

claude --files "specs/features/tag-management/*.md" \
       --files "model/resource_catalog/tag/*.go" \
       --files "api/internal/logic/resource_catalog/tag/*.go" \
       < gen_pr.txt > pr_description.md

# 创建PR时使用这个描述
```

---

## ⚠️ CLI使用注意事项

### DO ✅

1. **分阶段处理**
   - 不要一次生成所有代码
   - 每个Phase单独处理
   - 每层单独生成和验证

2. **保存中间结果**
   - 所有输出都保存到文件
   - 便于Review和调整
   - 可追踪生成过程

3. **验证输出**
   - 生成后立即编译测试
   - 不要盲目信任输出
   - 手动Review关键代码

4. **使用版本控制**
   - 每个阶段都提交
   - 便于回滚
   - 保持历史清晰

### DON'T ❌

1. **不要直接覆盖现有文件**
   - 先输出到临时文件
   - Review后再覆盖

2. **不要跳过质量检查**
   - 每次生成都要测试
   - 覆盖率必须达标

3. **不要忽略错误**
   - CLI可能生成不完整代码
   - 需要手动补充

4. **不要过度自动化**
   - 关键逻辑需要人工Review
   - 不能完全依赖AI

---

## 📊 时间对比

| 阶段 | 手动 | Cursor | CLI |
|------|------|--------|-----|
| Phase 1 | 60min | 30min | **15min** |
| Phase 2 | 90min | 40min | **20min** |
| Phase 3 | 40min | 20min | **10min** |
| Phase 4 | 8h | 5h | **3h** |
| **总计** | **10.5h** | **6.5h** | **4h** |

**CLI优势**：批量处理，速度最快

---

## 🎯 总结

### CLI最适合的场景

1. ✅ **批量生成文件** - 一次生成多个相似文件
2. ✅ **文档生成** - API文档、测试报告
3. ✅ **代码重构** - 批量重命名、添加注释
4. ✅ **CI/CD集成** - 自动化检查和生成
5. ✅ **大规模代码审查** - 检查多个文件

### CLI不适合的场景

1. ❌ **需要频繁交互** - 用Cursor更好
2. ❌ **探索性编程** - CLI太慢
3. ❌ **复杂业务逻辑** - 需要人工判断
4. ❌ **新手学习** - Cursor更直观

### 工具组合建议

- **Cursor**: 日常开发、快速迭代
- **CLI**: 批量处理、文档生成、CI/CD
- **Kiro**: 大型功能规划、团队协作

---

**Claude CLI + 自动化 = 高效批量处理！** 🤖
