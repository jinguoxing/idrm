# 场景：AI辅助Code Review

> 自动化质量检查

---

## 📋 场景

**需求**：Review PR代码  
**工具**：Claude CLI  
**时间**：5分钟

---

## 🔍 手动Review

```bash
# 获取变更文件
git diff main --name-only

# Claude Review
claude --files "sdd_doc/spec/**/*.md" \\
       --files "api/internal/handler/category/*.go" \\
       "Review代码，检查：
       1. 分层架构
       2. 中文注释
       3. 错误处理
       4. 函数<50行
       
       输出markdown报告"
```

---

## 🤖 自动化Review

### GitHub Actions

```yaml
# .github/workflows/code-review.yml
name: AI Review

on: [pull_request]

jobs:
  review:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      
      - name: Claude Review
        env:
          ANTHROPIC_API_KEY: ${{ secrets.ANTHROPIC_API_KEY }}
        run: |
          claude --files "sdd_doc/spec/**/*.md" \\
                 --files "${{ github.event.pull_request.changed_files }}" \\
                 "Review against specs" > review.md
      
      - name: Comment PR
        uses: actions/github-script@v6
        with:
          script: |
            const fs = require('fs');
            const review = fs.readFileSync('review.md', 'utf8');
            github.rest.issues.createComment({
              issue_number: context.issue.number,
              body: review
            });
```

---

## 📊 Review清单

### MUST Fix
- [ ] 违反架构
- [ ] 安全问题

### SHOULD Fix
- [ ] 缺少注释
- [ ] 函数过长

---

**AI Review，提升质量！** 🎯
