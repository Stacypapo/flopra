package repository

import (
	"database/sql"
	"errors"
	"flowershy/internal/models"
	"time"
)

type UserPostgres struct {
	db *sql.DB
}

func NewUserPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) Create(user *models.User) (int64, error) {
	// user_id is BIGSERIAL -> let DB generate it
	query := `INSERT INTO users (email, password, name, phone_number, role, created_at)
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id`
	var id int64
	createdAt := user.CreatedAt
	if createdAt.IsZero() {
		createdAt = time.Now()
	}
	err := r.db.QueryRow(query, user.Email, user.Password, user.Name, user.PhoneNumber, user.Role, createdAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserPostgres) ReadById(id int64) (*models.User, error) {
	query := `SELECT user_id, email, password, name, phone_number, role, created_at FROM users WHERE user_id = $1`
	row := r.db.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.UserId, &user.Email, &user.Password, &user.Name, &user.PhoneNumber, &user.Role, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgres) ReadByEmail(email string) (*models.User, error) {
	query := `SELECT user_id, email, password, name, phone_number, role, created_at FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)

	var user models.User
	err := row.Scan(&user.UserId, &user.Email, &user.Password, &user.Name, &user.PhoneNumber, &user.Role, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserPostgres) ReadAll(limit, offset int) ([]*models.User, error) {
	query := `SELECT user_id, email, password, name, phone_number, role, created_at FROM users ORDER BY user_id LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.UserId, &u.Email, &u.Password, &u.Name, &u.PhoneNumber, &u.Role, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, rows.Err()
}

func (r *UserPostgres) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM users`
	var cnt int64
	err := r.db.QueryRow(query).Scan(&cnt)
	return cnt, err
}

func (r *UserPostgres) Update(user *models.User) (int64, error) {
	query := `UPDATE users SET email = $1, password = $2, name = $3, phone_number = $4, role = $5 WHERE user_id = $6`
	_, err := r.db.Exec(query, user.Email, user.Password, user.Name, user.PhoneNumber, user.Role, user.UserId)
	if err != nil {
		return 0, err
	}
	return user.UserId, nil
}

func (r *UserPostgres) Delete(id int64) (int64, error) {
	query := `DELETE FROM users WHERE user_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
