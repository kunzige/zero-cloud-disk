package svc

import (
	"zero-cloud-disk/app/share-rpc/internal/config"
	"zero-cloud-disk/app/share/models"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	ShareModel      models.TbShareModel
	TbUserFileModel models.TbUserFileModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		ShareModel:      models.NewTbShareModel(sqlx.NewMysql(c.Mysql.DataSource)),
		TbUserFileModel: models.NewTbUserFileModel(sqlx.NewSqlConn("mysql", c.Mysql.DataSource)),
	}
}
