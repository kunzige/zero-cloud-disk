package logic

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/share-rpc/internal/svc"
	"zero-cloud-disk/app/share-rpc/pb"
	"zero-cloud-disk/app/share/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FileShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileShareLogic {
	return &FileShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileShareLogic) FileShare(in *pb.FileShareReq) (*pb.FileShareRes, error) {
	// todo: add your logic here and delete this line
	data := models.ShareData{
		UserName:      in.UserName,
		UserEmail:     in.UserEmail,
		ShareIdentity: in.FileHash,
		Path:          in.Path,
		ShareUuid:     in.ShareUuid,
	}
	_, err := l.svcCtx.ShareModel.Share(l.ctx, &data)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("分享失败:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &pb.FileShareRes{Message: "分享成功"}, nil
}
