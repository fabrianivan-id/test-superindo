package repository

import (
	"superindo-api/internal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *model.Product) error
	GetProducts(filter map[string]interface{}, sort string) ([]model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetProducts(filter map[string]interface{}, sort string) ([]model.Product, error) {
	var products []model.Product
	query := r.db

	if name, ok := filter["name"]; ok {
		query = query.Where("name LIKE ?", "%"+name.(string)+"%")
	}

	if id, ok := filter["id"]; ok {
		query = query.Where("id = ?", id)
	}

	if productType, ok := filter["type"]; ok {
		query = query.Where("type = ?", productType)
	}

	if sort != "" {
		query = query.Order(sort)
	}

	err := query.Find(&products).Error
	return products, err
}
