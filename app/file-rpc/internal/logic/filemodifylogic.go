package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/file-rpc/internal/svc"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileModifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileModifyLogic {
	return &FileModifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileModifyLogic) FileModify(in *pb.FileModifyReq) (*pb.FileModifyRes, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.TbFileModel.UpdateBySha1(l.ctx, in.FileHash, in.NewFileName)
	if err != nil {
		return nil, fmt.Errorf("更新文件信息失败")
	}
	return &pb.FileModifyRes{Message: "ok"}, nil
}
