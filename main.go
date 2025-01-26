package main

import (
    "database/sql"
    "net/http"
    "test-superindo/internal/handler"
    "test-superindo/internal/repository"
    "test-superindo/internal/service"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/superindo")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    productRepo := repository.NewProductRepository(db)
    productService := service.NewProductService(productRepo)
    productHandler := handler.NewProductHandler(productService)

    http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            productHandler.CreateProduct(w, r)
        } else if r.Method == http.MethodGet {
            productHandler.GetProducts(w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    http.ListenAndServe(":8080", nil)
}
