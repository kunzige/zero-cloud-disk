package share

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
	"zero-cloud-disk/app/share-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelFileShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelFileShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelFileShareLogic {
	return &CancelFileShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelFileShareLogic) CancelFileShare(req *types.CancelFileShareRequest) (resp *types.CancelFileShareResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.ShareRpcClient.CancelFileShare(l.ctx, &pb.CancelFileShareReq{
		ShareUuid: req.ShareUuid,
	})
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("取消分享失败:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &types.CancelFileShareResponse{Message: "取消分享成功"}, nil
}
