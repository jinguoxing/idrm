package handler

import (
	"net/http"

	"{{MODULE_PATH}}/api/internal/logic"
	"{{MODULE_PATH}}/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// HealthCheckHandler 健康检查处理器
func HealthCheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHealthCheckLogic(r.Context(), svcCtx)
		resp, err := l.HealthCheck()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
