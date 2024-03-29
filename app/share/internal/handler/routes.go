// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"zero-cloud-disk/app/share/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/file/share",
				Handler: shareFileHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/file/share/cancel",
				Handler: cancelShareFileHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/file/share/get",
				Handler: getShareFileHandler(serverCtx),
			},
		},
	)
}
