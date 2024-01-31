package file

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
	"zero-cloud-disk/app/file-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFolderLogic {
	return &CreateFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFolderLogic) CreateFolder(req *types.CreateFolderRequest) (resp *types.CreateFolderResponse, err error) {
	// todo: add your logic here and delete this line
	// 将用户的文件夹添加到数据库表中

	_, err = l.svcCtx.FileRpcClient.CreateFolder(l.ctx, &pb.CreateFolderReq{UserEmail: req.UserEmail, UserName: req.UserName, PathName: req.PathName, ParentDir: req.ParentDir})
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("创建失败:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &types.CreateFolderResponse{Message: "创建成功"}, nil
}
