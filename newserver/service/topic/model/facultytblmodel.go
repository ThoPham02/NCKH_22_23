package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FacultyTblModel = (*customFacultyTblModel)(nil)

type (
	// FacultyTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFacultyTblModel.
	FacultyTblModel interface {
		facultyTblModel
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
