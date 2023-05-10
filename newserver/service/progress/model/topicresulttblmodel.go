package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TopicResultTblModel = (*customTopicResultTblModel)(nil)

type (
	// TopicResultTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTopicResultTblModel.
	TopicResultTblModel interface {
		topicResultTblModel
	}

	customTopicResultTblModel struct {
		*defaultTopicResultTblModel
	}
)

// NewTopicResultTblModel returns a model for the database table.
func NewTopicResultTblModel(conn sqlx.SqlConn) TopicResultTblModel {
	return &customTopicResultTblModel{
		defaultTopicResultTblModel: newTopicResultTblModel(conn),
	}
}
