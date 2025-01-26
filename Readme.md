# Super Indo Product API

## Overview
This API allows users to manage Super Indo products with functionalities such as:
1. Adding new products.
2. Fetching all products.
3. Searching products by name or ID.
4. Filtering products by type.
5. Sorting products by price, date, or name.

## Endpoints

### `POST /product`
**Description**: Add a new product.  
**Body**:
```json
{
    "name": "Tomat",
    "type": "Sayuran",
    "price": 12000
}
```

### `GET /product`
**Description**: Fetch all products or apply filters and sorting.  
**Query Parameters**:
- `query`: Search by name or ID.
- `type`: Filter by product type (Sayuran, Protein, Buah, Snack).
- `sort`: Field to sort by (name, price, created_at).
- `order`: Sort order (asc, desc).

**Example**: `/product?type=Buah&sort=price&order=asc`

## Running Locally
1. Clone the repository.
2. Run `docker-compose up`.

## Seeder
Run the seeder using:
```bash
go run seeder/product_seeder.go
```

## Unit Testing
Run tests using:
```bash
go test ./...
```

## Tech Stack
- **Language**: Golang
- **Database**: MySQL
- **Cache**: Redis
- **Dependency Injection**: Wire
- **Testing**: Testify
- **Containerization**: Docker

Jika ada bagian yang perlu diperjelas atau diperluas, beri tahu saya!
