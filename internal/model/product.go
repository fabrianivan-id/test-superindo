package models

import "time"

type Product struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Type        string    `json:"type" db:"type"`
	Price       float64   `json:"price" db:"price"`
	CreatedAt   time.time `json:"created_at" db:"created_at"`
}