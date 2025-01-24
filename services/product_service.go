package services

import (
	"errors"
	"time"

	"github.com/fabrianivan-id/test-superindo/models"
	"github.com/fabrianivan-id/test-superindo/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) AddProduct(product *models.Product) error {
	if product.Name == "" || product.Type == "" {
		return errors.New("product name and type are required")
	}
	product.CreatedAt = time.Now()
	return s.repo.CreateProduct(product)
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *ProductService) SearchProduct(id int, name string) (*models.Product, error) {
	if id > 0 {
		return s.repo.FindProductByID(id)
	}
	if name != "" {
		return s.repo.FindProductByName(name)
	}
	return nil, errors.New("either id or name must be provided for search")
}

func (s *ProductService) FilterProducts(productType string) ([]models.Product, error) {
	return s.repo.FilterProducts(productType)
}

func (s *ProductService) SortProducts(sortBy string) ([]models.Product, error) {
	return s.repo.SortProducts(sortBy)
}