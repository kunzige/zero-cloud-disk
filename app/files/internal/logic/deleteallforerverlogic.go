package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAllForerverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAllForerverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAllForerverLogic {
	return &DeleteAllForerverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteAllForerverLogic) DeleteAllForerver(req *types.DeleteAllForeverRequest) (resp *types.DeleteAllForeverResponse, err error) {
	// todo: add your logic here and delete this line
	// 清空回收站
	err = l.svcCtx.TbUserFileModel.DeleteAllForever(l.ctx, req.UserEmail)
	if err != nil {
		return nil, fmt.Errorf("删除失败")
	}
	return &types.DeleteAllForeverResponse{Message: "删除成功"}, nil
}
