package repository

import (
	"database/sql"
	"errors"
	"flowershy/internal/models"
	"time"
)

type ProductPostgres struct {
	db *sql.DB
}

func NewProductPostgres(db *sql.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(product *models.Product) (int64, error) {
	query := `INSERT INTO products (sku, name, description, price, created_at, url)
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING product_id`
	createdAt := product.CreatedAt
	if createdAt.IsZero() {
		createdAt = time.Now()
	}
	var id int64
	err := r.db.QueryRow(query, product.SKU, product.Name, product.Description, product.Price, createdAt, product.URL).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProductPostgres) ReadById(id int64) (*models.Product, error) {
	query := `SELECT product_id, sku, name, description, price, created_at, url FROM products WHERE product_id = $1`
	row := r.db.QueryRow(query, id)
	var p models.Product
	err := row.Scan(&p.ProductId, &p.SKU, &p.Name, &p.Description, &p.Price, &p.CreatedAt, &p.URL)
	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductPostgres) ReadByName(name string, limit, offset int) ([]*models.Product, error) {
	query := `SELECT product_id, sku, name, description, price, created_at, url FROM products WHERE name ILIKE '%' || $1 || '%' ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(query, name, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ProductId, &p.SKU, &p.Name, &p.Description, &p.Price, &p.CreatedAt, &p.URL); err != nil {
			return nil, err
		}
		res = append(res, &p)
	}
	return res, rows.Err()
}

func (r *ProductPostgres) ReadBySKU(sku string) (*models.Product, error) {
	query := `SELECT product_id, sku, name, description, price, created_at, url FROM products WHERE sku = $1`
	row := r.db.QueryRow(query, sku)
	var p models.Product
	err := row.Scan(&p.ProductId, &p.SKU, &p.Name, &p.Description, &p.Price, &p.CreatedAt, &p.URL)
	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductPostgres) ReadAll(limit, offset int) ([]*models.Product, error) {
	query := `SELECT product_id, sku, name, description, price, created_at, url FROM products ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ProductId, &p.SKU, &p.Name, &p.Description, &p.Price, &p.CreatedAt, &p.URL); err != nil {
			return nil, err
		}
		res = append(res, &p)
	}
	return res, rows.Err()
}

func (r *ProductPostgres) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM products`
	var cnt int64
	err := r.db.QueryRow(query).Scan(&cnt)
	return cnt, err
}

func (r *ProductPostgres) Update(product *models.Product) (int64, error) {
	query := `UPDATE products SET sku = $1, name = $2, description = $3, price = $4, url = $5 WHERE product_id = $6`
	_, err := r.db.Exec(query, product.SKU, product.Name, product.Description, product.Price, product.URL, product.ProductId)
	if err != nil {
		return 0, err
	}
	return product.ProductId, nil
}

func (r *ProductPostgres) Delete(id int64) (int64, error) {
	query := `DELETE FROM products WHERE product_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
