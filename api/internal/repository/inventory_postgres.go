package repository

import (
	"database/sql"
	"flowershy/internal/models"
)

type InventoryPostgres struct {
	db *sql.DB
}

func NewInventoryPostgres(db *sql.DB) *InventoryPostgres {
	return &InventoryPostgres{db: db}
}

func (r *InventoryPostgres) Add(item *models.Inventory) error {
	query := `INSERT INTO inventory (warehouse_id, product_id, quantity) VALUES ($1, $2, $3)
	          ON CONFLICT (warehouse_id, product_id) DO UPDATE SET quantity = EXCLUDED.quantity`
	_, err := r.db.Exec(query, item.WarehouseId, item.ProductId, item.Quantity)
	return err
}

func (r *InventoryPostgres) ReadByProductId(productId int64) ([]*models.Inventory, error) {
	query := `SELECT warehouse_id, product_id, quantity FROM inventory WHERE product_id = $1`
	rows, err := r.db.Query(query, productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.Inventory
	for rows.Next() {
		var it models.Inventory
		if err := rows.Scan(&it.WarehouseId, &it.ProductId, &it.Quantity); err != nil {
			return nil, err
		}
		res = append(res, &it)
	}
	return res, rows.Err()
}

func (r *InventoryPostgres) ReadByWarehouseId(warehouseId int64) ([]*models.Inventory, error) {
	query := `SELECT warehouse_id, product_id, quantity FROM inventory WHERE warehouse_id = $1`
	rows, err := r.db.Query(query, warehouseId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*models.Inventory
	for rows.Next() {
		var it models.Inventory
		if err := rows.Scan(&it.WarehouseId, &it.ProductId, &it.Quantity); err != nil {
			return nil, err
		}
		res = append(res, &it)
	}
	return res, rows.Err()
}

func (r *InventoryPostgres) UpdateQuantity(inventory *models.Inventory) error {
	query := `UPDATE inventory SET quantity = $1 WHERE warehouse_id = $2 AND product_id = $3`
	_, err := r.db.Exec(query, inventory.Quantity, inventory.WarehouseId, inventory.ProductId)
	return err
}

func (r *InventoryPostgres) Delete(warehouseId, productId int64) error {
	query := `DELETE FROM inventory WHERE warehouse_id = $1 AND product_id = $2`
	_, err := r.db.Exec(query, warehouseId, productId)
	return err
}
