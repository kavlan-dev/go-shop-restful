package services

import "go-shop-restful/internal/models"

type ProductStorage interface {
	FindProducts(limit, offset int) (*[]models.Product, error)
	CreateProduct(product *models.Product) error
	FindProductById(id int) (*models.Product, error)
	FindProductByTitle(title string) (*models.Product, error)
	UpdateProduct(id int, updateProduct *models.Product) error
	DeleteProduct(product *models.Product) error
}

func (s *Services) Products(limit, offset int) (*[]models.Product, error) {
	return s.storage.FindProducts(limit, offset)
}

func (s *Services) CreateProduct(product *models.Product) error {
	return s.storage.CreateProduct(product)
}

func (s *Services) ProductById(id int) (*models.Product, error) {
	return s.storage.FindProductById(id)
}

func (s *Services) ProductByTitle(title string) (*models.Product, error) {
	return s.storage.FindProductByTitle(title)
}

func (s *Services) UpdateProduct(id int, updateProduct *models.Product) error {
	_, err := s.ProductById(id)
	if err != nil {
		return err
	}
	return s.storage.UpdateProduct(id, updateProduct)
}

func (s *Services) DeleteProduct(id int) error {
	product, err := s.ProductById(id)
	if err != nil {
		return err
	}

	return s.storage.DeleteProduct(product)
}
