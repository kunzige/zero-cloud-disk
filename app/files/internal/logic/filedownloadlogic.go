package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"
	"zero-cloud-disk/app/files/models"

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
	fileMeta, err := l.svcCtx.TbFileModel.FindOneByFileSha1(l.ctx, req.FileHash)
	// 内部错误
	if err != nil && err != models.ErrNotFound {
		return nil, fmt.Errorf("下载失败")
	}
	// 不存在
	if err == models.ErrNotFound {
		return nil, fmt.Errorf("下载失败")
	}
	return &types.FileDownloadResponse{Message: "ok", FileName: fileMeta.FileName, FilePath: fileMeta.FileAddr}, nil
}
