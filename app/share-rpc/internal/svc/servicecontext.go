package svc

import (
	"zero-cloud-disk/app/share-rpc/internal/config"
	"zero-cloud-disk/app/share/models"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	ShareModel models.TbShareModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ShareModel: models.NewTbShareModel(sqlx.NewMysql(c.Mysql.DataSource)),
	}
}
