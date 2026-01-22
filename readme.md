# Cashier API

## Testing Product (Windows)

curl http://localhost:8080/health

curl http://localhost:8080/api/products

curl -X POST http://localhost:8080/api/products ^
-H "Content-Type: application/json" ^
-d "{ \"name\": \"Bag\", \"price\": 50000, \"stock\": 42 }"

curl http://localhost:8080/api/products

curl -X PUT http://localhost:8080/api/products/2 ^
-H "Content-Type: application/json" ^
-d "{ \"name\": \"Flowers\", \"price\": 42000, \"stock\": 41 }"

curl -X DELETE http://localhost:8080/api/products/1

## Testing Category (Windows)

curl http://localhost:8080/health

curl http://localhost:8080/categories

curl -X POST http://localhost:8080/categories ^
-H "Content-Type: application/json" ^
-d "{ \"name\": \"Sports\", \"description\": \"Sport goods\" }"

curl http://localhost:8080/categories

curl -X PUT http://localhost:8080/categories/2 ^
-H "Content-Type: application/json" ^
-d "{ \"name\": \"Toys\", \"description\": \"Toys for kids\" }"

curl -X DELETE http://localhost:8080/categories/1
