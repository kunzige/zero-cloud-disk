package handler

import (
	"fmt"
	"net/http"

	"zero-cloud-disk/app/files/internal/logic"
	"zero-cloud-disk/app/files/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	defaultMultipartMemory = 32 << 20 // 32 MB
)

func fileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 就是限制一下解析表单时的内存大小。
		if err := r.ParseMultipartForm(defaultMultipartMemory); err != nil {
			fmt.Println(err)
			httpx.Error(w, err)
			return
		}

		// uploadedFiles := r.MultipartForm.File

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)

		// 多文件上传
		// for key, files := range uploadedFiles {
		// 	if strings.HasPrefix(key, "file") {
		// 		l.Files = append(l.Files, files...)
		// 	}
		// }

		// 单文件上传
		l.Files = r.MultipartForm.File["file"]

		resp, err := l.FileUpload()
		if err != nil {
			fmt.Println(err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
