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
	sendEmailRes, err := l.svcCtx.UserRpcClient.SendEmail(l.ctx, &pb.UserEmailReq{
		Email: req.Email,
	})
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		return nil, fmt.Errorf("创建失败:%s\nLine:%d\n%v", file, line+1, err)
	}
	return &types.UserEmailResponse{Code: sendEmailRes.Code}, nil

}
