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

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	// todo: add your logic here and delete this line

	_, err = l.svcCtx.UserRpcClient.UserRegister(l.ctx, &pb.UserRegisterReq{UserName: req.UserName, UserPassword: req.UserPassword,
		Email: req.Email, Code: req.Code,
	})

	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("注册失败:%s\nLine:%d\n%v", file, line+1, err)
	}

	return &types.UserRegisterResponse{Message: "注册成功"}, err
}
