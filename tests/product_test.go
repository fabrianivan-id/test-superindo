package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabrianivan-id/test-superindo/controllers"
	"github.com/fabrianivan-id/test-superindo/models"
	"github.com/fabrianivan-id/test-superindo/services"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	productController := controllers.NewProductController(services.NewProductService())
	router.POST("/product", productController.AddProduct)
	router.GET("/product", productController.ListProducts)
	router.GET("/product/search", productController.SearchProduct)
	return router
}

func TestAddProduct(t *testing.T) {
	router := setupRouter()

	product := models.Product{
		Name:  "Test Product",
		Type:  "Sayuran",
		Price: 10000,
	}

	jsonProduct, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(jsonProduct))
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestListProducts(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/product", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestSearchProduct(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/product/search?id=1", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
