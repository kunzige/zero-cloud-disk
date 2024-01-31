package file

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-cloud-disk/app/applet/internal/logic/file"
	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
)

func FileListByPathHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileListByPathRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := file.NewFileListByPathLogic(r.Context(), svcCtx)
		resp, err := l.FileListByPath(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
