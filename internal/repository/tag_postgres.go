package repository

import (
	"database/sql"
)

type TagPostgres struct {
	db *sql.DB
}

// NewTransactionPostgres создает новый экземпляр TransactionPostgres.
func NewTagPostgres(db *sql.DB) *TagPostgres {
	return &TagPostgres{db: db}
}
