package service

import (
	"flowershy/internal/models"
	"flowershy/internal/repository"
)

type CartService struct {
	repo repository.CartItem
}

func NewCartService(repo repository.CartItem) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) AddToCart(userId, productId int64, qty int) error {
	return s.repo.Create(&models.CartItem{
		UserId:    userId,
		ProductId: productId,
		Quantity:  qty,
	})
}

func (s *CartService) RemoveFromCart(userId, productId int64) error {
	return s.repo.Delete(userId, productId)
}

func (s *CartService) GetCartItems(userId int64) ([]models.CartItem, error) {
	return s.repo.ReadByUserId(userId)
}

func (s *CartService) ClearCart(userId int64) error {
	return s.repo.ClearCart(userId)
}

func (s *CartService) CountCartItems(userId int64) (int64, error) {
	return s.repo.CountByUserId(userId)
}
