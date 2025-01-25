package config

import (
	"context"
	"log"
	"superindo-api/internal/model"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "user:password@tcp(127.0.0.1:3306)/superindo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&model.Product{})
	return db
}

func CloseDatabase(db *gorm.DB) {
	dbInstance, err := db.DB()
	if err != nil {
		log.Fatal("Failed to close database connection:", err)
	}
	dbInstance.Close()
}

func SetupRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "", // no password set
		DB: 0,          // use default DB
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	return client
}

func CloseRedis(client *redis.Client) {
	client.Close()
}
