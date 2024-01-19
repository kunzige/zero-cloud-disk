package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/user/internal/svc"
	"zero-cloud-disk/app/user/internal/types"
	"zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendEmailLogic) SendEmail(req *types.UserEmailRequest) (resp *types.UserEmailResponse, err error) {
	// todo: add your logic here and delete this line
	randomCode := utils.RandomCode()
	_, err = l.svcCtx.RedisCache.SetnxEx(req.Email, randomCode, 120)
	if err != nil {
		return nil, fmt.Errorf("发送验证码失败")
	}
	return &types.UserEmailResponse{Code: string(randomCode)}, nil
}
