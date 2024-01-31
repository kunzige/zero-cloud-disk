package file

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
	"zero-cloud-disk/app/file-rpc/pb"

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
	_, err = l.svcCtx.FileRpcClient.FileDeleteAllForever(l.ctx, &pb.DeleteAllForeverReq{UserEmail: req.UserEmail})
	if err != nil {
		return nil, fmt.Errorf("删除失败")
	}
	return &types.DeleteAllForeverResponse{Message: "删除成功"}, nil
}
