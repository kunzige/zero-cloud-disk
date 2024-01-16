package svc

import (
	"zero-cloud-disk/app/files/internal/config"
	"zero-cloud-disk/app/files/models"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	TbFileModel models.TbFileModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		TbFileModel: models.NewTbFileModel(sqlx.NewSqlConn("mysql", c.DataSource)),
	}
}
