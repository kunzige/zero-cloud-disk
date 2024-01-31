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
	// 单体服务
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

	// // 通过rpc调用
	// fileInfo, err := l.svcCtx.FileRpc.GetFileInfo(l.ctx, &pb.FileInfoReq{FileHash: req.FileHash})
	// if err != nil {
	// 	return nil, fmt.Errorf("下载失败")
	// }
	// return &types.FileInfoResponse{UpdateAt: fileInfo.UpdateAt, FileName: fileInfo.FileName, FileSize: fileInfo.FileSize, CreateAt: fileInfo	.CreateAt}, nil

}
