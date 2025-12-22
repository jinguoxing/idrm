// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"idrm/api/internal/logic/data_view/category"
	"idrm/api/internal/svc"
	"idrm/api/internal/types"
)

// 类别列表
func ListCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewListCategoryLogic(r.Context(), svcCtx)
		resp, err := l.ListCategory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
