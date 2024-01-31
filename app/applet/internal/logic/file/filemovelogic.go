package file

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
	"zero-cloud-disk/app/file-rpc/pb"
	"zero-cloud-disk/app/user/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileMoveLogic {
	return &FileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileMoveLogic) FileMove(req *types.FileMoveRequest) (resp *types.FileMoveResponse, err error) {
	// todo: add your logic here and delete this line
	// 更改对应路径即可
	_, err = l.svcCtx.FileRpcClient.FileMove(l.ctx, &pb.FileMoveReq{FileHash: req.FileHash, NewPath: req.NewPath})
	if err != nil && err != models.ErrNotFound {
		return nil, fmt.Errorf("移动失败")
	}
	if err == models.ErrNotFound {
		return nil, fmt.Errorf("文件不存在")
	}
	return &types.FileMoveResponse{Message: "移动成功"}, nil
}
