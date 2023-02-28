package store

import (
	"database/sql"
	db "github/ThoPham02/research_management/db/sqlc"
)

type Store struct {
	*db.Queries
	db *sql.DB
}

func NewStore(dbSQl *sql.DB) *Store {
	return &Store{
		Queries: db.New(dbSQl),
		db:      dbSQl,
	}
}
