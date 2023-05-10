package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ LibraryTblModel = (*customLibraryTblModel)(nil)

type (
	// LibraryTblModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLibraryTblModel.
	LibraryTblModel interface {
		libraryTblModel
	}

	customLibraryTblModel struct {
		*defaultLibraryTblModel
	}
)

// NewLibraryTblModel returns a model for the database table.
func NewLibraryTblModel(conn sqlx.SqlConn) LibraryTblModel {
	return &customLibraryTblModel{
		defaultLibraryTblModel: newLibraryTblModel(conn),
	}
}
