package handler

import (
	"net/http"

	"{{MODULE_PATH}}/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

// RegisterHandlers 注册路由处理器
func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/health",
				Handler: HealthCheckHandler(serverCtx),
			},
		},
	)
}
