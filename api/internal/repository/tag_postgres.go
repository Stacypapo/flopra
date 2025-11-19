package repository

import (
	"database/sql"
	"errors"
	"flowershy/internal/models"
)

type TagPostgres struct {
	db *sql.DB
}

func NewTagPostgres(db *sql.DB) *TagPostgres {
	return &TagPostgres{db: db}
}

func (r *TagPostgres) Create(t *models.Tag) (int64, error) {
	query := `INSERT INTO tags (name, type, color_hex) VALUES ($1, $2, $3) RETURNING tag_id`
	var id int64
	err := r.db.QueryRow(query, t.Name, t.Type, t.ColorHex).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TagPostgres) ReadById(id int64) (*models.Tag, error) {
	query := `SELECT tag_id, name, type, color_hex FROM tags WHERE tag_id = $1`
	row := r.db.QueryRow(query, id)
	var t models.Tag
	if err := row.Scan(&t.TagId, &t.Name, &t.Type, &t.ColorHex); err == sql.ErrNoRows {
		return nil, errors.New("tag not found")
	} else if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TagPostgres) ReadAll(limit, offset int) ([]models.Tag, error) {
	query := `SELECT tag_id, name, type, color_hex FROM tags ORDER BY tag_id LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, limit, offset)
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

func (r *TagPostgres) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM tags`
	var cnt int64
	err := r.db.QueryRow(query).Scan(&cnt)
	return cnt, err
}

func (r *TagPostgres) Update(t *models.Tag) (int64, error) {
	query := `UPDATE tags SET name = $1, type = $2, color_hex = $3 WHERE tag_id = $4`
	_, err := r.db.Exec(query, t.Name, t.Type, t.ColorHex, t.TagId)
	if err != nil {
		return 0, err
	}
	return t.TagId, nil
}

func (r *TagPostgres) Delete(id int64) (int64, error) {
	query := `DELETE FROM tags WHERE tag_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
