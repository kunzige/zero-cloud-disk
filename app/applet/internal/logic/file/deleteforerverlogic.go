package file

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteForerverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteForerverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteForerverLogic {
	return &DeleteForerverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteForerverLogic) DeleteForerver(req *types.DeleteForeverRequest) (resp *types.DeleteForeverResponse, err error) {
	// todo: add your logic here and delete this line
	// 永久删除文件
	_, err = l.svcCtx.FileRpcClient.FileDeleteForever(l.ctx, &pb.DeleteForeverReq{Hash: req.Hash})
	if err != nil {
		return nil, fmt.Errorf("删除失败")
	}
	return &types.DeleteForeverResponse{Message: "删除成功"}, nil
}
