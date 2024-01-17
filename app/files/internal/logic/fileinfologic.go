package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"
	"zero-cloud-disk/app/files/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileInfoLogic {
	return &FileInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileInfoLogic) FileInfo(req *types.FileInfoRequest) (resp *types.FileInfoResponse, err error) {
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
	return &types.FileInfoResponse{UpdateAt: fileMeta.UpdateAt, FileName: fileMeta.FileName, FileSize: fileMeta.FileSize, CreateAt: fileMeta.CreateAt}, nil
}
