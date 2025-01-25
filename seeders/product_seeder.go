package seeders

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type Product struct {
	ID        int
	Name      string
	Type      string
	Price     float64
	CreatedAt time.Time
}

func SeedProducts(db *sql.DB) {
	products := []Product{
		{ID: 1, Name: "Carrot", Type: "Sayuran", Price: 5000, CreatedAt: time.Now()},
		{ID: 2, Name: "Chicken Breast", Type: "Protein", Price: 30000, CreatedAt: time.Now()},
		{ID: 3, Name: "Apple", Type: "Buah", Price: 10000, CreatedAt: time.Now()},
		{ID: 4, Name: "Potato Chips", Type: "Snack", Price: 15000, CreatedAt: time.Now()},
	}

	for _, product := range products {
		_, err := db.Exec("INSERT INTO products (id, name, type, price, created_at) VALUES ($1, $2, $3, $4, $5)",
			product.ID, product.Name, product.Type, product.Price, product.CreatedAt)
		if err != nil {
			log.Fatalf("Failed to seed product %s: %v", product.Name, err)
		}
	}
}