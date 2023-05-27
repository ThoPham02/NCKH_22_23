package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ StageTblModel = (*customStageTblModel)(nil)

type (
	// StageTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStageTblModel.
	StageTblModel interface {
		stageTblModel
		FindStages(ctx context.Context, eventID int64) ([]StageTbl, error)
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

func (m *customStageTblModel) FindStages(ctx context.Context, eventID int64) ([]StageTbl, error) {
	var values = []interface{}{}

	query := fmt.Sprintf("select %s from %s where", stageTblRows, m.table)
	if eventID != 0 {
		values = append(values, eventID)
		query += fmt.Sprintf(" event_id = $%d and ", len(values))
	}

	query = query[0 : len(query)-5]
	query += " order by id asc"
	var resp []StageTbl
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
