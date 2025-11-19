package repository

import (
	"database/sql"
	"flowershy/internal/models"
)

type ProductTagPostgres struct {
	db *sql.DB
}

func NewProductTagPostgres(db *sql.DB) *ProductTagPostgres {
	return &ProductTagPostgres{db: db}
}

func (r *ProductTagPostgres) Add(pt *models.ProductTag) error {
	query := `INSERT INTO product_tags (tag_id, product_id) VALUES ($1, $2) ON CONFLICT (tag_id, product_id) DO NOTHING`
	_, err := r.db.Exec(query, pt.TagId, pt.ProductId)
	return err
}

func (r *ProductTagPostgres) ReadByProductId(productId int64) ([]models.Tag, error) {
	query := `SELECT t.tag_id, t.name, t.type, t.color_hex FROM tags t JOIN product_tags pt ON t.tag_id = pt.tag_id WHERE pt.product_id = $1`
	rows, err := r.db.Query(query, productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []models.Tag
	for rows.Next() {
		var t models.Tag
		if err := rows.Scan(&t.TagId, &t.Name, &t.Type, &t.ColorHex); err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return res, rows.Err()
}

func (r *ProductTagPostgres) ReadByTagId(tagId int64) ([]models.Product, error) {
	query := `SELECT p.product_id, p.sku, p.name, p.description, p.price, p.created_at, p.url FROM products p JOIN product_tags pt ON p.product_id = pt.product_id WHERE pt.tag_id = $1`
	rows, err := r.db.Query(query, tagId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ProductId, &p.SKU, &p.Name, &p.Description, &p.Price, &p.CreatedAt, &p.URL); err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	return res, rows.Err()
}

func (r *ProductTagPostgres) Delete(tagId, productId int64) error {
	query := `DELETE FROM product_tags WHERE tag_id = $1 AND product_id = $2`
	_, err := r.db.Exec(query, tagId, productId)
	return err
}

func (r *ProductTagPostgres) DeleteByProductId(productId int64) error {
	query := `DELETE FROM product_tags WHERE product_id = $1`
	_, err := r.db.Exec(query, productId)
	return err
}
