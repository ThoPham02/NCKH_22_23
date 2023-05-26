package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserTblModel = (*customUserTblModel)(nil)

type UserCondition struct {
	Name      string `json:"name"`
	Role      int64  `json:"role"`
	FacultyID int64  `json:"facultyID"`
}

type (
	// UserTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTblModel.
	UserTblModel interface {
		userTblModel
		FindOneByName(ctx context.Context, name string) (*UserTbl, error)
		FindUserByCondition(ctx context.Context, condition UserCondition) ([]UserTbl, error)
		ResetPassword(ctx context.Context, hashPassword string) error
	}

	customUserTblModel struct {
		*defaultUserTblModel
	}
)

// NewUserTblModel returns a model for the database table.
func NewUserTblModel(conn sqlx.SqlConn) UserTblModel {
	return &customUserTblModel{
		defaultUserTblModel: newUserTblModel(conn),
	}
}

func (m *customUserTblModel) FindOneByName(ctx context.Context, name string) (*UserTbl, error) {
	query := fmt.Sprintf("select %s from %s where username = $1 limit 1", userTblRows, m.table)
	var resp UserTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, name)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserTblModel) FindUserByCondition(ctx context.Context, condition UserCondition) ([]UserTbl, error) {
	query := fmt.Sprintf("select %s from %s", userTblRows, m.table)
	var resp []UserTbl
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserTblModel) ResetPassword(ctx context.Context, hashPassword string) error {
	query := fmt.Sprintf("update %s set hash_password = $1 ", m.table)
	_, err := m.conn.ExecCtx(ctx, query, hashPassword)
	return err
}
