package repository

import (
	"database/sql"
	"errors"
	"flowershy/internal/models"
)

type WarehousePostgres struct {
	db *sql.DB
}

func NewWarehousePostgres(db *sql.DB) *WarehousePostgres {
	return &WarehousePostgres{db: db}
}

func (r *WarehousePostgres) Create(w *models.Warehouse) (int64, error) {
	query := `INSERT INTO warehouses (name, address) VALUES ($1, $2) RETURNING warehouse_id`
	var id int64
	err := r.db.QueryRow(query, w.Name, w.Address).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *WarehousePostgres) ReadById(id int64) (*models.Warehouse, error) {
	query := `SELECT warehouse_id, name, address FROM warehouses WHERE warehouse_id = $1`
	row := r.db.QueryRow(query, id)
	var w models.Warehouse
	if err := row.Scan(&w.WarehouseId, &w.Name, &w.Address); err == sql.ErrNoRows {
		return nil, errors.New("warehouse not found")
	} else if err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *WarehousePostgres) ReadAll(limit, offset int) ([]*models.Warehouse, error) {
	query := `SELECT warehouse_id, name, address FROM warehouses ORDER BY warehouse_id LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.Warehouse
	for rows.Next() {
		var w models.Warehouse
		if err := rows.Scan(&w.WarehouseId, &w.Name, &w.Address); err != nil {
			return nil, err
		}
		res = append(res, &w)
	}
	return res, rows.Err()
}

func (r *WarehousePostgres) Count() (int64, error) {
	query := `SELECT COUNT(*) FROM warehouses`
	var cnt int64
	err := r.db.QueryRow(query).Scan(&cnt)
	return cnt, err
}

func (r *WarehousePostgres) Update(w *models.Warehouse) (int64, error) {
	query := `UPDATE warehouses SET name = $1, address = $2 WHERE warehouse_id = $3`
	_, err := r.db.Exec(query, w.Name, w.Address, w.WarehouseId)
	if err != nil {
		return 0, err
	}
	return w.WarehouseId, nil
}

func (r *WarehousePostgres) Delete(id int64) (int64, error) {
	query := `DELETE FROM warehouses WHERE warehouse_id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
