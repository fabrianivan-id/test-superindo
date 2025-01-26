package seeders

import (
	"database/sql"
	"log"
)

type Product struct {
	ID       int
	Name     string
	Type     string
	Price    float64
	CreatedAt string
}

func SeedProducts(db *sql.DB) {
	products := []Product{
		{ID: 1, Name: "Carrot", Type: "Vegetables", Price: 1.50, CreatedAt: "2023-10-10"},
		{ID: 2, Name: "Chicken Breast", Type: "Protein", Price: 5.00, CreatedAt: "2023-10-10"},
		{ID: 3, Name: "Apple", Type: "Fruits", Price: 0.80, CreatedAt: "2023-10-10"},
		{ID: 4, Name: "Potato Chips", Type: "Snacks", Price: 2.00, CreatedAt: "2023-10-10"},
	}

	for _, product := range products {
		_, err := db.Exec("INSERT INTO products (id, name, type, price, created_at) VALUES (?, ?, ?, ?, ?)",
			product.ID, product.Name, product.Type, product.Price, product.CreatedAt)
		if err != nil {
			log.Fatalf("Failed to seed product %s: %v", product.Name, err)
		}
	}
}