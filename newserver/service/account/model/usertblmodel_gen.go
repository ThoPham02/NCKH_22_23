// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userTblFieldNames          = builder.RawFieldNames(&UserTbl{}, true)
	userTblRows                = strings.Join(userTblFieldNames, ",")
	userTblRowsExpectAutoSet   = strings.Join(stringx.Remove(userTblFieldNames), ",")
	userTblRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(userTblFieldNames, "id"))
)

type (
	userTblModel interface {
		Insert(ctx context.Context, data *UserTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UserTbl, error)
		Update(ctx context.Context, data *UserTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UserTbl struct {
		Id           int64          `db:"id"`
		Username     string         `db:"username"`
		HashPassword string         `db:"hash_password"`
		Role         int64          `db:"role"`
		Name         string         `db:"name"`
		Email        sql.NullString `db:"email"`
		Phone        sql.NullString `db:"phone"`
		FacultyId    int64          `db:"faculty_id"`
		YearStart    int64          `db:"year_start"`
		Degree       int64          `db:"degree"`
		AvataUrl     sql.NullString `db:"avata_url"`
		Birthday     sql.NullString `db:"birthday"`
		BankAccount  sql.NullString `db:"bank_account"`
		Department   sql.NullInt64  `db:"department"`
	}
)

func newUserTblModel(conn sqlx.SqlConn) *defaultUserTblModel {
	return &defaultUserTblModel{
		conn:  conn,
		table: `"public"."user_tbl"`,
	}
}

func (m *defaultUserTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUserTblModel) FindOne(ctx context.Context, id int64) (*UserTbl, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", userTblRows, m.table)
	var resp UserTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserTblModel) Insert(ctx context.Context, data *UserTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)", m.table, userTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Username, data.HashPassword, data.Role, data.Name, data.Email, data.Phone, data.FacultyId, data.YearStart, data.Degree, data.AvataUrl, data.Birthday, data.BankAccount, data.Department)
	return ret, err
}

func (m *defaultUserTblModel) Update(ctx context.Context, data *UserTbl) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, userTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Id, data.Username, data.HashPassword, data.Role, data.Name, data.Email, data.Phone, data.FacultyId, data.YearStart, data.Degree, data.AvataUrl, data.Birthday, data.BankAccount, data.Department)
	return err
}

func (m *defaultUserTblModel) tableName() string {
	return m.table
}
