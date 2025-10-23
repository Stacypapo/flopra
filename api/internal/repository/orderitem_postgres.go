package repository

import (
	"database/sql"
	"flowershy/internal/models"
)

type OrderItemPostgres struct {
	db *sql.DB
}

func NewOrderItemPostgres(db *sql.DB) *OrderItemPostgres {
	return &OrderItemPostgres{db: db}
}

func (r *OrderItemPostgres) Create(item *models.OrderItem) error {
	query := `INSERT INTO order_items (order_id, product_id, quantity) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, item.OrderId, item.ProductId, item.Quantity)
	return err
}

func (r *OrderItemPostgres) ReadByOrderId(orderId int64) ([]*models.OrderItem, error) {
	query := `SELECT order_id, product_id, quantity FROM order_items WHERE order_id = $1`
	rows, err := r.db.Query(query, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.OrderItem
	for rows.Next() {
		var it models.OrderItem
		if err := rows.Scan(&it.OrderId, &it.ProductId, &it.Quantity); err != nil {
			return nil, err
		}
		res = append(res, &it)
	}
	return res, rows.Err()
}

func (r *OrderItemPostgres) ReadByProductId(productId int64) ([]*models.OrderItem, error) {
	query := `SELECT order_id, product_id, quantity FROM order_items WHERE product_id = $1`
	rows, err := r.db.Query(query, productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.OrderItem
	for rows.Next() {
		var it models.OrderItem
		if err := rows.Scan(&it.OrderId, &it.ProductId, &it.Quantity); err != nil {
			return nil, err
		}
		res = append(res, &it)
	}
	return res, rows.Err()
}

func (r *OrderItemPostgres) Delete(orderId, productId int64) error {
	query := `DELETE FROM order_items WHERE order_id = $1 AND product_id = $2`
	_, err := r.db.Exec(query, orderId, productId)
	return err
}

func (r *OrderItemPostgres) DeleteByOrderId(orderId int64) error {
	query := `DELETE FROM order_items WHERE order_id = $1`
	_, err := r.db.Exec(query, orderId)
	return err
}
