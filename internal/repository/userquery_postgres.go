package repository

import (
	"database/sql"
)

type UserQueryPostgres struct {
	db *sql.DB
}

// NewTransactionPostgres создает новый экземпляр TransactionPostgres.
func NewUserQueryPostgres(db *sql.DB) *UserQueryPostgres {
	return &UserQueryPostgres{db: db}
}
