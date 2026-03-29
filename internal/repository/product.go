package repository

import (
	"go-shop-restful/internal/model"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (s productRepository) FindProducts(limit, offset int) ([]model.Product, error) {
	var products []model.Product
	err := s.db.Limit(limit).Offset(offset).Find(&products).Error

	return products, err
}

func (s productRepository) CreateProduct(newProduct *model.Product) error {
	return s.db.Create(&newProduct).Error
}

func (s productRepository) FindProductById(id int) (*model.Product, error) {
	var product model.Product
	err := s.db.First(&product, id).Error

	return &product, err
}

func (s productRepository) FindProductByTitle(title string) ([]model.Product, error) {
	var product []model.Product
	err := s.db.Where("title = ?", title).Find(&product).Error

	return product, err
}

func (s productRepository) UpdateProduct(id int, updateProduct *model.Product) error {
	var user model.User
	if err := s.db.First(&user, id).Error; err != nil {
		return err
	}

	return s.db.Model(&model.Product{Model: gorm.Model{ID: uint(id)}}).Updates(&updateProduct).Error
}

func (s productRepository) DeleteProduct(id int) error {
	return s.db.Delete(&model.Product{}, id).Error
}
