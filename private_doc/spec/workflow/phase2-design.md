# Phase 2: Design (技术方案)

> **Level**: MUST  
> **Version**: 2.0.0

## 目标

创建详细的技术方案，包括架构设计、接口定义、序列图。

## 输入
- Phase 1的requirements.md

## 输出模板

创建 `specs/features/{name}/design.md`：

```markdown
# Design: {功能名}

## Architecture Overview
遵循IDRM分层架构

## File Structure
详细的文件清单

## Interface Definitions
\`\`\`go
type Model interface { ... }
\`\`\`

## Sequence Diagrams
流程图

## Implementation Considerations
实现注意事项
```

## 质量门禁

- [ ] 符合分层架构
- [ ] 文件清单完整
- [ ] 接口定义清晰
- [ ] 序列图完整

参考：`../quality/quality-gates.md#gate-2`

## 下一步
→ Phase 3: Tasks
