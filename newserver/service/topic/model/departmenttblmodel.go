package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ DepartmentTblModel = (*customDepartmentTblModel)(nil)

type (
	// DepartmentTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDepartmentTblModel.
	DepartmentTblModel interface {
		departmentTblModel
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
