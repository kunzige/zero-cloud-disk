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

type SaveFileShareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveFileShareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveFileShareLogic {
	return &SaveFileShareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveFileShareLogic) SaveFileShare(in *pb.SaveFileReq) (*pb.SaveFileRes, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.TbUserFileModel.SaveFile(l.ctx, &models.SaveData{UserName: in.UserName, FileName: in.FileName, FileSize: in.FileSize, ShareIdentity: in.ShareIdentity, UserEmail: in.UserEmail, Type: in.Type})

	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("错误:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &pb.SaveFileRes{Message: "转存成功"}, nil
}
