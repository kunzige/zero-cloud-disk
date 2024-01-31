package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/file-rpc/internal/svc"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeleteLogic {
	return &FileDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDeleteLogic) FileDelete(in *pb.FileDeleteReq) (*pb.FileDeleteRes, error) {
	// todo: add your logic here and delete this line
	// 删除文件
	err := l.svcCtx.TbUserFileModel.DeleteByHash(l.ctx, in.FileHash)
	if err != nil {
		return nil, fmt.Errorf("删除失败")
	}
	return &pb.FileDeleteRes{Message: "删除成功"}, nil
}
