package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/file-rpc/internal/svc"
	"zero-cloud-disk/app/file-rpc/models"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDownloadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDownloadLogic {
	return &FileDownloadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDownloadLogic) FileDownload(in *pb.FileDownloadReq) (*pb.FileDownloadRes, error) {
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
	return &pb.FileDownloadRes{Message: "ok", FileName: fileMeta.FileName, FilePath: fileMeta.FileAddr, FileSize: fileMeta.FileSize}, nil
}
