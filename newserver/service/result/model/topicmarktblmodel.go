package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TopicMarkTblModel = (*customTopicMarkTblModel)(nil)

type TopicMarkConditions struct {
	TopicIDs []int64 `json:"topic_ids"`
}

type (
	// TopicMarkTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTopicMarkTblModel.
	TopicMarkTblModel interface {
		topicMarkTblModel
		FindTopicMarks(ctx context.Context, conditions TopicMarkConditions) ([]TopicMarkTbl, error)
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

func (m *customTopicMarkTblModel) FindTopicMarks(ctx context.Context, conditions TopicMarkConditions) ([]TopicMarkTbl, error) {
	query := fmt.Sprintf("select %s from %s where", topicMarkTblRows, m.table)
	var resp []TopicMarkTbl
	var values = []interface{}{}
	if len(conditions.TopicIDs) > 0 {
		query += " topic_id in ("
		for _, id := range conditions.TopicIDs {
			query += "?, "
			values = append(values, id)
		}
		query = query[0:len(query)-1] + ") and "
	}
	query = query[0 : len(query)-5]
	err := m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
