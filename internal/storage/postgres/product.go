package postgres

import (
	"go-shop-restful/internal/model"

	"gorm.io/gorm"
)

func (s *storage) FindProducts(limit, offset int) (*[]model.Product, error) {
	var products []model.Product
	err := s.db.Limit(limit).Offset(offset).Find(&products).Error

	return &products, err
}

func (s *storage) CreateProduct(newProduct *model.Product) error {
	return s.db.Create(&newProduct).Error
}

func (s *storage) FindProductById(id int) (*model.Product, error) {
	var product model.Product
	err := s.db.First(&product, id).Error

	return &product, err
}

func (s *storage) FindProductByTitle(title string) (*[]model.Product, error) {
	var product []model.Product
	err := s.db.Where("title = ?", title).Find(&product).Error

	return &product, err
}

func (s *storage) UpdateProduct(id int, updateProduct *model.Product) error {
	var user model.User
	if err := s.db.First(&user, id).Error; err != nil {
		return err
	}

	return s.db.Model(&model.Product{Model: gorm.Model{ID: uint(id)}}).Updates(&updateProduct).Error
}

func (s *storage) DeleteProduct(id int) error {
	return s.db.Delete(&model.Product{}, id).Error
}
