#!/bin/bash

# 初始化数据库脚本

echo "Creating databases..."

# 数据库配置
DB_HOST="127.0.0.1"
DB_PORT="3306"
DB_USER="root"
DB_PASS="password"

# 创建数据库
mysql -h${DB_HOST} -P${DB_PORT} -u${DB_USER} -p${DB_PASS} << EOF
CREATE DATABASE IF NOT EXISTS idrm_data_view DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS idrm_data_understanding DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS idrm_resource_catalog DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
EOF

echo "Databases created successfully!"

# 创建资源目录表
echo "Creating resource_catalog tables..."
mysql -h${DB_HOST} -P${DB_PORT} -u${DB_USER} -p${DB_PASS} idrm_resource_catalog << EOF
CREATE TABLE IF NOT EXISTS categories (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL COMMENT '类别名称',
    code VARCHAR(50) NOT NULL UNIQUE COMMENT '类别编码',
    parent_id BIGINT DEFAULT 0 COMMENT '父级ID',
    level INT DEFAULT 1 COMMENT '层级',
    sort INT DEFAULT 0 COMMENT '排序',
    description TEXT COMMENT '描述',
    status TINYINT DEFAULT 1 COMMENT '状态：1启用 0禁用',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_parent_id (parent_id),
    INDEX idx_code (code),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源类别表';
EOF

echo "Tables created successfully!"
echo "Database initialization completed!"
