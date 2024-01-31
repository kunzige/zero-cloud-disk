package logic

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/share/internal/svc"
	"zero-cloud-disk/app/share/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelShareFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelShareFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelShareFileLogic {
	return &CancelShareFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelShareFileLogic) CancelShareFile(req *types.CancelFileShareRequest) (resp *types.CancelFileShareResponse, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.ShareModel.DeleteByUuid(l.ctx, req.ShareUuid)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("取消分享失败:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &types.CancelFileShareResponse{Message: "取消分享成功"}, nil
}
