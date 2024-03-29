// Code generated by goctl. DO NOT EDIT.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tbUserFieldNames          = builder.RawFieldNames(&TbUser{})
	tbUserRows                = strings.Join(tbUserFieldNames, ",")
	tbUserRowsExpectAutoSet   = strings.Join(stringx.Remove(tbUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tbUserRowsWithPlaceHolder = strings.Join(stringx.Remove(tbUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tbUserModel interface {
		Insert(ctx context.Context, data *TbUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TbUser, error)
		FindOneByEmail(ctx context.Context, email string) (*TbUser, error)
		FindOneByPhone(ctx context.Context, phone sql.NullInt64) (*TbUser, error)
		Update(ctx context.Context, data *TbUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTbUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TbUser struct {
		Id         int64          `db:"id"`          // id
		UserName   string         `db:"user_name"`   // 用户名
		UserPwd    string         `db:"user_pwd"`    // 密码
		Email      string         `db:"email"`       // 邮件
		Phone      sql.NullInt64  `db:"phone"`       // 手机号
		SignupAt   time.Time      `db:"signup_at"`   // 注册时间
		LastActive time.Time      `db:"last_active"` // 最近活跃
		Profile    sql.NullString `db:"profile"`     // 用户属性
		Status     int64          `db:"status"`      // 标记账号的状态
	}
)

func newTbUserModel(conn sqlx.SqlConn) *defaultTbUserModel {
	return &defaultTbUserModel{
		conn:  conn,
		table: "`tb_user`",
	}
}

func (m *defaultTbUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTbUserModel) FindOne(ctx context.Context, id int64) (*TbUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbUserRows, m.table)
	var resp TbUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbUserModel) FindOneByEmail(ctx context.Context, email string) (*TbUser, error) {
	var resp TbUser
	query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", tbUserRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, email)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbUserModel) FindOneByPhone(ctx context.Context, phone sql.NullInt64) (*TbUser, error) {
	var resp TbUser
	query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", tbUserRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, phone)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbUserModel) Insert(ctx context.Context, data *TbUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, tbUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.UserName, data.UserPwd, data.Email, data.Phone, data.SignupAt, data.LastActive, data.Profile, data.Status)
	return ret, err
}

func (m *defaultTbUserModel) Update(ctx context.Context, newData *TbUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tbUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.UserName, newData.UserPwd, newData.Email, newData.Phone, newData.SignupAt, newData.LastActive, newData.Profile, newData.Status, newData.Id)
	return err
}

func (m *defaultTbUserModel) tableName() string {
	return m.table
}
