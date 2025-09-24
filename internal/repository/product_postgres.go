package repository

import (
	"database/sql"
	"flowershy/internal/models"
)

type ProductPostgres struct {
	db *sql.DB
}

// NewTransactionPostgres создает новый экземпляр TransactionPostgres.
func NewProductPostgres(db *sql.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(user *models.Product) (int64, error) {
	return 0, nil
}
func (r *ProductPostgres) ReadById(id int64) (*models.Product, error) {
	return nil, nil
}
func (r *ProductPostgres) ReadByName(name string) (*models.Product, error) {
	return nil, nil
}
func (r *ProductPostgres) ReadBySKU(sku string) (*models.Product, error) {
	return nil, nil
}
func (r *ProductPostgres) Update(user *models.Product) (int64, error) {
	return 0, nil
}
func (r *ProductPostgres) Delete(id int64) (int64, error) {
	return 0, nil
}
