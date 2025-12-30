# Claude CLI使用指南

> 批量处理，自动化利器

---

## 🎯 为什么选Claude CLI？

- ⭐⭐⭐ 批量处理能力强
- ⭐⭐⭐ CI/CD集成方便
- ⭐⭐⭐ 脚本化自动化

**最适合**：批量生成测试、自动Review、CI/CD

---

## ⚙️ 配置

### 1. 安装

```bash
pip install anthropic-cli
# 或
brew install anthropic-cli
```

### 2. 配置API Key

```bash
export ANTHROPIC_API_KEY="your-api-key"

# 写入.bashrc或.zshrc
echo 'export ANTHROPIC_API_KEY="your-key"' >> ~/.zshrc
```

### 3. 配置.clinerules

项目根目录创建`.clinerules`：

```markdown
# IDRM Rules for Claude CLI

## Workflow
5-Phase工作流

## Architecture
Handler → Logic → Model

## Standards
- Chinese comments
- Functions < 50 lines

## Quality
- Build pass
- Test >80%

## Specs
sdd_doc/spec/
```

---

## 🚀 基础用法

### 1. 单文件处理

```bash
claude --files "spec/coding-standards/go-style-guide.md" \
       --files "api/internal/handler/category/handler.go" \
       "Review这个handler是否符合规范"
```

### 2. 批量处理

```bash
# 为所有handler生成测试
for file in api/internal/handler/**/*handler.go; do
    claude --files "spec/coding-standards/testing-standards.md" \
           --files "$file" \
           "生成单元测试" > "${file%.*}_test.go"
done
```

### 3. Review PR

```bash
# 获取变更文件
changed_files=$(git diff --name-only main)

# Claude Review
claude --files "spec/**/*.md" \
       --files "$changed_files" \
       "Review代码，输出markdown报告"
```

---

## 📖 实战场景

### 场景1：批量生成测试

```bash
#!/bin/bash
# scripts/generate-tests.sh

for file in model/**/*_dao.go; do
    test_file="${file%.*}_test.go"
    if [ ! -f "$test_file" ]; then
        echo "生成测试: $file"
        claude --files "spec/coding-standards/testing-standards.md" \
               --files "$file" \
               "生成单元测试，要求：
               1. 表驱动测试
               2. 覆盖所有方法
               3. Mock外部依赖" > "$test_file"
    fi
done
```

### 场景2：自动Code Review

```bash
#!/bin/bash
# scripts/auto-review.sh

# 获取PR变更
git fetch origin main
changed_files=$(git diff --name-only origin/main...HEAD)

# Claude Review
claude --files "spec/**/*.md" \
       --files "$changed_files" \
      "Review代码，检查：
       1. 分层架构
       2. 编码规范
       3. 错误处理
       4. 测试覆盖
       
       输出markdown格式，分MUST/SHOULD/COULD三级" > review-report.md

echo "✅ Review完成，查看 review-report.md"
```

### 场景3：生成文档

```bash
claude --files "api/internal/handler/**/*.go" \
       "分析这些handler，生成API文档
       格式：
       - 接口路径
       - 请求参数
       - 响应格式
       - 示例" > docs/API.md
```

---

## 🤖 CI/CD集成

### GitHub Actions

```yaml
# .github/workflows/claude-review.yml
name: Claude Review

on: [pull_request]

jobs:
  review:
    runs-on: ubuntu-latest
    env:
      ANTHROPIC_API_KEY: ${{ secrets.ANTHROPIC_API_KEY }}
    
    steps:
      - uses: actions/checkout@v2
        with:
          fetch-depth: 0
      
      - name: Get changed files
        id: files
        run: |
          echo "files=$(git diff --name-only origin/main...HEAD | tr '\n' ' ')" >> $GITHUB_OUTPUT
      
      - name: Claude Review
        run: |
          claude --files "sdd_doc/spec/**/*.md" \
                 --files "${{ steps.files.outputs.files }}" \
                 "Review code against specs" > review.md
      
      - name: Comment PR
        uses: actions/github-script@v6
        with:
          script: |
            const fs = require('fs');
            const review = fs.readFileSync('review.md', 'utf8');
            github.rest.issues.createComment({
              issue_number: context.issue.number,
              owner: context.repo.owner,
              repo: context.repo.repo,
              body: review
            });
```

---

## 💡 最佳实践

1. **引用规范** - 总是用--files引用spec
2. **明确输出格式** - 指定markdown、json等格式
3. **批处理脚本** - 封装常用操作
4. **版本控制** - 生成内容提交git
5. **人工验证** - AI输出需要Review

---

## 🔧 辅助脚本

### spec.sh - 快速引用规范

```bash
#!/bin/bash
# scripts/spec.sh

spec_files="sdd_doc/spec/core/*.md sdd_doc/spec/architecture/*.md"

claude --files "$spec_files" "$@"
```

使用：
```bash
./scripts/spec.sh "总结项目规范"
```

---

详细指南：`sdd_doc/ClaudeCodeCLI使用指南.md`

**Claude CLI，自动化首选！** 🤖
