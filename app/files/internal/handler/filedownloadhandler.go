package handler

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"

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
		_ = resp
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// f, err := os.Open("." + resp.FilePath + resp.FileName)
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
			}
			// data, err := ioutil.ReadAll(f)

			// 分片下载

			// 后缀名
			fileExt := filepath.Ext(resp.FileName)
			// 不包括后缀
			fileWithoutExt := filepath.Base(resp.FileName[:len(resp.FileName)-len(fileExt)])
			chunks := 1 << 20
			block := int(math.Ceil(float64(resp.FileSize / int64(chunks))))
			// 识别不同类型的文件
			w.Header().Add("Transfer-Encoding", "chunked")
			w.Header().Set("Content-Type", getMimeType(fileExt))
			w.Header().Set("Content-Disposition", "attachment; filename="+resp.FileName)
			// 结合上传的逻辑，向上取整，这里需要取等于
			for i := 0; i <= block; i++ {
				full_path := fmt.Sprintf(".%s%s_%d%s", resp.FilePath, fileWithoutExt, i, fileExt)
				f, _ := os.Open(full_path)
				defer f.Close()
				io.Copy(w, f)
				w.(http.Flusher).Flush()
			}
			if err != nil {
				httpx.ErrorCtx(r.Context(), w, err)
			}
			// w.Write(data)
		}
	}
}

// 根据文件扩展名获取对应的 MIME 类型
func getMimeType(fileExt string) string {
	switch fileExt {
	case ".mp4":
		return "video/mp4"
	case ".txt":
		return "text/plain"
	// 添加其他文件类型的处理
	default:
		return "application/octet-stream"
	}
}
