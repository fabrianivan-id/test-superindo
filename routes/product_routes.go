package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/fabrianivan-id/test-superindo/controllers"
)

func SetupProductRoutes(router *gin.Engine, productController *controllers.ProductController) {
    productRoutes := router.Group("/product")
    {
        productRoutes.POST("/", productController.AddProduct)
        productRoutes.GET("/", productController.GetProducts)
        productRoutes.GET("/search", productController.SearchProduct)
        productRoutes.GET("/filter", productController.FilterProducts)
        productRoutes.GET("/sort", productController.SortProducts)
    }
}