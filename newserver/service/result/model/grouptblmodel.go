package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GroupTblModel = (*customGroupTblModel)(nil)

type (
	// GroupTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupTblModel.
	GroupTblModel interface {
		groupTblModel
	}

	customGroupTblModel struct {
		*defaultGroupTblModel
	}
)

// NewGroupTblModel returns a model for the database table.
func NewGroupTblModel(conn sqlx.SqlConn) GroupTblModel {
	return &customGroupTblModel{
		defaultGroupTblModel: newGroupTblModel(conn),
	}
}
