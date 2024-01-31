package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/file-rpc/internal/svc"
	"zero-cloud-disk/app/file-rpc/models"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileInfoLogic {
	return &GetFileInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileInfoLogic) GetFileInfo(in *pb.FileInfoReq) (*pb.FileInfoRes, error) {
	// todo: add your logic here and delete this line

	// todo: add your logic here and delete this line
	fileMeta, err := l.svcCtx.TbFileModel.FindOneByFileSha1(l.ctx, in.FileHash)
	// 内部错误
	if err != nil && err != models.ErrNotFound {
		return nil, fmt.Errorf("下载失败")
	}
	// 不存在
	if err == models.ErrNotFound {
		return nil, fmt.Errorf("下载失败")
	}

	return &pb.FileInfoRes{UpdateAt: fileMeta.UpdateAt, FileName: fileMeta.FileName, FileSize: fileMeta.FileSize, CreateAt: fileMeta.CreateAt}, nil
}
