package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/file-rpc/internal/svc"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDeleteForeverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDeleteForeverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeleteForeverLogic {
	return &FileDeleteForeverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDeleteForeverLogic) FileDeleteForever(in *pb.DeleteForeverReq) (*pb.DeleteForeverRes, error) {
	// todo: add your logic here and delete this line
	// 永久删除文件
	err := l.svcCtx.TbUserFileModel.DeleteByHashForever(l.ctx, in.Hash)
	if err != nil {
		return nil, fmt.Errorf("删除失败")
	}
	return &pb.DeleteForeverRes{Message: "删除成功"}, nil
}
