package data_view

import (
	"idrm/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// TODO: 添加数据视图模块路由
	server.AddRoutes(
		[]rest.Route{
			// 待添加
		},
	)
}
