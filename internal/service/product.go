package service

import (
	"context"
	"encoding/json"
	"superindo-api/internal/model"
	"superindo-api/internal/repository"

	"github.com/go-redis/redis/v8"
)

type ProductService interface {
	CreateProduct(product *model.Product) error
	GetProducts(filter map[string]interface{}, sort string) ([]model.Product, error)
}

type productService struct {
	repo  repository.ProductRepository
	cache *redis.Client
}

func NewProductService(repo repository.ProductRepository, cache *redis.Client) ProductService {
	return &productService{repo: repo, cache: cache}
}

func (s *productService) CreateProduct(product *model.Product) error {
	if err := s.repo.CreateProduct(product); err != nil {
		return err
	}

	// Clear cache after creating a product
	ctx := context.Background()
	s.cache.Del(ctx, "products")
	return nil
}

func (s *productService) GetProducts(filter map[string]interface{}, sort string) ([]model.Product, error) {
	ctx := context.Background()
	cacheKey := "products"

	if cachedData, err := s.cache.Get(ctx, cacheKey).Result(); err == nil {
		var products []model.Product
		if err := json.Unmarshal([]byte(cachedData), &products); err == nil {
			return products, nil
		}
	}

	products, err := s.repo.GetProducts(filter, sort)
	if err != nil {
		return nil, err
	}

	// Cache the results
	if data, err := json.Marshal(products); err == nil {
		s.cache.Set(ctx, cacheKey, data, 0)
	}

	return products, nil
}
