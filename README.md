# hot-coffee
Learning Objectives
REST API development

JSON data handling

Logging with Go's log/slog

Basic software design principles (SOLID)

Layered software architecture

Abstract
hot-coffee replicates the backend systems used by real-world coffee shops. Staff can:

Manage customer orders (create, update, close, delete)

Track and update inventory

Manage menu items and ingredients

The project focuses on building a maintainable, scalable codebase using RESTful principles, JSON storage, and a three-layered architecture (Handlers, Services, Repositories).

Technologies
Language: Go (Golang)

Storage: Local JSON files (no databases)

Logging: Go's log/slog

Architecture: Layered (Presentation, Business Logic, Data Access)
```go
hot-coffee/
├── cmd/
│   └── main.go
├── internal/
│   ├── handler/
│   │   ├── order_handler.go
│   │   ├── menu_handler.go
│   │   └── inventory_handler.go
│   ├── service/
│   │   ├── order_service.go
│   │   ├── menu_service.go
│   │   └── inventory_service.go
│   └── dal/
│       ├── order_repository.go
│       ├── menu_repository.go
│       └── inventory_repository.go
├── models/
│   ├── order.go
│   ├── menu_item.go
│   └── inventory_item.go
├── go.mod
├── go.sum
└── data/
    ├── orders.json
    ├── menu_items.json
    └── inventory.json
```
API Endpoints
Orders
```go
POST /orders — Create a new order

GET /orders — Retrieve all orders

GET /orders/{id} — Retrieve an order by ID

PUT /orders/{id} — Update an order

DELETE /orders/{id} — Delete an order

POST /orders/{id}/close — Close an order
```
Menu Items
```go
POST /menu — Add a menu item

GET /menu — Retrieve all menu items

GET /menu/{id} — Retrieve a menu item by ID

PUT /menu/{id} — Update a menu item

DELETE /menu/{id} — Delete a menu item
```
Inventory
```go
POST /inventory — Add an inventory item

GET /inventory — Retrieve all inventory items

GET /inventory/{id} — Retrieve an inventory item by ID

PUT /inventory/{id} — Update an inventory item

DELETE /inventory/{id} — Delete an inventory item
```
Aggregations
```go
GET /reports/total-sales — Get total sales amount

GET /reports/popular-items — Get popular menu items
```
Data Storage
All application data is stored in structured JSON files located in the data/ directory:
```go
orders.json

menu_items.json

inventory.json
```
Data models include Order, MenuItem, and InventoryItem, each properly serialized and deserialized to JSON.

Requirements
Code must conform to gofumpt style.

No external packages allowed (only Go standard library).

Application must compile with:

```go
go build -o hot-coffee .
```
Proper error handling and HTTP status codes are mandatory.

Server configuration options:

```bash
./hot-coffee --port <PORT> --dir <DATA_DIRECTORY>
./hot-coffee --help
```
Error Handling
400 Bad Request for invalid inputs

404 Not Found for missing resources

500 Internal Server Error for server-side failures

All errors and significant events are logged using Go's log/slog package.

```go
slog.Info("Order created", "orderID", newOrder.ID)
slog.Error("Failed to update inventory", err)
```
Example JSON Structures
Order

```json

  "order_id": "order123",
  "customer_name": "Alice Smith",
  "items": [
    {"product_id": "latte", "quantity": 2},
    {"product_id": "muffin", "quantity": 1}
  ],
  "status": "open",
  "created_at": "2023-10-01T09:00:00Z"
}
```
Menu Item

```json
{
  "product_id": "latte",
  "name": "Caffe Latte",
  "description": "Espresso with steamed milk",
  "price": 3.50,
  "ingredients": [
    {"ingredient_id": "espresso_shot", "quantity": 1},
    {"ingredient_id": "milk", "quantity": 200}
  ]
}
```
Inventory Item

```json
{
  "ingredient_id": "milk",
  "name": "Milk",
  "quantity": 5000,
  "unit": "ml"
}
```
Guidelines
Focus first on setting up inventory functionality across all three layers.

Implement clean, maintainable interfaces.

Follow SOLID principles.

Plan the structure carefully before starting coding.
