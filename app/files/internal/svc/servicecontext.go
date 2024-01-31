package svc

import (
	"zero-cloud-disk/app/files/internal/config"
	"zero-cloud-disk/app/files/models"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config          config.Config
	TbFileModel     models.TbFileModel
	TbUserFileModel models.TbUserFileModel
	// FileRpc         filecenter.Filecenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		TbFileModel:     models.NewTbFileModel(sqlx.NewSqlConn("mysql", c.Mysql.DataSource)),
		TbUserFileModel: models.NewTbUserFileModel(sqlx.NewSqlConn("mysql", c.Mysql.DataSource)),
		// FileRpc:         filecenter.NewFilecenter(zrpc.MustNewClient(c.FileRpcConf)),
	}
}
