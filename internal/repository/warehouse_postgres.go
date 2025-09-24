package repository

import (
	"database/sql"
)

type WarehousePostgres struct {
	db *sql.DB
}

// NewTransactionPostgres создает новый экземпляр TransactionPostgres.
func NewWarehousePostgres(db *sql.DB) *WarehousePostgres {
	return &WarehousePostgres{db: db}
}
