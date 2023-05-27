package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ StageDetailTblModel = (*customStageDetailTblModel)(nil)

type (
	// StageDetailTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStageDetailTblModel.
	StageDetailTblModel interface {
		stageDetailTblModel
	}

	customStageDetailTblModel struct {
		*defaultStageDetailTblModel
	}
)

// NewStageDetailTblModel returns a model for the database table.
func NewStageDetailTblModel(conn sqlx.SqlConn) StageDetailTblModel {
	return &customStageDetailTblModel{
		defaultStageDetailTblModel: newStageDetailTblModel(conn),
	}
}
