package services

import "go-shop-restful/internal/models"

type productStorage interface {
	FindProducts(limit, offset int) (*[]models.Product, error)
	CreateProduct(product *models.Product) error
	FindProductById(id int) (*models.Product, error)
	FindProductByTitle(title string) (*models.Product, error)
	UpdateProduct(id int, updateProduct *models.Product) error
	DeleteProduct(product *models.Product) error
}

func (s *service) Products(limit, offset int) (*[]models.Product, error) {
	return s.storage.FindProducts(limit, offset)
}

func (s *service) CreateProduct(product *models.Product) error {
	return s.storage.CreateProduct(product)
}

func (s *service) ProductById(id int) (*models.Product, error) {
	return s.storage.FindProductById(id)
}

func (s *service) ProductByTitle(title string) (*models.Product, error) {
	return s.storage.FindProductByTitle(title)
}

func (s *service) UpdateProduct(id int, updateProduct *models.Product) error {
	_, err := s.ProductById(id)
	if err != nil {
		return err
	}
	return s.storage.UpdateProduct(id, updateProduct)
}

func (s *service) DeleteProduct(id int) error {
	product, err := s.ProductById(id)
	if err != nil {
		return err
	}

	return s.storage.DeleteProduct(product)
}
