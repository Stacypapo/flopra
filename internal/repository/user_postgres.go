package repository

import (
	"database/sql"
	"flowershy/internal/models"
)

type UserPostgres struct {
	db *sql.DB
}

// NewTransactionPostgres создает новый экземпляр TransactionPostgres.
func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user *models.User) (int64, error) {
	return 0, nil
}
func (r *UserPostgres) ReadById(id int64) (*models.User, error) {
	return nil, nil
}
func (r *UserPostgres) Update(user *models.User) (int64, error) {
	return 0, nil
}
func (r *UserPostgres) Delete(id int64) (int64, error) {
	return 0, nil
}
