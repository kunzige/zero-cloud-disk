package logic

import (
	"context"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareLogic {
	return &GetShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareLogic) GetShare(req *types.GetShareRequest) (resp *types.GetShareResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
