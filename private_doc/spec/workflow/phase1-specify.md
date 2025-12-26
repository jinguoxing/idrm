# Phase 1: Specify (需求规范)

> **Level**: MUST  
> **Version**: 2.0.0

## 目标

定义清晰的需求和验收标准，使用EARS notation确保可测试性。

## EARS Notation

**格式**：
```
WHEN [condition/event] THE SYSTEM SHALL [expected behavior]
```

**示例**：
```
WHEN a user submits invalid data
THE SYSTEM SHALL display validation errors
```

## 输出模板

`specs/features/{name}/requirements.md`

## AI工具使用

- **Kiro.dev**: 创建Spec → Requirements
- **Cursor**: @spec对话生成
- **Claude CLI**: 批量生成

## 质量门禁

- [ ] 用户故事完整
- [ ] 使用EARS notation
- [ ] 技术约束明确

参考：`../quality/quality-gates.md#gate-1`
