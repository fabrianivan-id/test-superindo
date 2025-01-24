package main

import (
	"superindo-api/config"
	"superindo-api/internal/handler"
	"superindo-api/internal/repository"
	"superindo-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Setup database connection
	db := config.SetupDatabase()
	defer config.CloseDatabase(db)

	// Setup Redis connection
	redisClient := config.SetupRedis()
	defer config.CloseRedis(redisClient)

	// Initialize dependencies
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, redisClient)
	productHandler := handler.NewProductHandler(productService)

	// Create router
	r := gin.Default()

	// Product endpoints
	r.POST("/product", productHandler.CreateProduct)
	r.GET("/product", productHandler.GetProducts)

	// Start server
	r.Run(":8080")
}
