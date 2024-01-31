package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-cloud-disk/app/files/internal/logic"
	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"
)

func deleteAllForerverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteAllForeverRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteAllForerverLogic(r.Context(), svcCtx)
		resp, err := l.DeleteAllForerver(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
