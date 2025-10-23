package service

import (
	"flowershy/internal/models"
	"flowershy/internal/repository"
)

type InventoryService struct {
	repo repository.Inventory
}

func NewInventoryService(repo repository.Inventory) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) UpdateInventory(warehouseId, productId int64, quantity int) error {
	return s.repo.UpdateQuantity(&models.Inventory{
		WarehouseId: warehouseId,
		ProductId:   productId,
		Quantity:    quantity,
	})
}

func (s *InventoryService) GetInventory(productId int64) ([]models.Inventory, error) {
	return s.repo.ReadByProductId(productId)
}

func (s *InventoryService) GetByWarehouse(warehouseId int64) ([]models.Inventory, error) {
	return s.repo.ReadByWarehouseId(warehouseId)
}

func (s *InventoryService) ReserveStock(orderId int64, items map[int64]int) error {
	return s.repo.ReserveStock(orderId, items)
}
