package service

import (
	"flowershy/internal/models"
	"flowershy/internal/repository"
)

type ProductService struct {
	productRepo repository.Product
}

func NewProductService(p repository.Product) *ProductService {
	return &ProductService{productRepo: p}
}

func (s *ProductService) CreateProduct(p *models.Product) (int64, error) {
	return s.productRepo.Create(p)
}

func (s *ProductService) GetProductById(id int64) (*models.Product, error) {
	return s.productRepo.ReadById(id)
}

func (s *ProductService) GetProductsByName(name string, limit int, offset int) ([]models.Product, error) {
	return s.productRepo.ReadByName(name, limit, offset)
}

func (s *ProductService) GetProductBySKU(sku string) (*models.Product, error) {
	return s.productRepo.ReadBySKU(sku)
}

func (s *ProductService) SearchProducts(query string, tags []int64) ([]models.Product, error) {
	return s.productRepo.Search(query, tags)
}

func (s *ProductService) GetAllProducts(limit, offset int) ([]models.Product, error) {
	return s.productRepo.ReadAll(limit, offset)
}

func (s *ProductService) CountProducts() (int64, error) {
	return s.productRepo.Count()
}

func (s *ProductService) UpdateProduct(p *models.Product) (int64, error) {
	return s.productRepo.Update(p)
}

func (s *ProductService) DeleteProduct(id int64) (int64, error) {
	return s.productRepo.Delete(id)
}
