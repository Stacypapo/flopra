package repository

import (
	"database/sql"
)

type OrderItemPostgres struct {
	db *sql.DB
}

// NewTransactionPostgres создает новый экземпляр TransactionPostgres.
func NewOrderItemPostgres(db *sql.DB) *OrderItemPostgres {
	return &OrderItemPostgres{db: db}
}
