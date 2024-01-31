package logic

import (
	"context"
	"fmt"
	"time"

	"zero-cloud-disk/app/user-rpc/internal/svc"
	"zero-cloud-disk/app/user-rpc/pb"
	"zero-cloud-disk/app/user/models"
	"zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *pb.UserRegisterReq) (*pb.UserRegisterRes, error) {
	// todo: add your logic here and delete this line
	// 验证验证码是否正确

	code, err := l.svcCtx.RedisCache.Get(in.Email)
	if code != in.Code {
		return nil, fmt.Errorf("验证码不正确")
	}

	// 获取注册的用户名和密码并写入数据库
	_, err = l.svcCtx.TbUserModel.Insert(l.ctx, &models.TbUser{UserName: in.UserName, UserPwd: utils.Encrypt(in.UserPassword), Email: in.Email, SignupAt: time.Now(), LastActive: time.Now()})
	if err != nil {
		return nil, fmt.Errorf("注册失败")
	}

	return &pb.UserRegisterRes{Message: "注册成功"}, nil
}
