package postgres

import "go-shop-restful/internal/models"

func (s *storage) FindProducts(limit, offset int) (*[]models.Product, error) {
	var products []models.Product
	err := s.db.Limit(limit).Offset(offset).Find(&products).Error
	return &products, err
}

func (s *storage) CreateProduct(product *models.Product) error {
	return s.db.Create(&product).Error
}

func (s *storage) FindProductById(id int) (*models.Product, error) {
	var product models.Product
	err := s.db.First(&product, id).Error
	return &product, err
}

func (s *storage) FindProductByTitle(title string) (*models.Product, error) {
	var product models.Product
	err := s.db.Where("title = ?", title).Find(&product).Error
	return &product, err
}

func (s *storage) UpdateProduct(id int, updateProduct *models.Product) error {
	product, err := s.FindProductById(id)
	if err != nil {
		return err
	}
	return s.db.Model(&product).Updates(&updateProduct).Error
}

func (s *storage) DeleteProduct(product *models.Product) error {
	return s.db.Delete(&product).Error
}
