package repository

import (
	"database/sql"
)

type ProductTagPostgres struct {
	db *sql.DB
}

// NewTransactionPostgres создает новый экземпляр TransactionPostgres.
func NewProductTagPostgres(db *sql.DB) *ProductTagPostgres {
	return &ProductTagPostgres{db: db}
}
