package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SubcommitteeTblModel = (*customSubcommitteeTblModel)(nil)

type SubcommitteeConditions struct {
}

type (
	// SubcommitteeTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSubcommitteeTblModel.
	SubcommitteeTblModel interface {
		subcommitteeTblModel
		FindSubcommittee(ctx context.Context, condition SubcommitteeConditions) ([]SubcommitteeTbl, error)
	}

	customSubcommitteeTblModel struct {
		*defaultSubcommitteeTblModel
	}
)

// NewSubcommitteeTblModel returns a model for the database table.
func NewSubcommitteeTblModel(conn sqlx.SqlConn) SubcommitteeTblModel {
	return &customSubcommitteeTblModel{
		defaultSubcommitteeTblModel: newSubcommitteeTblModel(conn),
	}
}

func (m *customSubcommitteeTblModel) FindSubcommittee(ctx context.Context, condition SubcommitteeConditions) ([]SubcommitteeTbl, error) {
	query := fmt.Sprintf("select %s from %s where", subcommitteeTblRows, m.table)
	var resp []SubcommitteeTbl
	query = query[0 : len(query)-5]
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
