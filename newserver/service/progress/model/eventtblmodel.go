package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ EventTblModel = (*customEventTblModel)(nil)

type (
	// EventTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customEventTblModel.
	EventTblModel interface {
		eventTblModel
		FindCurrentEvent(ctx context.Context) (*EventTbl, error)
		UpdateCurrentEvent(ctx context.Context, eventID int64) error
		FindEvents(ctx context.Context, isCurrent int64) ([]EventTbl, error)
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

func (m *customEventTblModel) FindCurrentEvent(ctx context.Context) (*EventTbl, error) {
	query := fmt.Sprintf("select %s from %s where is_current = 1 limit 1", eventTblRows, m.table)
	var resp EventTbl
	err := m.conn.QueryRowCtx(ctx, &resp, query)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
func (m *customEventTblModel) UpdateCurrentEvent(ctx context.Context, eventID int64) error {
	query := fmt.Sprintf("update %s set %s", m.table, "is_current=2")
	_, err := m.conn.ExecCtx(ctx, query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("update %s set %s where id = $1", m.table, "is_current=1")
	_, err = m.conn.ExecCtx(ctx, query, eventID)
	return err
}

func (m *customEventTblModel) FindEvents(ctx context.Context, isCurrent int64) ([]EventTbl, error) {
	query := fmt.Sprintf("select %s from %s", eventTblRows, m.table)
	var resp []EventTbl
	var values = []interface{}{}
	if isCurrent > 0 {
		query += "where is_current = $1"
		values = append(values, isCurrent)
	}
	err := m.conn.QueryRowsCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
