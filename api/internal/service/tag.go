package service

import (
	"flowershy/internal/models"
	"flowershy/internal/repository"
)

type TagService struct {
	tagRepo        repository.Tag
	productTagRepo repository.ProductTag
}

func NewTagService(tagrepo repository.Tag, prodrepo repository.ProductTag) *TagService {
	return &TagService{tagRepo: tagrepo, productTagRepo: prodrepo}
}

func (s *TagService) AddTag(name, prodtype, color string) (int64, error) {
	return s.tagRepo.Create(&models.Tag{
		Name:     name,
		Type:     prodtype,
		ColorHex: color,
	})
}

func (s *TagService) GetTagById(tagId int64) (*models.Tag, error) {
	return s.tagRepo.ReadById(tagId)
}

func (s *TagService) GetAllTags(limit, offset int) ([]models.Tag, error) {
	return s.tagRepo.ReadAll(limit, offset)
}

func (s *TagService) ApplyTagToProduct(tagId, productId int64) error {
	return s.productTagRepo.Add(&models.ProductTag{
		TagId:     tagId,
		ProductId: productId,
	})
}
func (s *TagService) RemoveTagFromProduct(tagId, productId int64) error {
	return s.productTagRepo.Delete(tagId, productId)
}
func (s *TagService) GetTagsByProductId(productId int64) ([]models.Tag, error) {
	return s.productTagRepo.ReadByProductId(productId)
}
