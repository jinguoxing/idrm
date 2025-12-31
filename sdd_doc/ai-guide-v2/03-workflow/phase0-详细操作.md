# Phase 0 详细操作指南

> Context准备 - 万事开头难

---

## 🎯 本阶段目标

在开始编码前，充分理解项目规范，准备好开发环境。

**时间投入**：
- 新人：2-4小时
- 熟手：15-30分钟

---

## 📋 详细操作步骤

### Step 1: 阅读核心规范（必做）

#### 1.1 项目宪章 (15分钟)

```
打开：sdd_doc/spec/core/project-charter.md

重点理解：
✅ 项目使命和价值观
✅ 核心开发原则
✅ AI协作规范
✅ 质量标准
```

**自检问题**：
- [ ] 理解了IDRM的4大核心价值观？
- [ ] 知道哪些是MUST原则？
- [ ] 明白AI的角色定位？

#### 1.2 技术栈 (10分钟)

```
打开：sdd_doc/spec/core/tech-stack.md

重点理解：
✅ Go 1.21+ 
✅ Go-Zero v1.9+
✅ MySQL 8.0
✅ 双ORM（GORM + SQLx）
```

**自检问题**：
- [ ] 知道项目使用的核心技术？
- [ ] 理解为什么用双ORM？
- [ ] 了解版本要求？

#### 1.3 工作流 (10分钟)

```
打开：sdd_doc/spec/core/workflow.md

重点理解：
✅ 5阶段工作流
✅ 每个阶段的目标
✅ 质量门禁
```

**自检问题**：
- [ ] 能说出5个Phase名称？
- [ ] 知道每个Phase的输出？
- [ ] 理解质量门禁的作用？

---

### Step 2: 了解架构规范（按需）

#### 2.1 分层架构（重要）

```
打开：sdd_doc/spec/architecture/layered-architecture.md

重点理解：
✅ Handler → Logic → Model
✅ 每层的职责
✅ 依赖规则
```

**实操练习**：
```
找一个现有的handler代码
查看它如何调用Logic
查看Logic如何调用Model
```

#### 2.2 双ORM模式

```
打开：sdd_doc/spec/architecture/dual-orm-pattern.md

重点理解：
✅ GORM用于复杂查询
✅ SQLx用于简单CRUD
✅ 如何选择
```

#### 2.3 API设计规范

```
打开：sdd_doc/spec/architecture/api-service-guide.md

快速浏览：
- RESTful规范
- 请求/响应格式
- 错误码
```

---

### Step 3: 熟悉编码规范（重要）

#### 3.1 Go代码风格 (20分钟)

```
打开：sdd_doc/spec/coding-standards/go-style-guide.md

重点理解：
✅ Import分组规则
✅ 命名规范
✅ 函数<50行
✅ 中文注释要求
```

**实操练习**：
```go
// 检查现有代码
cat api/internal/handler/xxx/xxxhandler.go

观察：
1. import是否分组？
2. 函数是否<50行？
3. 是否有中文注释？
```

#### 3.2 错误处理

```
打开：sdd_doc/spec/coding-standards/error-handling.md

重点理解：
✅ 错误封装用%w
✅ 分层错误处理
✅ 日志记录
```

#### 3.3 测试规范

```
打开：sdd_doc/spec/coding-standards/testing-standards.md

重点理解：
✅ 覆盖率>80%
✅ 表驱动测试
✅ Mock使用
```

---

### Step 4: 准备开发环境（必做）

#### 4.1 代码克隆

```bash
# 如果还没有
git clone https://github.com/xxx/idrm.git
cd idrm
```

#### 4.2 依赖安装

```bash
# Go依赖
go mod download

# 验证
go build ./...
```

#### 4.3 数据库启动

```bash
# Docker启动MySQL
docker-compose up -d mysql

# 验证连接
mysql -h 127.0.0.1 -P 3306 -u root -p
```

#### 4.4 工具配置

**Cursor配置**：
```bash
# 确认.cursorrules存在
ls -la .cursorrules

# 如果不存在，从模板复制
cp sdd_doc/spec/tools/cursor/.cursorrules.template .cursorrules
```

**Claude CLI配置**：
```bash
# 设置API Key
export ANTHROPIC_API_KEY="your-key"

# 确认.clinerules存在
ls -la .clinerules

# 如果不存在，从模板复制
cp sdd_doc/spec/tools/claude-cli/.clinerules.template .clinerules
```

**Kiro.dev配置**：
```bash
# 创建steering目录
mkdir -p .kiro/steering/

# 复制模板
cp sdd_doc/spec/tools/kiro/steering.template.md .kiro/steering/idrm-rules.md
```

---

### Step 5: AI工具测试（推荐）

#### 5.1 Cursor测试

```
打开Cursor
Cmd+L 打开Chat

输入：
@sdd_doc/spec/core/project-charter.md
总结IDRM的核心价值观

验证：AI能正确读取规范
```

#### 5.2 Claude CLI测试

```bash
claude --files "sdd_doc/spec/core/project-charter.md" \
       "总结项目的技术栈"

# 应该返回：Go, Go-Zero, MySQL等
```

#### 5.3 Kiro.dev测试

```
打开Kiro
检查Steering是否生效
创建一个测试Spec验证
```

---

## ✅ Phase 0完成清单

### 必做项（MUST）
- [ ] 阅读project-charter.md
- [ ] 阅读tech-stack.md  
- [ ] 阅读workflow.md
- [ ] 阅读layered-architecture.md
- [ ] 阅读go-style-guide.md
- [ ] 代码已克隆
- [ ] 依赖已安装
- [ ] 数据库已启动
- [ ] 至少配置一个AI工具

### 推荐项（SHOULD）
- [ ] 阅读dual-orm-pattern.md
- [ ] 阅读error-handling.md
- [ ] 阅读testing-standards.md
- [ ] 配置所有3个AI工具
- [ ] 完成AI工具测试

### 可选项（COULD）
- [ ] 阅读所有architecture/文档
- [ ] 阅读所有coding-standards/文档
- [ ] 浏览quality/文档
- [ ] 浏览governance/文档

---

## 🔧 常见问题

### Q1: 一定要全部读完吗？
**A**: 不是。MUST项必读，其他按需。

### Q2: 读不懂某些规范怎么办？
**A**: 
1. 先跳过，做中学
2. 问AI：`@spec文档 "xx是什么意思？"`
3. 问团队成员

### Q3: 环境配置失败怎么办？
**A**: 
1. 查看`docs/环境搭建.md`
2. 检查版本要求
3. 寻求帮助

### Q4: 可以跳过Phase 0吗？
**A**: 
- 简单功能：可以快速浏览
- 复杂功能：强烈建议完整执行
- 新人：必须执行

---

## 💡 效率提升技巧

### 技巧1：制作自己的Checklist

```markdown
# 我的Phase 0清单

## 核心规范（每次必看）
- [ ] project-charter
- [ ] workflow  
- [ ] layered-architecture

## 按需规范（看情况）
- [ ] dual-orm（Model层开发时）
- [ ] api-design（API开发时）
- [ ] error-handling（错误处理时）

## 环境（第一次）
- [x] 代码克隆
- [x] 依赖安装
- [x] 数据库
- [x] AI工具
```

###技巧2：用AI辅助学习

```
Cursor对话：
@spec/core/ "总结核心规范要点，输出思维导图"

@spec/architecture/ "对比GORM和SQLx的使用场景"

@spec/coding-standards/ "列出最重要的10条编码规范"
```

### 技巧3：边做边学

不要等全部学会再开始，先快速浏览MUST项，然后边开发边查阅。

---

## 🎓 学习路径

### 第1次（完整学习）
- 时间：2-4小时
- 认真阅读所有MUST项
- 完成环境配置
- 测试AI工具

### 第2次（复习）
- 时间：30分钟
- 快速浏览核心规范
- 检查环境
- 确认AI工具可用

### 第3次及以后（熟练）
- 时间：5-10分钟
- 扫一眼workflow
- 启动环境
- 开始Phase 1

---

## 📊 Phase 0输出

**无形资产**：
- ✅ 对项目规范的理解
- ✅ 准备就绪的开发环境
- ✅ 可用的AI工具

**有形资产**（可选）：
- 个人学习笔记
- 自己的Checklist
- 环境配置脚本

---

**Phase 0做好了，后面都顺利！** 🚀
