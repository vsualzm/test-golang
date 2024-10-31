package service

import (
	"test-golang/model"
	"test-golang/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product *model.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *ProductService) GetProductByID(id uint) (model.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) UpdateProduct(product *model.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}
