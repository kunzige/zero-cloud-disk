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

	data, err := l.svcCtx.ShareRpcClient.GetFileShare(l.ctx, &pb.GetShareReq{
		ShareUuid: req.ShareUuid,
	})
	res := make(map[string]interface{})
	res["user_name"] = data.SharedFiles.UserName
	res["path"] = data.SharedFiles.FileAddr
	res["file_name"] = data.SharedFiles.FileName
	res["file_size"] = data.SharedFiles.FileSize
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("获取分享链接失败:%s\nLine:%d\n%v", file, line+1, err)
	}

	return &types.GetShareResponse{Data: res}, nil
}
