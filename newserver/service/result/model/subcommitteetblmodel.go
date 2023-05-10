package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ SubcommitteeTblModel = (*customSubcommitteeTblModel)(nil)

type (
	// SubcommitteeTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSubcommitteeTblModel.
	SubcommitteeTblModel interface {
		subcommitteeTblModel
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
