# Super Indo API

## Overview
Super Indo API is a RESTful API for managing products in a supermarket. It allows users to add, retrieve, search, filter, and sort products. The API is built using Go and utilizes a SQL database for data storage, Redis for caching, and follows best practices for dependency injection and testing.

## Features
- Add new products
- Retrieve a list of products
- Search products by name or ID
- Filter products by type (Vegetables, Protein, Fruits, Snacks)
- Sort products by date, price, or name

## Tech Stack
- **Language**: Go
- **Database**: SQL
- **Cache**: Redis
- **Dependency Injection**: Wire
- **Unit Testing**: Go testing framework
- **Containerization**: Docker

## Project Structure
```
github.com/fabrianivan-id/test-superindo
├── cmd
│   └── main.go
├── config
│   └── config.go
├── controllers
│   └── product_controller.go
├── models
│   └── product.go
├── repositories
│   └── product_repository.go
├── services
│   └── product_service.go
├── cache
│   └── redis.go
├── migrations
│   └── 20231010120000_create_products_table.sql
├── seeders
│   └── product_seeder.go
├── tests
│   └── product_test.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

## Setup Instructions
1. Clone the repository:
   ```
   git clone https://github.com/fabrianivan-id/test-superindo.git
   cd test-superindo
   ```

2. Build the Docker image:
   ```
   docker build -t test-superindo .
   ```

3. Run the application using Docker Compose:
   ```
   docker-compose up
   ```

4. Access the API at `http://localhost:8080/product`.

## API Usage
### Add Product
- **Endpoint**: `POST /product`
- **Request Body**: JSON object containing product details.

### Get Products
- **Endpoint**: `GET /product`
- **Query Parameters**: Optional filters for searching and sorting.

### Search Product
- **Endpoint**: `GET /product/search`
- **Query Parameters**: `id` or `name`.

### Filter Products
- **Endpoint**: `GET /product/filter`
- **Query Parameters**: `type` (e.g., Vegetables, Protein, Fruits, Snacks).

### Sort Products
- **Endpoint**: `GET /product/sort`
- **Query Parameters**: `sortBy` (e.g., date, price, name).

## Testing
Run unit tests using:
```
go test ./tests
```
