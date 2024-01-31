package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	// FileRpcConf zrpc.RpcClientConf
	Mysql struct {
		DataSource string
	}
}
