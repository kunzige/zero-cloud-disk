package logic

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/share/internal/svc"
	"zero-cloud-disk/app/share/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShareFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetShareFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShareFileLogic {
	return &GetShareFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetShareFileLogic) GetShareFile(req *types.GetShareRequest) (resp *types.GetShareResponse, err error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.ShareModel.FindByUuid(l.ctx, req.ShareUuid)
	res := make(map[string]interface{})
	res["user_name"] = data.UserName
	res["path"] = data.Path
	res["file_name"] = data.FileName
	res["file_size"] = data.FileSize
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("获取分享链接失败:%s\nLine:%d\n%v", file, line+1, err)
	}

	return &types.GetShareResponse{Data: res}, nil
}
