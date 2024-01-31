// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	file "zero-cloud-disk/app/applet/internal/handler/file"
	share "zero-cloud-disk/app/applet/internal/handler/share"
	user "zero-cloud-disk/app/applet/internal/handler/user"
	"zero-cloud-disk/app/applet/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/createfolder",
				Handler: file.CreateFolderHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: file.FileDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/deleteallforever",
				Handler: file.DeleteAllForerverHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/deleteforever",
				Handler: file.DeleteForerverHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/download",
				Handler: file.FileDownloadHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: file.FileInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/modify",
				Handler: file.FileModifyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/move",
				Handler: file.FileMoveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/pathlist",
				Handler: file.FileListByPathHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/recycle",
				Handler: file.RecycleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/upload",
				Handler: file.FileUploadHandler(serverCtx),
			},
		},
		rest.WithPrefix("/file"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodDelete,
				Path:    "/cancel",
				Handler: share.CancelFileShareHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/file",
				Handler: share.FileShareHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get",
				Handler: share.GetShareHandler(serverCtx),
			},
		},
		rest.WithPrefix("/share"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/info",
				Handler: user.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/sendemail",
				Handler: user.SendEmailHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user"),
	)
}
