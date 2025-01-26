package main

import (
    "database/sql"
    "fmt"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/superindo")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    products := []struct {
        Name  string
        Type  string
        Price float64
    }{
        {"Wortel", "Sayuran", 10000},
        {"Ayam", "Protein", 50000},
        {"Apel", "Buah", 20000},
        {"Keripik Kentang", "Snack", 15000},
    }

    for _, product := range products {
        query := `INSERT INTO products (name, type, price, created_at) VALUES (?, ?, ?, ?)`
        _, err := db.Exec(query, product.Name, product.Type, product.Price, time.Now())
        if err != nil {
            fmt.Printf("Failed to seed product %s: %v\n", product.Name, err)
        }
    }

    fmt.Println("Seeder completed successfully")
}
