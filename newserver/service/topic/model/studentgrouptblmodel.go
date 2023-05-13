package model

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StudentGroupTblModel = (*customStudentGroupTblModel)(nil)

type (
	// StudentGroupTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStudentGroupTblModel.
	StudentGroupTblModel interface {
		studentGroupTblModel
		InsertMutil(ctx context.Context, data []StudentGroupTbl) (sql.Result, error)
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

func (m *customStudentGroupTblModel) InsertMutil(ctx context.Context, data []StudentGroupTbl) (sql.Result, error) {
	return nil, nil
}
