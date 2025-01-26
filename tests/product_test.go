package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"github.com/fabrianivan-id/test-superindo/controllers"
	"github.com/fabrianivan-id/test-superindo/models"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	productController := controllers.ProductController{}
	r.POST("/product", productController.AddProduct)
	r.GET("/product", productController.GetProducts)
	r.GET("/product/search", productController.SearchProduct)
	r.GET("/product/filter", productController.FilterProducts)
	r.GET("/product/sort", productController.SortProducts)
	return r
}

func TestAddProduct(t *testing.T) {
	router := setupRouter()

	product := models.Product{
		Name:  "Test Product",
		Type:  "Fruits",
		Price: 10.99,
	}

	jsonProduct, _ := json.Marshal(product)

	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonProduct))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetProducts(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/product", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSearchProduct(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/product/search?name=Test Product", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFilterProducts(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/product/filter?type=Fruits", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSortProducts(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/product/sort?by=price", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}