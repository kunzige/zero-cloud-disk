package logic

import (
	"context"
	"fmt"
	"runtime"

	"zero-cloud-disk/app/user-rpc/internal/svc"
	"zero-cloud-disk/app/user-rpc/pb"
	"zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *pb.UserLoginReq) (*pb.UserLoginRes, error) {
	// todo: add your logic here and delete this line

	// 校验用户名与密码是否正确
	user, err := l.svcCtx.TbUserModel.FindOneByEmail(l.ctx, in.UserEmail)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("错误:%s\nLine:%d\n%v", file, line+1, err)
	}
	if user.UserPwd != utils.Encrypt(in.UserPassword) {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("密码错误:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &pb.UserLoginRes{Message: "登陆成功"}, nil
}
