package repository

import (
    "database/sql"
    "test-superindo/internal/model"
    "fmt"
)

type ProductRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
    return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(product model.Product) error {
    query := `INSERT INTO products (name, type, price, created_at) VALUES (?, ?, ?, ?)`
    _, err := r.db.Exec(query, product.Name, product.Type, product.Price, product.CreatedAt)
    return err
}

func (r *ProductRepository) FindAll() ([]model.Product, error) {
    rows, err := r.db.Query("SELECT id, name, type, price, created_at FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []model.Product
    for rows.Next() {
        var product model.Product
        err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Price, &product.CreatedAt)
        if err != nil {
            return nil, err
        }
        products = append(products, product)
    }
    return products, nil
}

func (r *ProductRepository) SearchByNameOrID(query string) ([]model.Product, error) {
    sqlQuery := `SELECT id, name, type, price, created_at FROM products WHERE name LIKE ? OR id = ?`
    rows, err := r.db.Query(sqlQuery, fmt.Sprintf("%%%s%%", query), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []model.Product
    for rows.Next() {
        var product model.Product
        err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Price, &product.CreatedAt)
        if err != nil {
            return nil, err
        }
        products = append(products, product)
    }
    return products, nil
}

func (r *ProductRepository) FilterByType(productType string) ([]model.Product, error) {
    rows, err := r.db.Query("SELECT id, name, type, price, created_at FROM products WHERE type = ?", productType)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []model.Product
    for rows.Next() {
        var product model.Product
        err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Price, &product.CreatedAt)
        if err != nil {
            return nil, err
        }
        products = append(products, product)
    }
    return products, nil
}

func (r *ProductRepository) SortBy(field, order string) ([]model.Product, error) {
    query := fmt.Sprintf("SELECT id, name, type, price, created_at FROM products ORDER BY %s %s", field, order)
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []model.Product
    for rows.Next() {
        var product model.Product
        err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Price, &product.CreatedAt)
        if err != nil {
            return nil, err
        }
        products = append(products, product)
    }
    return products, nil
}
