package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TopicReportTblModel = (*customTopicReportTblModel)(nil)

type (
	// TopicReportTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTopicReportTblModel.
	TopicReportTblModel interface {
		topicReportTblModel
	}

	customTopicReportTblModel struct {
		*defaultTopicReportTblModel
	}
)

// NewTopicReportTblModel returns a model for the database table.
func NewTopicReportTblModel(conn sqlx.SqlConn) TopicReportTblModel {
	return &customTopicReportTblModel{
		defaultTopicReportTblModel: newTopicReportTblModel(conn),
	}
}
