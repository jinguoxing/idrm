package handler

import (
	"idrm/api/internal/handler/data_understanding"
	"idrm/api/internal/handler/data_view"
	"idrm/api/internal/handler/resource_catalog"
	"idrm/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 注册资源目录模块
	resource_catalog.RegisterHandlers(server, serverCtx)

	// 注册数据视图模块
	data_view.RegisterHandlers(server, serverCtx)

	// 注册数据理解模块
	data_understanding.RegisterHandlers(server, serverCtx)
}
