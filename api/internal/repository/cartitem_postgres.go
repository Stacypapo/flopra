package repository

import (
	"database/sql"
	"flowershy/internal/models"
)

type CartItemPostgres struct {
	db *sql.DB
}

func NewCartItemPostgres(db *sql.DB) *CartItemPostgres {
	return &CartItemPostgres{db: db}
}

func (r *CartItemPostgres) Add(item *models.CartItem) error {
	// upsert: если уже есть — обновляем количество
	query := `INSERT INTO cart_items (user_id, product_id, quantity) VALUES ($1, $2, $3)
	          ON CONFLICT (user_id, product_id) DO UPDATE SET quantity = EXCLUDED.quantity`
	_, err := r.db.Exec(query, item.UserId, item.ProductId, item.Quantity)
	return err
}

func (r *CartItemPostgres) ReadByUserId(userId int64) ([]*models.CartItem, error) {
	query := `SELECT user_id, product_id, quantity FROM cart_items WHERE user_id = $1`
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.CartItem
	for rows.Next() {
		var it models.CartItem
		if err := rows.Scan(&it.UserId, &it.ProductId, &it.Quantity); err != nil {
			return nil, err
		}
		res = append(res, &it)
	}
	return res, rows.Err()
}

func (r *CartItemPostgres) UpdateQuantity(userId, productId int64, quantity int) error {
	query := `UPDATE cart_items SET quantity = $1 WHERE user_id = $2 AND product_id = $3`
	_, err := r.db.Exec(query, quantity, userId, productId)
	return err
}

func (r *CartItemPostgres) Delete(userId, productId int64) error {
	query := `DELETE FROM cart_items WHERE user_id = $1 AND product_id = $2`
	_, err := r.db.Exec(query, userId, productId)
	return err
}

func (r *CartItemPostgres) ClearCart(userId int64) error {
	query := `DELETE FROM cart_items WHERE user_id = $1`
	_, err := r.db.Exec(query, userId)
	return err
}

func (r *CartItemPostgres) CountByUserId(userId int64) (int64, error) {
	query := `SELECT COUNT(*) FROM cart_items WHERE user_id = $1`
	var cnt int64
	err := r.db.QueryRow(query, userId).Scan(&cnt)
	return cnt, err
}
