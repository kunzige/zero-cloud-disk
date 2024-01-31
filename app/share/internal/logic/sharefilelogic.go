package logic

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/share/internal/svc"
	"zero-cloud-disk/app/share/internal/types"
	"zero-cloud-disk/app/share/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareFileLogic {
	return &ShareFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareFileLogic) ShareFile(req *types.FileShareRequest) (resp *types.FileShareResponse, err error) {
	// todo: add your logic here and delete this line
	data := models.ShareData{
		UserName:      req.UserName,
		UserEmail:     req.UserEmail,
		ShareIdentity: req.FileHash,
		Path:          req.Path,
		ShareUuid:     req.ShareUuid,
	}
	_, err = l.svcCtx.ShareModel.Share(l.ctx, &data)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("创建失败:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &types.FileShareResponse{Message: "分享成功"}, nil
}
