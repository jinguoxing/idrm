# Docker Compose 部署指南

## 📋 概述

IDRM API 的 Docker Compose 部署方案，包含完整的服务栈和监控体系。

### 服务列表

| 服务 | 端口 | 说明 |
|------|------|------|
| **idrm-api** | 8888 | API 服务 |
| **mysql** | 3306 | MySQL 8.0 数据库 |
| **redis** | 6379 | Redis 缓存 |
| **jaeger** | 16686, 4317 | 链路追踪 |
| **elasticsearch** | 9200, 9300 | 日志存储 |
| **kibana** | 5601 | 日志查询界面 |
| **filebeat** | - | 日志收集器 |

---

## 🚀 快速开始

### 1. 前置要求

```bash
# 检查 Docker 版本（需要 20.10+）
docker --version

# 检查 Docker Compose 版本（需要 2.0+）
docker-compose --version

# 检查可用内存（建议 8GB+）
free -h
```

### 2. 启动所有服务

```bash
# 克隆项目后，进入项目目录
cd /path/to/idrm

# 启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f
```

### 3. 等待服务就绪

```bash
# 等待所有服务健康检查通过（约2-3分钟）
watch docker-compose ps

# 或者检查特定服务
docker-compose logs -f idrm-api
```

---

## 📝 服务访问

启动完成后，可以访问以下服务：

### API 服务
```bash
# 健康检查
curl http://localhost:8888/health

# API 文档（如果配置了）
http://localhost:8888/swagger
```

### Kibana 日志查询
1. 访问: http://localhost:5601
2. 首次访问需要配置 Index Pattern
3. 创建 Index Pattern: `idrm-api-*`
4. 选择时间字段: `@timestamp`
5. 开始查询日志

### Jaeger 链路追踪
- 访问: http://localhost:16686
- 选择服务: `idrm-api`
- 查看链路trace

### Elasticsearch
```bash
# 检查集群状态
curl http://localhost:9200/_cluster/health?pretty

# 查看索引
curl http://localhost:9200/_cat/indices?v
```

---

## 🔧 常用操作

### 查看日志

```bash
# 查看所有服务日志
docker-compose logs -f

# 查看特定服务日志
docker-compose logs -f idrm-api
docker-compose logs -f mysql
docker-compose logs -f elasticsearch

# 查看最近100行日志
docker-compose logs --tail=100 idrm-api
```

### 重启服务

```bash
# 重启所有服务
docker-compose restart

# 重启特定服务
docker-compose restart idrm-api

# 重新构建并启动API服务
docker-compose up -d --build idrm-api
```

### 停止服务

```bash
# 停止所有服务
docker-compose stop

# 停止并删除容器（保留数据卷）
docker-compose down

# 停止并删除所有（包括数据卷）⚠️ 慎用
docker-compose down -v
```

### 数据库操作

```bash
# 进入MySQL容器
docker-compose exec mysql mysql -uroot -pidrm@2024

# 查看数据库
SHOW DATABASES;

# 使用数据库
USE idrm_resource_catalog;

# 查看表
SHOW TABLES;
```

### 查看资源使用

```bash
# 查看容器资源使用情况
docker stats

# 查看磁盘使用
docker system df
```

---

## 🗂️ 目录结构

```
idrm/
├── docker-compose.yml           # Docker Compose 配置
├── Dockerfile                   # API 服务镜像
├── .dockerignore               # Docker 构建忽略文件
├── api/
│   └── etc/
│       ├── api.yaml            # 本地配置
│       └── api.docker.yaml     # Docker 配置
├── docker/                      # Docker 相关配置
│   ├── mysql/
│   │   └── init.sql           # MySQL 初始化脚本
│   └── filebeat/
│       └── filebeat.yml       # Filebeat 配置
└── logs/                        # 应用日志目录（自动创建）
```

---

## ⚙️ 配置说明

### 环境变量

可以通过 `.env` 文件自定义配置：

```bash
# .env
MYSQL_ROOT_PASSWORD=your_password
API_PORT=8888
JAEGER_PORT=16686
KIBANA_PORT=5601
```

### 修改配置

#### 修改 MySQL 密码

1. 编辑 `docker-compose.yml`:
```yaml
mysql:
  environment:
    MYSQL_ROOT_PASSWORD: your_new_password
```

2. 编辑 `api/etc/api.docker.yaml`:
```yaml
DB:
  ResourceCatalog:
    Password: your_new_password
```

3. 重新启动:
```bash
docker-compose down -v
docker-compose up -d
```

#### 修改API配置

编辑 `api/etc/api.docker.yaml` 后重启 API 服务：

```bash
docker-compose restart idrm-api
```

---

## 🔍 故障排查

### 服务无法启动

```bash
# 查看服务日志
docker-compose logs <service-name>

# 查看容器状态
docker-compose ps

# 检查端口占用
netstat -tulpn | grep <port>
```

### Elasticsearch 启动失败

```bash
# 检查内存限制
docker logs idrm-elasticsearch

# 如果提示内存不足，调整 docker-compose.yml:
elasticsearch:
  environment:
    - "ES_JAVA_OPTS=-Xms256m -Xmx256m"  # 降低内存
```

### MySQL 连接失败

```bash
# 检查 MySQL 是否就绪
docker-compose exec mysql mysqladmin ping -h localhost

# 查看 MySQL 日志
docker-compose logs mysql

# 测试连接
docker-compose exec mysql mysql -uroot -pidrm@2024 -e "SELECT 1"
```

### API 无法访问数据库

```bash
# 检查网络
docker network ls
docker network inspect idrm_idrm-network

# 在 API 容器内测试连接
docker-compose exec idrm-api ping mysql
```

---

## 📊 监控和维护

### 清理日志

```bash
# 清理 Docker 日志
docker system prune -a

# 清理应用日志（保留最近7天）
find ./logs -name "*.log" -mtime +7 -delete
```

### 备份数据

```bash
# 备份 MySQL 数据
docker-compose exec mysql mysqldump -uroot -pidrm@2024 --all-databases > backup.sql

# 备份 Docker 卷
docker run --rm -v idrm_mysql_data:/data -v $(pwd):/backup alpine \
  tar czf /backup/mysql_backup.tar.gz /data
```

### 更新服务

```bash
# 拉取最新镜像
docker-compose pull

# 重新构建 API
docker-compose build --no-cache idrm-api

# 重启服务
docker-compose up -d
```

---

## 🚦 生产环境建议

### 1. 安全加固

```yaml
# docker-compose.yml
elasticsearch:
  environment:
    - xpack.security.enabled=true
    - ELASTIC_PASSWORD=strong_password
```

### 2. 数据持久化

确保使用命名卷而不是绑定挂载：

```yaml
volumes:
  mysql_data:
    driver: local
```

### 3. 资源限制

```yaml
idrm-api:
  deploy:
    resources:
      limits:
        cpus: '1.0'
        memory: 512M
      reservations:
        cpus: '0.5'
        memory: 256M
```

### 4. 健康检查

所有关键服务都配置了健康检查，确保服务可用。

### 5. 日志轮转

应用日志自动轮转（最大10MB，保留3个文件）。

---

## 📚 参考资料

- [Docker Compose 文档](https://docs.docker.com/compose/)
- [go-zero 文档](https://go-zero.dev/)
- [Elasticsearch 文档](https://www.elastic.co/guide/en/elasticsearch/reference/current/index.html)
- [Jaeger 文档](https://www.jaegertracing.io/docs/)

---

## ❓ 常见问题

**Q: 如何扩展 API 服务？**

A: 使用 `docker-compose up -d --scale idrm-api=3` 启动3个实例。

**Q: 如何查看实时日志？**

A: 使用 `docker-compose logs -f --tail=100 idrm-api`

**Q: 如何进入容器内部？**

A: 使用 `docker-compose exec idrm-api sh`

**Q: 如何完全重置环境？**

A: 使用 `docker-compose down -v` 然后 `docker-compose up -d`

---

## ✅ 下一步

1. 访问 Kibana 配置日志查询
2. 访问 Jaeger 查看链路追踪
3. 测试 API 接口
4. 根据需要调整配置

祝您使用愉快！🎉
