package repository

import (
	"database/sql"
)

type InventoryPostgres struct {
	db *sql.DB
}

// NewTransactionPostgres создает новый экземпляр TransactionPostgres.
func NewInventoryPostgres(db *sql.DB) *InventoryPostgres {
	return &InventoryPostgres{db: db}
}
