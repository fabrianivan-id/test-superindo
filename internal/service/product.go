package service

import (
    "test-superindo/internal/model"
    "test-superindo/internal/repository"
)

type ProductService struct {
    repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(product model.Product) error {
    return s.repo.Create(product)
}

func (s *ProductService) GetProducts() ([]model.Product, error) {
    return s.repo.FindAll()
}

func (s *ProductService) SearchProduct(query string) ([]model.Product, error) {
    return s.repo.SearchByNameOrID(query)
}

func (s *ProductService) FilterProductsByType(productType string) ([]model.Product, error) {
    return s.repo.FilterByType(productType)
}

func (s *ProductService) SortProducts(field, order string) ([]model.Product, error) {
    return s.repo.SortBy(field, order)
}
