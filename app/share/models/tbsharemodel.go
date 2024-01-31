package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TbShareModel = (*customTbShareModel)(nil)

type (
	// TbShareModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbShareModel.
	TbShareModel interface {
		tbShareModel
		withSession(session sqlx.Session) TbShareModel
	}

	customTbShareModel struct {
		*defaultTbShareModel
	}
)

// NewTbShareModel returns a model for the database table.
func NewTbShareModel(conn sqlx.SqlConn) TbShareModel {
	return &customTbShareModel{
		defaultTbShareModel: newTbShareModel(conn),
	}
}

func (m *customTbShareModel) withSession(session sqlx.Session) TbShareModel {
	return NewTbShareModel(sqlx.NewSqlConnFromSession(session))
}
