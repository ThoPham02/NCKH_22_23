package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ StudentGroupTblModel = (*customStudentGroupTblModel)(nil)

type (
	// StudentGroupTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStudentGroupTblModel.
	StudentGroupTblModel interface {
		studentGroupTblModel
	}

	customStudentGroupTblModel struct {
		*defaultStudentGroupTblModel
	}
)

// NewStudentGroupTblModel returns a model for the database table.
func NewStudentGroupTblModel(conn sqlx.SqlConn) StudentGroupTblModel {
	return &customStudentGroupTblModel{
		defaultStudentGroupTblModel: newStudentGroupTblModel(conn),
	}
}
