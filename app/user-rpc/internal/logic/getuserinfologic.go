package logic

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/user-rpc/internal/svc"
	"zero-cloud-disk/app/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.UserInfoReq) (*pb.UserInfoRes, error) {
	// todo: add your logic here and delete this line
	// 获取用户信息
	user, err := l.svcCtx.TbUserModel.FindOneByEmail(l.ctx, in.UserEmail)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("错误:%s\nLine:%d\n%v", file, line+1, err)
	}
	data := make(map[string]string)
	data["user_email"] = user.Email
	data["user_name"] = user.UserName
	data["last_active"] = user.LastActive.Format("2006-01-02 15:04:05")
	return &pb.UserInfoRes{Data: data}, nil
}
