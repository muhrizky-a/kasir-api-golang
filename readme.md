# Cashier API

## Testing (Windows)

curl http://localhost:8080/health

curl http://localhost:8080/categories

curl -X POST http://localhost:8080/categories ^
-H "Content-Type: application/json" ^
-d "{
\"name\": \"Sports\",
\"description\": \"Sport goods\" }"

curl -X POST https://kasir-api-golang-production-c6bc.up.railway.app/categories ^
-H "Content-Type: application/json" ^
-d "{ \"name\": \"Sports\", \"description\": \"Sport goods\" }"

curl http://localhost:8080/categories

curl -X PUT http://localhost:8080/categories/2 ^
-H "Content-Type: application/json" ^
-d "{
\"name\": \"Toys\",
\"description\": \"Toys for kids\" }"

curl -X DELETE http://localhost:8080/categories/1
