package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TopicTblModel = (*customTopicTblModel)(nil)

type (
	// TopicTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTopicTblModel.
	TopicTblModel interface {
		topicTblModel
	}

	customTopicTblModel struct {
		*defaultTopicTblModel
	}
)

// NewTopicTblModel returns a model for the database table.
func NewTopicTblModel(conn sqlx.SqlConn) TopicTblModel {
	return &customTopicTblModel{
		defaultTopicTblModel: newTopicTblModel(conn),
	}
}
