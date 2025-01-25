package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"super-indo-api/models"
	"super-indo-api/services"
)

type ProductController struct {
	ProductService services.ProductService
}

func (pc *ProductController) AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := pc.ProductService.AddProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) ListProducts(c *gin.Context) {
	products, err := pc.ProductService.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) SearchProduct(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	product, err := pc.ProductService.SearchProduct(id, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) FilterProducts(c *gin.Context) {
	productType := c.Query("type")
	products, err := pc.ProductService.FilterProducts(productType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) SortProducts(c *gin.Context) {
	sortBy := c.Query("sort_by")
	products, err := pc.ProductService.SortProducts(sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}