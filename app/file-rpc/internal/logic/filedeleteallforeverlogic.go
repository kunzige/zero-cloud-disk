package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/file-rpc/internal/svc"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileDeleteAllForeverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileDeleteAllForeverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileDeleteAllForeverLogic {
	return &FileDeleteAllForeverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileDeleteAllForeverLogic) FileDeleteAllForever(in *pb.DeleteAllForeverReq) (*pb.DeleteAllForeverRes, error) {
	// todo: add your logic here and delete this line
	// 清空回收站
	err := l.svcCtx.TbUserFileModel.DeleteAllForever(l.ctx, in.UserEmail)
	if err != nil {
		return nil, fmt.Errorf("删除失败")
	}
	return &pb.DeleteAllForeverRes{Message: "删除成功"}, nil
}
