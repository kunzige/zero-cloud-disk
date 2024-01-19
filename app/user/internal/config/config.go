package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	CacheRedis struct {
		Addr string
		Pass string
	}
}
