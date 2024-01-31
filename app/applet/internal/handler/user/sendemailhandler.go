package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zero-cloud-disk/app/applet/internal/logic/user"
	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
)

func SendEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserEmailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSendEmailLogic(r.Context(), svcCtx)
		resp, err := l.SendEmail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
