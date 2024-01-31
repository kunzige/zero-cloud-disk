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

type FileShareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileShareLogic {
	return &FileShareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileShareLogic) FileShare(req *types.FileShareRequest) (resp *types.FileShareResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.ShareRpcClient.FileShare(l.ctx, &pb.FileShareReq{
		FileHash:  req.FileHash,
		UserEmail: req.UserEmail,
		UserName:  req.UserName,
		Path:      req.Path,
		ShareUuid: req.ShareUuid,
	})
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("创建失败:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &types.FileShareResponse{Message: "分享成功"}, nil
}
