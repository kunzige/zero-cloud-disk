package handler

import (
	"io/ioutil"
	"net/http"
	"os"

	"zero-cloud-disk/app/files/internal/logic"
	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func fileDownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileDownloadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFileDownloadLogic(r.Context(), svcCtx)
		resp, err := l.FileDownload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			f, err := os.Open("." + resp.FilePath + resp.FileName)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
			}
			data, err := ioutil.ReadAll(f)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
			}
			w.Write(data)
		}
	}
}
