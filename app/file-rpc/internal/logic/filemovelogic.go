package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/file-rpc/internal/svc"
	"zero-cloud-disk/app/file-rpc/models"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileMoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileMoveLogic {
	return &FileMoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileMoveLogic) FileMove(in *pb.FileMoveReq) (*pb.FileMoveRes, error) {
	// todo: add your logic here and delete this line
	// 更改对应路径即可
	err := l.svcCtx.TbFileModel.UpdatePathBySha1(l.ctx, in.FileHash, in.NewPath)
	if err != nil && err != models.ErrNotFound {
		return nil, fmt.Errorf("移动失败")
	}
	if err == models.ErrNotFound {
		return nil, fmt.Errorf("文件不存在")
	}
	return &pb.FileMoveRes{Message: "移动成功"}, nil
}
