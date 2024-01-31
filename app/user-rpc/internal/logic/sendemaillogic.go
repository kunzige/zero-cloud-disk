package logic

import (
	"context"
	"fmt"

	"zero-cloud-disk/app/user-rpc/internal/svc"
	"zero-cloud-disk/app/user-rpc/pb"
	"zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *pb.UserEmailReq) (*pb.UserEmailRes, error) {
	// todo: add your logic here and delete this line
	randomCode := utils.RandomCode()
	_, err := l.svcCtx.RedisCache.SetnxEx(in.Email, randomCode, 120)
	if err != nil {
		return nil, fmt.Errorf("发送验证码失败")
	}
	return &pb.UserEmailRes{Code: string(randomCode)}, nil
}
