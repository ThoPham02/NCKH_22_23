package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DepartmentTblModel = (*customDepartmentTblModel)(nil)

type (
	// DepartmentTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDepartmentTblModel.
	DepartmentTblModel interface {
		departmentTblModel
		FindDepartments(ctx context.Context, facultyID int64) ([]DepartmentTbl, error)
	}

	customDepartmentTblModel struct {
		*defaultDepartmentTblModel
	}
)

// NewDepartmentTblModel returns a model for the database table.
func NewDepartmentTblModel(conn sqlx.SqlConn) DepartmentTblModel {
	return &customDepartmentTblModel{
		defaultDepartmentTblModel: newDepartmentTblModel(conn),
	}
}

func (m *customDepartmentTblModel) FindDepartments(ctx context.Context, facultyID int64) ([]DepartmentTbl, error) {
	var data []DepartmentTbl
	var values = []interface{}{}
	query := fmt.Sprintf("select %s from %s ", departmentTblRows, m.table)
	if facultyID != 0 {
		query = query + "where faculty_id = $1"
		values = append(values, facultyID)
	}
	query += " ;"
	err := m.conn.QueryRowsCtx(ctx, &data, query, values...)
	switch err {
	case nil:
		return data, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
