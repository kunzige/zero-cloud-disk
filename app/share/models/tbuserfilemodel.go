package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TbUserFileModel = (*customTbUserFileModel)(nil)

type (
	// TbUserFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbUserFileModel.
	TbUserFileModel interface {
		tbUserFileModel
		withSession(session sqlx.Session) TbUserFileModel
	}

	customTbUserFileModel struct {
		*defaultTbUserFileModel
	}
)

// NewTbUserFileModel returns a model for the database table.
func NewTbUserFileModel(conn sqlx.SqlConn) TbUserFileModel {
	return &customTbUserFileModel{
		defaultTbUserFileModel: newTbUserFileModel(conn),
	}
}

func (m *customTbUserFileModel) withSession(session sqlx.Session) TbUserFileModel {
	return NewTbUserFileModel(sqlx.NewSqlConnFromSession(session))
}
