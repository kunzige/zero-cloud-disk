package logic

import (
	"context"
	"fmt"
	"time"

	"zero-cloud-disk/app/user/internal/svc"
	"zero-cloud-disk/app/user/internal/types"
	"zero-cloud-disk/app/user/models"
	"zero-cloud-disk/utils"

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

	// 验证验证码是否正确

	code, err := l.svcCtx.RedisCache.Get(req.Email)
	if code != req.Code {
		return nil, fmt.Errorf("验证码不正确")
	}

	// 获取注册的用户名和密码并写入数据库
	_, err = l.svcCtx.TbUserModel.Insert(l.ctx, &models.TbUser{UserName: req.UserName, UserPwd: utils.Encrypt(req.UserPassword), Email: req.Email, SignupAt: time.Now(), LastActive: time.Now()})
	if err != nil {
		return nil, fmt.Errorf("注册失败")
	}

	return &types.UserRegisterResponse{Message: "注册成功"}, nil
}
