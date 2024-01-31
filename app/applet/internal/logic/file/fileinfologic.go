package file

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
	"zero-cloud-disk/app/file-rpc/pb"

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

	// 通过rpc调用
	fileInfo, err := l.svcCtx.FileRpcClient.GetFileInfo(l.ctx, &pb.FileInfoReq{FileHash: req.FileHash})
	if err != nil {
		return nil, fmt.Errorf("下载失败")
	}
	return &types.FileInfoResponse{UpdateAt: fileInfo.UpdateAt, FileName: fileInfo.FileName, FileSize: fileInfo.FileSize, CreateAt: fileInfo.CreateAt}, nil
}
