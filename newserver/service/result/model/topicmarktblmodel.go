package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TopicMarkTblModel = (*customTopicMarkTblModel)(nil)

type (
	// TopicMarkTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTopicMarkTblModel.
	TopicMarkTblModel interface {
		topicMarkTblModel
	}

	customTopicMarkTblModel struct {
		*defaultTopicMarkTblModel
	}
)

// NewTopicMarkTblModel returns a model for the database table.
func NewTopicMarkTblModel(conn sqlx.SqlConn) TopicMarkTblModel {
	return &customTopicMarkTblModel{
		defaultTopicMarkTblModel: newTopicMarkTblModel(conn),
	}
}
