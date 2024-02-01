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

type SaveFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveFileLogic {
	return &SaveFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveFileLogic) SaveFile(req *types.SaveFileRequest) (resp *types.SaveFileResponse, err error) {
	// todo: add your logic here and delete this line

	_, err = l.svcCtx.ShareRpcClient.SaveFileShare(l.ctx, &pb.SaveFileReq{UserName: req.UserName, FileName: req.FileName, FileSize: req.FileSize, ShareIdentity: req.ShareIdentity, UserEmail: req.UserEmail, Type: req.Type})

	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("错误:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &types.SaveFileResponse{Message: "转存成功"}, nil
}
