package service

import (
	"flowershy/internal/models"
	"flowershy/internal/repository"
)

type ProductService struct {
	product_repo repository.Product
}

func NewProductService(product_repo repository.Product) *ProductService {
	return &ProductService{
		product_repo: product_repo,
	}
}

func (s *ProductService) CreateProduct(product *models.Product) (int64, error) {
	return s.product_repo.Create(product)
}
func (s *ProductService) GetProductById(productId int64) (*models.Product, error) {
	return s.product_repo.ReadById(productId)
}
func (s *ProductService) GetProductByName(name string) (*models.Product, error) {
	return s.product_repo.ReadByName(name)
}
func (s *ProductService) GetProductBySKU(sku string) (*models.Product, error) {
	return s.product_repo.ReadBySKU(sku)
}
func (s *ProductService) UpdateProduct(product *models.Product) (int64, error) {
	return s.product_repo.Update(product)
}
func (s *ProductService) DeleteProduct(productId int64) (int64, error) {
	return s.product_repo.Delete(productId)
}
