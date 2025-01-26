# Super Indo API

## Overview
The Super Indo API is a RESTful API designed to manage product data for Super Indo. It allows users to add, retrieve, search, filter, and sort product information. The API is built using Go and follows best practices for structuring a web application.

## Features
- Add and retrieve product data
- Search products by name and ID
- Filter products by type (Vegetables, Protein, Fruits, Snacks)
- Sort products by date, price, and product name
- Optional Redis caching for improved performance
- Dependency injection using Wire
- Unit tests for product functionalities
- Docker setup for easy deployment

## Project Structure
```
github.com/fabrianivan-id/test-superindo
├── cmd
│   └── main.go
├── config
│   └── config.go
├── controllers
│   └── product_controller.go
├── database
│   ├── migrations
│   │   └── 20231010_create_products_table.sql
│   ├── seeders
│   │   └── product_seeder.go
│   └── database.go
├── models
│   └── product.go
├── repositories
│   └── product_repository.go
├── routes
│   └── product_routes.go
├── services
│   └── product_service.go
├── cache
│   └── redis_cache.go
├── di
│   └── wire.go
├── tests
│   └── product_test.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

## Setup Instructions
1. **Clone the repository:**
   ```
   git clone https://github.com/yourusername/github.com/fabrianivan-id/test-superindo.git
   cd github.com/fabrianivan-id/test-superindo
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run database migrations:**
   Ensure your database is set up and configured in `config/config.go`, then run:
   ```
   go run database/migrations/20231010_create_products_table.sql
   ```

4. **Seed the database:**
   ```
   go run database/seeders/product_seeder.go
   ```

5. **Start the application:**
   ```
   go run cmd/main.go
   ```

6. **Access the API:**
   The API will be available at `http://localhost:8080/product`.

## API Usage
- **Add Product:** `POST /product`
- **Get Products:** `GET /product`
- **Search Product:** `GET /product/search?name={name}`
- **Filter Products:** `GET /product/filter?type={type}`
- **Sort Products:** `GET /product/sort?by={date|price|name}`

## Docker Setup
To run the application in a Docker container, use the following commands:
1. Build the Docker image:
   ```
   docker build -t github.com/fabrianivan-id/test-superindo .
   ```

2. Run the Docker container:
   ```
   docker run -p 8080:8080 github.com/fabrianivan-id/test-superindo
   ```

## Testing
To run unit tests, execute:
```
go test ./tests
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.