package svc

import (
	"zero-cloud-disk/app/applet/internal/config"
	"zero-cloud-disk/app/file-rpc/filecenter"
	"zero-cloud-disk/app/share-rpc/sharecenter"
	"zero-cloud-disk/app/user-rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	UserRpcClient  usercenter.Usercenter
	FileRpcClient  filecenter.Filecenter
	ShareRpcClient sharecenter.Sharecenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserRpcClient:  usercenter.NewUsercenter(zrpc.MustNewClient(c.UserRpcConf)),
		FileRpcClient:  filecenter.NewFilecenter(zrpc.MustNewClient(c.FileRpcConf)),
		ShareRpcClient: sharecenter.NewSharecenter(zrpc.MustNewClient(c.ShareRpcConf)),
	}
}
