package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/file-rpc/internal/svc"
	"zero-cloud-disk/app/file-rpc/pb"
	"zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFolderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFolderLogic {
	return &CreateFolderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateFolderLogic) CreateFolder(in *pb.CreateFolderReq) (*pb.CreateFolderRes, error) {

	// 将用户的文件夹添加到数据库表中
	message, err := json.Marshal(in)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("创建文件夹失败")
	}

	encrypt_hash := utils.Encrypt(string(message))

	err = l.svcCtx.TbUserFileModel.InsertPath(l.ctx, in.UserEmail, in.UserName, in.PathName, in.ParentDir, encrypt_hash)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("创建文件夹失败:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &pb.CreateFolderRes{Message: "创建成功"}, nil
}
