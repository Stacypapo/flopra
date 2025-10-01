package repository

import (
	"database/sql"
	"errors"
	"flowershy/internal/models"
	"time"
)

type OrderPostgres struct {
	db *sql.DB
}

func NewOrderPostgres(db *sql.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) Create(order *models.Order) (int64, error) {
	query := `INSERT INTO orders (user_id, status, total_amount, created_at, shipping_address, warehouse_id)
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING order_id`
	createdAt := order.CreatedAt
	if createdAt.IsZero() {
		createdAt = time.Now()
	}
	var id int64
	err := r.db.QueryRow(query, order.UserId, order.Status, order.TotalAmount, createdAt, order.ShippingAddress, order.WarehouseId).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *OrderPostgres) ReadById(id int64) (*models.Order, error) {
	query := `SELECT order_id, user_id, status, total_amount, created_at, shipping_address, warehouse_id FROM orders WHERE order_id = $1`
	row := r.db.QueryRow(query, id)
	var o models.Order
	err := row.Scan(&o.OrderId, &o.UserId, &o.Status, &o.TotalAmount, &o.CreatedAt, &o.ShippingAddress, &o.WarehouseId)
	if err == sql.ErrNoRows {
		return nil, errors.New("order not found")
	}
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *OrderPostgres) ReadByUserId(userId int64, limit, offset int) ([]*models.Order, error) {
	query := `SELECT order_id, user_id, status, total_amount, created_at, shipping_address, warehouse_id FROM orders WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(query, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.OrderId, &o.UserId, &o.Status, &o.TotalAmount, &o.CreatedAt, &o.ShippingAddress, &o.WarehouseId); err != nil {
			return nil, err
		}
		res = append(res, &o)
	}
	return res, rows.Err()
}

func (r *OrderPostgres) ReadByWarehouseId(warehouseId int64, limit, offset int) ([]*models.Order, error) {
	query := `SELECT order_id, user_id, status, total_amount, created_at, shipping_address, warehouse_id FROM orders WHERE warehouse_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(query, warehouseId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.Order
	for rows.Next() {
		var o models.Order
		if err := rows.Scan(&o.OrderId, &o.UserId, &o.Status, &o.TotalAmount, &o.CreatedAt, &o.ShippingAddress, &o.WarehouseId); err != nil {
			return nil, err
		}
		res = append(res, &o)
	}
	return res, rows.Err()
}

func (r *OrderPostgres) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM orders`
	var cnt int64
	err := r.db.QueryRow(query).Scan(&cnt)
	return cnt, err
}

func (r *OrderPostgres) CountByUserId(userId int64) (int64, error) {
	query := `SELECT COUNT(*) FROM orders WHERE user_id = $1`
	var cnt int64
	err := r.db.QueryRow(query, userId).Scan(&cnt)
	return cnt, err
}

func (r *OrderPostgres) Update(order *models.Order) (int64, error) {
	query := `UPDATE orders SET status = $1, total_amount = $2, shipping_address = $3, warehouse_id = $4 WHERE order_id = $5`
	_, err := r.db.Exec(query, order.Status, order.TotalAmount, order.ShippingAddress, order.WarehouseId, order.OrderId)
	if err != nil {
		return 0, err
	}
	return order.OrderId, nil
}

func (r *OrderPostgres) Delete(id int64) (int64, error) {
	query := `DELETE FROM orders WHERE order_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
