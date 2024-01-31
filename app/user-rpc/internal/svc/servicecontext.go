package svc

import (
	"zero-cloud-disk/app/user-rpc/internal/config"
	"zero-cloud-disk/app/user/models"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	TbUserModel models.TbUserModel
	RedisCache  redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		TbUserModel: models.NewTbUserModel(sqlx.NewSqlConn("mysql", c.Mysql.DataSource)),
		RedisCache:  *redis.New(c.CacheRedis.Addr, redis.WithPass(c.CacheRedis.Pass)),
	}
}
