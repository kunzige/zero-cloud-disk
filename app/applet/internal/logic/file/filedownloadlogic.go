package file

import (
	"context"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDownloadLogic {
	return &FileDownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileDownloadLogic) FileDownload(req *types.FileDownloadRequest) (resp *types.FileDownloadResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
