package repository

import (
	"test-golang/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) CreateProduct(product *model.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetProductByID(id uint) (model.Product, error) {
	var product model.Product
	err := r.DB.First(&product, id).Error
	return product, err
}

func (r *ProductRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r *ProductRepository) UpdateProduct(product *model.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	return r.DB.Delete(&model.Product{}, id).Error
}
