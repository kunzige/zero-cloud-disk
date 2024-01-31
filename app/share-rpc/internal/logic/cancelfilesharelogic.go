package logic

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/share-rpc/internal/svc"
	"zero-cloud-disk/app/share-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelFileShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelFileShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelFileShareLogic {
	return &CancelFileShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CancelFileShareLogic) CancelFileShare(in *pb.CancelFileShareReq) (*pb.CancelFileShareRes, error) {
	// todo: add your logic here and delete this line
	err := l.svcCtx.ShareModel.DeleteByUuid(l.ctx, in.ShareUuid)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("取消分享失败:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &pb.CancelFileShareRes{Message: "取消分享成功"}, nil
}
