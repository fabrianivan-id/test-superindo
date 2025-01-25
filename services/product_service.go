package services

import (
	"errors"
	"time"

	"github.com/yourusername/super-indo-api/models"
	"github.com/yourusername/super-indo-api/repositories"
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
	return s.repo.Add(product)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) SearchProductByID(id int) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) SearchProductByName(name string) ([]models.Product, error) {
	return s.repo.FindByName(name)
}

func (s *ProductService) FilterProductsByType(productType string) ([]models.Product, error) {
	return s.repo.FilterByType(productType)
}

func (s *ProductService) SortProductsByDate() ([]models.Product, error) {
	return s.repo.SortByDate()
}

func (s *ProductService) SortProductsByPrice() ([]models.Product, error) {
	return s.repo.SortByPrice()
}

func (s *ProductService) SortProductsByName() ([]models.Product, error) {
	return s.repo.SortByName()
}