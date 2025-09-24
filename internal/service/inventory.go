package service

import (
	"flowershy/internal/repository"
)

type InventoryService struct {
	inventory_repo repository.Inventory
}

func NewInventoryService(inventory_repo repository.Inventory) *InventoryService {
	return &InventoryService{
		inventory_repo: inventory_repo,
	}
}
