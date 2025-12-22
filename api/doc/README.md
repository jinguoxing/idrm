# API 定义文件说明

本目录存放所有 API 定义文件（.api 格式），采用模块化组织。

## 目录结构

```
doc/
├── api.api                      # 主API文件，导入所有模块
├── resource_catalog/            # 资源目录模块
│   └── category.api            # 类别相关API
├── data_view/                  # 数据视图模块（待添加）
│   └── data_view.api
└── data_understanding/         # 数据理解模块（待添加）
    └── data_understanding.api
```

## 使用说明

### 1. 生成代码

在 `api/` 目录下执行：

```bash
# 使用 doc/api.api 作为入口文件
goctl api go -api doc/api.api -dir .
```

### 2. 添加新模块

1. 在 `doc/` 下创建模块目录，如 `data_view/`
2. 创建模块的 API 定义文件，如 `data_view.api`
3. 在 `doc/api.api` 中添加导入语句：
   ```
   import "data_view/data_view.api"
   ```

### 3. 模块内多文件

如果一个模块有多个 API 文件，可以在模块目录下创建多个 .api 文件，然后在主文件中逐一导入。

## 注意事项

- 所有 import 路径都是相对于 `doc/` 目录的
- 每个模块的 API 文件应该独立完整，包含类型定义和服务定义
- 使用 `@server(group: xxx)` 来指定生成的 handler 和 logic 的目录
- service 名称统一使用 `Api`
