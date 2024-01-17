// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"zero-cloud-disk/app/files/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/file/download",
				Handler: fileDownloadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/file/info",
				Handler: fileInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/modify",
				Handler: fileModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file/upload",
				Handler: fileUploadHandler(serverCtx),
			},
		},
	)
}
