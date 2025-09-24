package repository

import (
	"database/sql"
)

type BouqetItemPostgres struct {
	db *sql.DB
}

// NewTransactionPostgres создает новый экземпляр TransactionPostgres.
func NewBouqetItemPostgres(db *sql.DB) *BouqetItemPostgres {
	return &BouqetItemPostgres{db: db}
}
