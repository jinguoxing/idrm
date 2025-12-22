# Swagger 文档目录

本目录用于存放 goctl 生成的 Swagger/OpenAPI 文档。

## 📋 目录说明

此目录包含项目的 API 文档，使用 Swagger/OpenAPI 规范。

## 🚀 生成 Swagger 文档

### 方法一：使用 goctl-swagger 插件（推荐）

#### 1. 安装 goctl-swagger 插件

```bash
go install github.com/zeromicro/goctl-swagger@latest
```

#### 2. 生成 Swagger JSON 文件

```bash
cd api

# 生成 swagger.json
goctl api plugin -plugin goctl-swagger="swagger -filename idrm.json -basepath /" \
  -api doc/api/api.api -dir doc/swagger
```

生成的文件：`doc/swagger/idrm.json`

### 方法二：手动转换（备选）

如果插件安装有问题，可以使用在线工具手动转换：

1. 访问 [Swagger Editor](https://editor.swagger.io/)
2. 根据 API 定义手动编写 OpenAPI 规范

## 📖 查看文档

### 使用 Swagger UI

#### 选项 A：在线查看

1. 访问 [Swagger Editor](https://editor.swagger.io/)
2. 导入生成的 `idrm.json` 文件

#### 选项 B：本地部署 Swagger UI

使用 Docker 快速启动：

```bash
docker run -p 8080:8080 \
  -e SWAGGER_JSON=/swagger/idrm.json \
  -v $(pwd)/doc/swagger:/swagger \
  swaggerapi/swagger-ui
```

访问：http://localhost:8080

#### 选项 C：使用 VS Code 插件

安装 VS Code 插件：`Swagger Viewer`

## 📝 文档结构

生成后的目录结构：

```
swagger/
├── README.md          # 本文件
└── idrm.json          # Swagger/OpenAPI 规范文件
```

## 🔄 更新文档

每次修改 API 定义后，重新运行生成命令：

```bash
cd api
goctl api plugin -plugin goctl-swagger="swagger -filename idrm.json -basepath /" \
  -api doc/api/api.api -dir doc/swagger
```

## ⚙️ 配置选项

goctl-swagger 支持的参数：

| 参数 | 说明 | 默认值 |
|------|------|--------|
| `-filename` | 输出文件名 | `swagger.json` |
| `-basepath` | API 基础路径 | `/` |
| `-host` | API 主机地址 | `localhost:8888` |
| `-schemes` | 协议（http/https） | `http` |

### 自定义示例

```bash
goctl api plugin -plugin goctl-swagger="swagger \
  -filename idrm-api.json \
  -basepath /api/v1 \
  -host api.idrm.com \
  -schemes https" \
  -api doc/api/api.api -dir doc/swagger
```

## 📌 注意事项

1. **版本控制**
   - ✅ 建议提交 swagger 文档到 Git
   - ✅ 方便团队协作和 API 文档共享

2. **自动化**
   - 可以在 CI/CD pipeline 中自动生成
   - 确保文档始终与代码保持同步

3. **文档维护**
   - 每次 API 变更后记得更新文档
   - 可以将生成命令加入 Makefile

## 🔗 相关资源

- [go-zero 文档](https://go-zero.dev/)
- [goctl-swagger GitHub](https://github.com/zeromicro/goctl-swagger)
- [Swagger/OpenAPI 规范](https://swagger.io/specification/)
- [Swagger UI](https://swagger.io/tools/swagger-ui/)

## 📋 Makefile 集成

可以在项目根目录的 Makefile 中添加：

```makefile
.PHONY: swagger
swagger:
	cd api && goctl api plugin -plugin goctl-swagger="swagger -filename idrm.json -basepath /" \
		-api doc/api/api.api -dir doc/swagger
	@echo "Swagger documentation generated at api/doc/swagger/idrm.json"
```

使用：
```bash
make swagger
```
