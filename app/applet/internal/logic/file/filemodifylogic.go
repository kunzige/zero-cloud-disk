package file

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileModifyLogic {
	return &FileModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileModifyLogic) FileModify(req *types.FileModifyRequest) (resp *types.FileModifyResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.FileRpcClient.FileModify(l.ctx, &pb.FileModifyReq{FileHash: req.FileHash, NewFileName: req.NewFileName})
	if err != nil {
		return nil, fmt.Errorf("更新文件信息失败")
	}
	return &types.FileModifyResponse{Message: "ok"}, nil
}
