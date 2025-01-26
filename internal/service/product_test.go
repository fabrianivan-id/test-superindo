package service

import (
    "errors"
    "test-superindo/internal/model"
    "test-superindo/internal/repository"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
    mock.Mock
}

func (m *MockProductRepository) Create(product model.Product) error {
    args := m.Called(product)
    return args.Error(0)
}

func (m *MockProductRepository) FindAll() ([]model.Product, error) {
    args := m.Called()
    return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductRepository) SearchByNameOrID(query string) ([]model.Product, error) {
    args := m.Called(query)
    return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductRepository) FilterByType(productType string) ([]model.Product, error) {
    args := m.Called(productType)
    return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductRepository) SortBy(field, order string) ([]model.Product, error) {
    args := m.Called(field, order)
    return args.Get(0).([]model.Product), args.Error(1)
}

func TestCreateProduct(t *testing.T) {
    repo := new(MockProductRepository)
    service := NewProductService(repo)

    product := model.Product{
        Name:      "Tomat",
        Type:      "Sayuran",
        Price:     12000,
        CreatedAt: time.Now(),
    }

    repo.On("Create", product).Return(nil)

    err := service.CreateProduct(product)

    assert.NoError(t, err)
    repo.AssertExpectations(t)
}

func TestGetProducts(t *testing.T) {
    repo := new(MockProductRepository)
    service := NewProductService(repo)

    products := []model.Product{
        {ID: 1, Name: "Tomat", Type: "Sayuran", Price: 12000, CreatedAt: time.Now()},
        {ID: 2, Name: "Apel", Type: "Buah", Price: 15000, CreatedAt: time.Now()},
    }

    repo.On("FindAll").Return(products, nil)

    result, err := service.GetProducts()

    assert.NoError(t, err)
    assert.Equal(t, products, result)
    repo.AssertExpectations(t)
}
