# Cashier API

## Testing Product (Windows)

curl http://localhost:8080/health

curl http://localhost:8080/api/products

curl -X POST http://localhost:8080/api/products ^
-H "Content-Type: application/json" ^
-d "{ \"name\": \"Rice\", \"price\": 5000, \"stock\": 42, \"category_id\": 42 }"

curl http://localhost:8080/api/products
curl http://localhost:8080/api/products?name=ice

curl http://localhost:8080/api/products/1

curl -X PUT http://localhost:8080/api/products/1 ^
-H "Content-Type: application/json" ^
-d "{ \"name\": \"Flowers\", \"price\": 42000, \"stock\": 41 }"

curl -X DELETE http://localhost:8080/api/products/5

## Testing Category (Windows)

curl http://localhost:8080/health

curl -X POST http://localhost:8080/api/categories ^
-H "Content-Type: application/json" ^
-d "{ \"name\": \"Sports\", \"description\": \"Sport goods\" }"

curl http://localhost:8080/api/categories

curl http://localhost:8080/api/categories/1

curl http://localhost:8080/api/categories

curl -X PUT http://localhost:8080/api/categories/1 ^
-H "Content-Type: application/json" ^
-d "{ \"name\": \"Toys\", \"description\": \"Toys for kids\" }"

curl -X DELETE http://localhost:8080/api/categories/1
