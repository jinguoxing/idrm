-- 创建数据库
CREATE DATABASE IF NOT EXISTS `idrm_resource_catalog` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `idrm_data_view` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE DATABASE IF NOT EXISTS `idrm_data_understanding` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 使用数据库并创建category表（示例）
USE `idrm_resource_catalog`;

CREATE TABLE IF NOT EXISTS `category` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(100) NOT NULL COMMENT '名称',
  `code` varchar(50) NOT NULL COMMENT '编码',
  `parent_id` bigint DEFAULT '0' COMMENT '父级ID',
  `level` int DEFAULT '1' COMMENT '层级',
  `sort` int DEFAULT '0' COMMENT '排序',
  `description` text COMMENT '描述',
  `status` int DEFAULT '1' COMMENT '状态(1:启用 0:禁用)',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`),
  KEY `idx_parent_id` (`parent_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='类别表';

-- 插入测试数据
INSERT INTO `category` (`name`, `code`, `parent_id`, `level`, `sort`, `description`, `status`) VALUES
('根类别', 'ROOT', 0, 1, 0, '顶级类别', 1),
('子类别1', 'CHILD1', 1, 2, 1, '第一个子类别', 1),
('子类别2', 'CHILD2', 1, 2, 2, '第二个子类别', 1);
