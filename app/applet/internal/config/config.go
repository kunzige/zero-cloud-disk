package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf  zrpc.RpcClientConf
	FileRpcConf  zrpc.RpcClientConf
	ShareRpcConf zrpc.RpcClientConf
	Auth         struct {
		AccessSecret string
		AccessExpire int64
	}
}
