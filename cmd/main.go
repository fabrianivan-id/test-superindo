package main

import (
	"github.com/fabrianivan-id/test-superindo/config"
	"github.com/fabrianivan-id/test-superindo/controllers"
	"github.com/fabrianivan-id/test-superindo/repositories"
	"github.com/fabrianivan-id/test-superindo/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database and cache
	db := config.ConnectDatabase(cfg)
	redisClient := config.ConnectRedis(cfg)

	// Initialize repositories
	productRepo := repositories.NewProductRepository(db)

	// Initialize services
	productService := services.NewProductService(productRepo, redisClient)

	// Initialize controllers
	productController := controllers.NewProductController(productService)

	// Set up Gin router
	router := gin.Default()

	// Define routes
	router.POST("/product", productController.AddProduct)
	router.GET("/product", productController.ListProducts)
	router.GET("/product/search", productController.SearchProduct)
	router.GET("/product/filter", productController.FilterProducts)
	router.GET("/product/sort", productController.SortProducts)

	// Start the server
	router.Run(cfg.ServerAddress)
}
