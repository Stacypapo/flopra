package service

import (
	"errors"
	"flowershy/internal/models"
	"flowershy/internal/repository"
)

type OrderService struct {
	orderRepo repository.Order
	itemRepo  repository.OrderItem
	invRepo   repository.Inventory
}

func NewOrderService(order repository.Order, item repository.OrderItem, inv repository.Inventory) *OrderService {
	return &OrderService{orderRepo: order, itemRepo: item, invRepo: inv}
}

func (s *OrderService) CreateOrder(userId int64, items map[int64]int) (int64, error) {
	if len(items) == 0 {
		return 0, errors.New("empty order")
	}

	order := &models.Order{UserId: userId, Status: "pending"}
	orderId, err := s.orderRepo.Create(order)
	if err != nil {
		return 0, err
	}

	for productId, qty := range items {
		if err := s.invRepo.ReserveStock(orderId, map[int64]int{productId: qty}); err != nil {
			return 0, err
		}
		if err := s.itemRepo.Create(&models.OrderItem{
			OrderId:   orderId,
			ProductId: productId,
			Quantity:  qty}); err != nil {
			return 0, err
		}
	}

	return orderId, nil
}

func (s *OrderService) GetOrderById(id int64) (*models.Order, error) {
	return s.orderRepo.ReadById(id)
}

func (s *OrderService) GetOrdersByUserId(userId int64, limit int, offset int) ([]models.Order, error) {
	return s.orderRepo.ReadByUserId(userId, limit, offset)
}

func (s *OrderService) GetOrdersByWarehouseId(warehouseId int64, limit int, offset int) ([]models.Order, error) {
	return s.orderRepo.ReadByWarehouseId(warehouseId, limit, offset)
}

func (s *OrderService) GetAllOrders(limit, offset int) ([]models.Order, error) {
	return s.orderRepo.ReadAll(limit, offset)
}

func (s *OrderService) CountOrders() (int64, error) {
	return s.orderRepo.Count()
}

func (s *OrderService) UpdateOrderStatus(id int64, status string) (int64, error) {
	order, err := s.orderRepo.ReadById(id)
	if err != nil {
		return 0, err
	}
	order.Status = status
	return s.orderRepo.Update(order)
}

func (s *OrderService) DeleteOrder(id int64) (int64, error) {
	return s.orderRepo.Delete(id)
}
