package handler

import (
    "encoding/json"
    "net/http"
    "test-superindo/internal/model"
    "test-superindo/internal/service"
    "time"
)

type ProductHandler struct {
    service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
    return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product model.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    product.CreatedAt = time.Now()
    if err := h.service.CreateProduct(product); err != nil {
        http.Error(w, "Failed to create product", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Product created successfully"})
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    nameOrID := query.Get("query")
    productType := query.Get("type")
    sortField := query.Get("sort")
    sortOrder := query.Get("order")

    var products []model.Product
    var err error

    switch {
    case nameOrID != "":
        products, err = h.service.SearchProduct(nameOrID)
    case productType != "":
        products, err = h.service.FilterProductsByType(productType)
    case sortField != "" && sortOrder != "":
        products, err = h.service.SortProducts(sortField, sortOrder)
    default:
        products, err = h.service.GetProducts()
    }

    if err != nil {
        http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(products)
}
