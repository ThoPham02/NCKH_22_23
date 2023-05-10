package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ StageTblModel = (*customStageTblModel)(nil)

type (
	// StageTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStageTblModel.
	StageTblModel interface {
		stageTblModel
	}

	customStageTblModel struct {
		*defaultStageTblModel
	}
)

// NewStageTblModel returns a model for the database table.
func NewStageTblModel(conn sqlx.SqlConn) StageTblModel {
	return &customStageTblModel{
		defaultStageTblModel: newStageTblModel(conn),
	}
}
