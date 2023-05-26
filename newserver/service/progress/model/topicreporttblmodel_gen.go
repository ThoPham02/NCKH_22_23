// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	topicReportTblFieldNames          = builder.RawFieldNames(&TopicReportTbl{}, true)
	topicReportTblRows                = strings.Join(topicReportTblFieldNames, ",")
	topicReportTblRowsExpectAutoSet   = strings.Join(stringx.Remove(topicReportTblFieldNames), ",")
	topicReportTblRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(topicReportTblFieldNames, "id"))
)

type (
	topicReportTblModel interface {
		Insert(ctx context.Context, data *TopicReportTbl) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TopicReportTbl, error)
		Update(ctx context.Context, data *TopicReportTbl) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTopicReportTblModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TopicReportTbl struct {
		Id          int64          `db:"id"`
		TopicId     int64          `db:"topic_id"`
		StageId     int64          `db:"stage_id"`
		Description sql.NullString `db:"description"`
		ReportUrl   string         `db:"report_url"`
		CreatedAt   int64          `db:"created_at"`
		CreatedBy   int64          `db:"created_by"`
	}
)

func newTopicReportTblModel(conn sqlx.SqlConn) *defaultTopicReportTblModel {
	return &defaultTopicReportTblModel{
		conn:  conn,
		table: `"public"."topic_report_tbl"`,
	}
}

func (m *defaultTopicReportTblModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where id = $1", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTopicReportTblModel) FindOne(ctx context.Context, id int64) (*TopicReportTbl, error) {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", topicReportTblRows, m.table)
	var resp TopicReportTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTopicReportTblModel) Insert(ctx context.Context, data *TopicReportTbl) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7)", m.table, topicReportTblRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.TopicId, data.StageId, data.Description, data.ReportUrl, data.CreatedAt, data.CreatedBy)
	return ret, err
}

func (m *defaultTopicReportTblModel) Update(ctx context.Context, data *TopicReportTbl) error {
	query := fmt.Sprintf("update %s set %s where id = $1", m.table, topicReportTblRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Id, data.TopicId, data.StageId, data.Description, data.ReportUrl, data.CreatedAt, data.CreatedBy)
	return err
}

func (m *defaultTopicReportTblModel) tableName() string {
	return m.table
}
