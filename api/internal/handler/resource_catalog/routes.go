package resource_catalog

import (
	"net/http"

	"idrm/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/catalog/categories/:id",
				Handler: GetCategoryHandler(serverCtx),
			},
			// TODO: 添加其他路由
			// {
			//     Method:  http.MethodPost,
			//     Path:    "/api/v1/catalog/categories",
			//     Handler: CreateCategoryHandler(serverCtx),
			// },
		},
	)
}
