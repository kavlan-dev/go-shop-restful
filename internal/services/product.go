package services

import (
	"golang-shop-restful/internal/models"
)

func (s *Services) GetProducts(limit, offset int) ([]models.Product, error) {
	var products []models.Product
	if err := s.db.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Services) CreateProduct(product *models.Product) error {
	if err := s.db.Create(product).Error; err != nil {
		return err
	}

	return nil
}

func (s *Services) GetProductById(id int) (models.Product, error) {
	var product models.Product

	if err := s.db.First(&product, id).Error; err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (s *Services) UpdateProduct(id int, updateProduct *models.Product) error {
	existingProduct, err := s.GetProductById(id)
	if err != nil {
		return err
	}

	if err := s.db.Model(&existingProduct).Updates(updateProduct).Error; err != nil {
		return err
	}

	return nil
}

func (s *Services) DeleteProduct(id int) error {
	product, err := s.GetProductById(id)
	if err != nil {
		return err
	}

	if err := s.db.Delete(&product).Error; err != nil {
		return err
	}

	return nil
}
