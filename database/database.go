package database

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func InitDB(dataSourceName string) {
    var err error
    db, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
}

func GetDB() *sql.DB {
    return db
}