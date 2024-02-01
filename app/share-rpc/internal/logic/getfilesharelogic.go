package logic

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/share-rpc/internal/svc"
	"zero-cloud-disk/app/share-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileShareLogic {
	return &GetFileShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFileShareLogic) GetFileShare(in *pb.GetShareReq) (*pb.SharedFiles, error) {
	// todo: add your logic here and delete this line
	data, err := l.svcCtx.ShareModel.FindByUuid(l.ctx, in.ShareUuid)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("获取分享链接失败:%s\nLine:%d\n%v", file, line+1, err)
	}
	sharedFile := &pb.SharedFile{UserName: data.UserName, FileAddr: data.Path, FileSize: data.FileSize, FileName: data.FileName, SharedIdentity: data.ShareIdentity, Type: data.Type}
	return &pb.SharedFiles{SharedFiles: sharedFile}, nil
}
