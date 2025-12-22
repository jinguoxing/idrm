# API 定义文件说明

本目录存放所有 API 定义文件（.api 格式），采用模块化组织。

## 目录结构

```
doc/
├── README.md             # 本文件
├── STRUCTURE.md          # 详细结构说明
├── TYPES_NAMING.md       # 类型命名规范
└── api/                  # API 定义文件
    ├── api.api           # 主API定义（导入所有模块）
    ├── resource_catalog/ # 资源目录模块
    │   └── category.api
    └── data_view/        # 数据视图模块
        └── category.api
```

## 使用说明

### 1. 生成代码

在 `api/` 目录下执行：

```bash
# 使用 doc/api/api.api 作为入口文件
goctl api go -api doc/api/api.api -dir .
```

### 2. 添加新模块

1. 在 `doc/api/` 下创建模块目录，如 `data_understanding/`
2. 创建模块的 API 定义文件，如 `data_understanding.api`
3. 在 `doc/api/api.api` 中添加导入语句：
   ```
   import "data_understanding/data_understanding.api"
   ```

### 3. 模块内多文件

如果一个模块有多个 API 文件，可以在模块目录下创建多个 .api 文件，然后在主文件中逐一导入。

## 注意事项

- 所有 import 路径都是相对于 `doc/api/` 目录的
- 每个模块的 API 文件应该独立完整，包含类型定义和服务定义
- 使用 `@server(group: xxx)` 来指定生成的 handler 和 logic 的目录
- service 名称统一使用 `Api`
- **避免类型名称冲突**：不同模块使用不同的类型名称前缀（参考 TYPES_NAMING.md）
