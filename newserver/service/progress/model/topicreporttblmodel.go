package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TopicReportTblModel = (*customTopicReportTblModel)(nil)

type TopicReportConditions struct {
	StageID int64 `json:"stageID"`
	EventID int64 `json:"eventID"`
	TopicID int64 `json:"topicID"`
	Limit   int64 `json:"limit"`
	Offset  int64 `json:"offset"`
}

type (
	// TopicReportTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTopicReportTblModel.
	TopicReportTblModel interface {
		topicReportTblModel
		FindTopicReports(ctx context.Context, condition TopicReportConditions) ([]TopicReportTbl, error)
		CountTopicReports(ctx context.Context, condition TopicReportConditions) (int64, error)
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

func (m *customTopicReportTblModel) FindTopicReports(ctx context.Context, condition TopicReportConditions) ([]TopicReportTbl, error) {
	var data []TopicReportTbl
	var values = []interface{}{}
	query := fmt.Sprintf("select %s from %s where ", topicReportTblRows, m.table)
	if condition.StageID != 0 {
		values = append(values, condition.StageID)
		query += fmt.Sprintf(" stage_id = $%d and ", len(values))
	}
	if condition.TopicID != 0 {
		values = append(values, condition.TopicID)
		query += fmt.Sprintf(" topic_id = $%d and ", len(values))
	}
	if condition.EventID != 0 {
		values = append(values, condition.EventID)
		query += fmt.Sprintf(" stage_id in (select stage_id from stage_tbl where event_id = $%d) and ", len(values))
	}
	query = query[0 : len(query)-5]
	if condition.Limit > 0 {
		values = append(values, condition.Limit, condition.Offset)
		query += fmt.Sprintf(" limit $%d offset $%d", len(values)-1, len(values))
	}

	err := m.conn.QueryRowsCtx(ctx, &data, query, values...)
	switch err {
	case nil:
		return data, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customTopicReportTblModel) CountTopicReports(ctx context.Context, condition TopicReportConditions) (int64, error) {
	var count int64
	var values = []interface{}{}
	query := fmt.Sprintf("select count(*) from %s where ", m.table)
	if condition.StageID != 0 {
		values = append(values, condition.StageID)
		query += fmt.Sprintf(" stage_id = $%d and ", len(values))
	}
	if condition.TopicID != 0 {
		values = append(values, condition.TopicID)
		query += fmt.Sprintf(" topic_id = $%d and ", len(values))
	}
	if condition.EventID != 0 {
		values = append(values, condition.EventID)
		query += fmt.Sprintf(" stage_id in (select stage_id from stage_tbl where event_id = $%d) and ", len(values))
	}
	query = query[0 : len(query)-5]

	err := m.conn.QueryRowCtx(ctx, &count, query, values...)
	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}
