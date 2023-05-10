package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ EventTblModel = (*customEventTblModel)(nil)

type (
	// EventTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEventTblModel.
	EventTblModel interface {
		eventTblModel
	}

	customEventTblModel struct {
		*defaultEventTblModel
	}
)

// NewEventTblModel returns a model for the database table.
func NewEventTblModel(conn sqlx.SqlConn) EventTblModel {
	return &customEventTblModel{
		defaultEventTblModel: newEventTblModel(conn),
	}
}
