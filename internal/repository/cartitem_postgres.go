package repository

import (
	"database/sql"
)

type CartItemPostgres struct {
	db *sql.DB
}

// NewTransactionPostgres создает новый экземпляр TransactionPostgres.
func NewCartItemPostgres(db *sql.DB) *CartItemPostgres {
	return &CartItemPostgres{db: db}
}
