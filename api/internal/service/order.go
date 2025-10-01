package service

import (
	"flowershy/internal/repository"
)

type OrderService struct {
	order_repo repository.Order
}

func NewOrderService(order_repo repository.Order) *OrderService {
	return &OrderService{
		order_repo: order_repo,
	}
}
