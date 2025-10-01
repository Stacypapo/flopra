package repository

import (
	"database/sql"
	"flowershy/internal/models"
	"time"
)

type UserQueryPostgres struct {
	db *sql.DB
}

func NewUserQueryPostgres(db *sql.DB) *UserQueryPostgres {
	return &UserQueryPostgres{db: db}
}

func (r *UserQueryPostgres) Create(q *models.UserQuery) error {
	query := `INSERT INTO user_queries (product_id, user_id, query, model, rating, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	createdAt := q.CreatedAt
	if createdAt.IsZero() {
		createdAt = time.Now()
	}
	_, err := r.db.Exec(query, q.ProductId, q.UserId, q.Query, q.Model, q.Rating, createdAt)
	return err
}

func (r *UserQueryPostgres) ReadByUserId(userId int64, limit, offset int) ([]*models.UserQuery, error) {
	query := `SELECT product_id, user_id, query, model, rating, created_at FROM user_queries WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(query, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.UserQuery
	for rows.Next() {
		var q models.UserQuery
		if err := rows.Scan(&q.ProductId, &q.UserId, &q.Query, &q.Model, &q.Rating, &q.CreatedAt); err != nil {
			return nil, err
		}
		res = append(res, &q)
	}
	return res, rows.Err()
}

func (r *UserQueryPostgres) ReadByProductId(productId int64, limit, offset int) ([]*models.UserQuery, error) {
	query := `SELECT product_id, user_id, query, model, rating, created_at FROM user_queries WHERE product_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(query, productId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.UserQuery
	for rows.Next() {
		var q models.UserQuery
		if err := rows.Scan(&q.ProductId, &q.UserId, &q.Query, &q.Model, &q.Rating, &q.CreatedAt); err != nil {
			return nil, err
		}
		res = append(res, &q)
	}
	return res, rows.Err()
}

func (r *UserQueryPostgres) ReadAll(limit, offset int) ([]*models.UserQuery, error) {
	query := `SELECT product_id, user_id, query, model, rating, created_at FROM user_queries ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.UserQuery
	for rows.Next() {
		var q models.UserQuery
		if err := rows.Scan(&q.ProductId, &q.UserId, &q.Query, &q.Model, &q.Rating, &q.CreatedAt); err != nil {
			return nil, err
		}
		res = append(res, &q)
	}
	return res, rows.Err()
}

func (r *UserQueryPostgres) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM user_queries`
	var cnt int64
	err := r.db.QueryRow(query).Scan(&cnt)
	return cnt, err
}

func (r *UserQueryPostgres) CountByUserId(userId int64) (int64, error) {
	query := `SELECT COUNT(*) FROM user_queries WHERE user_id = $1`
	var cnt int64
	err := r.db.QueryRow(query, userId).Scan(&cnt)
	return cnt, err
}
