package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/files/internal/svc"
	"zero-cloud-disk/app/files/internal/types"
	"zero-cloud-disk/utils"

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
	message, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("上传文件失败")
	}

	encrypt_hash := utils.Encrypt(string(message))

	err = l.svcCtx.TbUserFileModel.InsertPath(l.ctx, req.UserEmail, req.UserName, req.PathName, req.ParentDir, encrypt_hash)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("创建失败:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &types.CreateFolderResponse{Message: "创建成功"}, nil
}
