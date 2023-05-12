package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FacultyTblModel = (*customFacultyTblModel)(nil)

type (
	// FacultyTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFacultyTblModel.
	FacultyTblModel interface {
		facultyTblModel
		FindFaculties(ctx context.Context) ([]FacultyTbl, error)
	}

	customFacultyTblModel struct {
		*defaultFacultyTblModel
	}
)

// NewFacultyTblModel returns a model for the database table.
func NewFacultyTblModel(conn sqlx.SqlConn) FacultyTblModel {
	return &customFacultyTblModel{
		defaultFacultyTblModel: newFacultyTblModel(conn),
	}
}

func (m *customFacultyTblModel) FindFaculties(ctx context.Context) ([]FacultyTbl, error) {
	var data []FacultyTbl
	query := fmt.Sprintf("select %s from %s ", facultyTblRows, m.table)
	query += " ;"
	err := m.conn.QueryRowsCtx(ctx, &data, query)
	switch err {
	case nil:
		return data, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
