package repository

import (
	"database/sql"
	"flowershy/internal/models"
)

type BouquetItemPostgres struct {
	db *sql.DB
}

func NewBouqetItemPostgres(db *sql.DB) *BouquetItemPostgres {
	return &BouquetItemPostgres{db: db}
}

func (r *BouquetItemPostgres) Add(item *models.BouquetItem) error {
	query := `INSERT INTO bouquet_items (product_id_parent, product_id_child, quantity) VALUES ($1, $2, $3)
	          ON CONFLICT (product_id_parent, product_id_child) DO UPDATE SET quantity = EXCLUDED.quantity`
	_, err := r.db.Exec(query, item.ProductIdParent, item.ProductIdChild, item.Quantity)
	return err
}

func (r *BouquetItemPostgres) ReadByParentId(productId int64) ([]*models.BouquetItem, error) {
	query := `SELECT product_id_parent, product_id_child, quantity FROM bouquet_items WHERE product_id_parent = $1`
	rows, err := r.db.Query(query, productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.BouquetItem
	for rows.Next() {
		var it models.BouquetItem
		if err := rows.Scan(&it.ProductIdParent, &it.ProductIdChild, &it.Quantity); err != nil {
			return nil, err
		}
		res = append(res, &it)
	}
	return res, rows.Err()
}

func (r *BouquetItemPostgres) Delete(productIdParent, productIdChild int64) error {
	query := `DELETE FROM bouquet_items WHERE product_id_parent = $1 AND product_id_child = $2`
	_, err := r.db.Exec(query, productIdParent, productIdChild)
	return err
}

func (r *BouquetItemPostgres) DeleteByParentId(productIdParent int64) error {
	query := `DELETE FROM bouquet_items WHERE product_id_parent = $1`
	_, err := r.db.Exec(query, productIdParent)
	return err
}
