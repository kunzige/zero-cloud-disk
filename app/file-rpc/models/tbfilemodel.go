package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TbFileModel = (*customTbFileModel)(nil)

type (
	// TbFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbFileModel.
	TbFileModel interface {
		tbFileModel
		withSession(session sqlx.Session) TbFileModel
	}

	customTbFileModel struct {
		*defaultTbFileModel
	}
)

// NewTbFileModel returns a model for the database table.
func NewTbFileModel(conn sqlx.SqlConn) TbFileModel {
	return &customTbFileModel{
		defaultTbFileModel: newTbFileModel(conn),
	}
}

func (m *customTbFileModel) withSession(session sqlx.Session) TbFileModel {
	return NewTbFileModel(sqlx.NewSqlConnFromSession(session))
}
