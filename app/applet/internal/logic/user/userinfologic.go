package user

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/applet/internal/svc"
	"zero-cloud-disk/app/applet/internal/types"
	"zero-cloud-disk/app/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.UserInfoReq{UserEmail: req.UserEmail})
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("错误:%s\nLine:%d\n%v", file, line+1, err)
	}
	data := make(map[string]string)
	data["user_email"] = userInfo.Data["user_email"]
	data["user_name"] = userInfo.Data["user_name"]
	data["last_active"] = userInfo.Data["last_active"]

	return &types.UserInfoResponse{Data: data}, nil
}
