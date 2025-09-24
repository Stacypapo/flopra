package service

import (
	"flowershy/internal/repository"
)

type CartItemService struct {
	cart_item_repo repository.CartItem
}

func NewCartItemService(cart_item_repo repository.CartItem) *CartItemService {
	return &CartItemService{
		cart_item_repo: cart_item_repo,
	}
}
