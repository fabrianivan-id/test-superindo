package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/fabrianivan-id/test-superindo/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) AddProduct(product *models.Product) error {
	query := "INSERT INTO products (name, type, price, created_at) VALUES (?, ?, ?, ?)"
	_, err := r.db.Exec(query, product.Name, product.Type, product.Price, time.Now())
	return err
}

func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
	query := "SELECT id, name, type, price, created_at FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Price, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) FindProductByID(id int) (*models.Product, error) {
	query := "SELECT id, name, type, price, created_at FROM products WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var product models.Product
	if err := row.Scan(&product.ID, &product.Name, &product.Type, &product.Price, &product.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) SearchProducts(name string) ([]models.Product, error) {
	query := "SELECT id, name, type, price, created_at FROM products WHERE name LIKE ?"
	rows, err := r.db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Price, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) FilterProductsByType(productType string) ([]models.Product, error) {
	query := "SELECT id, name, type, price, created_at FROM products WHERE type = ?"
	rows, err := r.db.Query(query, productType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Price, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) SortProductsByDate() ([]models.Product, error) {
	query := "SELECT id, name, type, price, created_at FROM products ORDER BY created_at"
	return r.getSortedProducts(query)
}

func (r *ProductRepository) SortProductsByPrice() ([]models.Product, error) {
	query := "SELECT id, name, type, price, created_at FROM products ORDER BY price"
	return r.getSortedProducts(query)
}

func (r *ProductRepository) SortProductsByName() ([]models.Product, error) {
	query := "SELECT id, name, type, price, created_at FROM products ORDER BY name"
	return r.getSortedProducts(query)
}

func (r *ProductRepository) getSortedProducts(query string) ([]models.Product, error) {
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Price, &product.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
