package service

import "go-shop-restful/internal/model"

type productStorage interface {
	FindProducts(limit, offset int) ([]model.Product, error)
	CreateProduct(newProduct *model.Product) error
	FindProductById(id int) (*model.Product, error)
	FindProductByTitle(title string) ([]model.Product, error)
	UpdateProduct(id int, updateProduct *model.Product) error
	DeleteProduct(id int) error
}

type productService struct {
	storage productStorage
}

func NewProductService(storage productStorage) *productService {
	return &productService{storage: storage}
}

func (s productService) Products(limit, offset int) ([]model.Product, error) {
	return s.storage.FindProducts(limit, offset)
}

func (s productService) CreateProduct(newProduct *model.Product) error {
	if err := newProduct.Validate(); err != nil {
		return err
	}

	return s.storage.CreateProduct(newProduct)
}

func (s productService) ProductById(id int) (*model.Product, error) {
	return s.storage.FindProductById(id)
}

func (s productService) ProductByTitle(title string) ([]model.Product, error) {
	return s.storage.FindProductByTitle(title)
}

func (s productService) UpdateProduct(id int, updateProduct *model.Product) error {
	product, err := s.storage.FindProductById(id)
	if err != nil {
		return nil
	}

	updateProduct.Title = product.Title

	if err := updateProduct.Validate(); err != nil {
		return err
	}

	return s.storage.UpdateProduct(id, updateProduct)
}

func (s productService) DeleteProduct(id int) error {
	if _, err := s.storage.FindProductById(id); err != nil {
		return err
	}

	return s.storage.DeleteProduct(id)
}
