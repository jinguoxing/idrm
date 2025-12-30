package svc

import (
	"{{MODULE_PATH}}/api/internal/config"
)

// ServiceContext 服务上下文
// 包含所有服务依赖，如Model、缓存、配置等
type ServiceContext struct {
	Config config.Config
	// TODO: 添加Model依赖
	// CategoryModel category.Model
}

// NewServiceContext 创建服务上下文
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
