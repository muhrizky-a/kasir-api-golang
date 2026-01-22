package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/domain"
	"net/http"
	"strconv"
	"strings"
)

var products = []domain.Product{
	{ID: 1, Name: "Indomie Godog", Price: 3500, Stock: 10},
	{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
	{ID: 3, Name: "kecap", Price: 12000, Stock: 20},
}

var categories = []domain.Category{
	{ID: 1, Name: "Electronics", Description: "Electronics goods"},
	{ID: 2, Name: "Cosmetics", Description: "Cosmetic goods"},
}

// POST localhost:8080/api/products
func AddProduct(w http.ResponseWriter, r *http.Request) {
	// baca data dari request
	var newProduct domain.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// append products with the new data
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(newProduct)
}

// GET localhost:8080/api/products
func GetProducts(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GET localhost:8080/api/products/{id}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL path e.g. /api/products/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// lookup product by ID with linear search
	for _, product := range products {
		if product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	// Return if product not found
	http.Error(w, "Product not found", http.StatusNotFound)
}

// PUT localhost:8080/api/products/{id}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	/// Parse ID from URL path e.g. /api/products/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// get data from request body
	var updateProduct domain.Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// lookup product by ID with linear search, then update the found product's data
	for i := range products {
		if products[i].ID == id {
			updateProduct.ID = id
			products[i] = updateProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduct)
			return
		}
	}

	// Return if products not found
	http.Error(w, "Product not found", http.StatusNotFound)
}

// DELETE localhost:8080/api/products/{id}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	/// Parse ID from URL path e.g. /api/products/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// lookup product by ID with linear search, then remove the product
	for i, product := range products {
		if product.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			// Create the new slice with data before index(es) and data after index(es)
			products = append(products[:i], products[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "delete product success",
			})
			return
		}
	}

	// Return if product not found
	http.Error(w, "Product not found", http.StatusNotFound)
}

// POST localhost:8080/categories
func AddCategory(w http.ResponseWriter, r *http.Request) {
	// baca data dari request
	var newCategory domain.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// append categories with the new data
	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(newCategory)
}

// GET localhost:8080/categories
func GetCategories(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// GET localhost:8080/categories/{id}
func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID from URL path e.g. /categories/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// lookup category by ID with linear search
	for _, category := range categories {
		if category.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
			return
		}
	}

	// Return if category not found
	http.Error(w, "Category not found", http.StatusNotFound)
}

// PUT localhost:8080/categories/{id}
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	/// Parse ID from URL path e.g. /categories/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// get data from request body
	var updateCategory domain.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// lookup category by ID with linear search, then update the found category's data
	for i := range categories {
		if categories[i].ID == id {
			updateCategory.ID = id
			categories[i] = updateCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateCategory)
			return
		}
	}

	// Return if category not found
	http.Error(w, "Category not found", http.StatusNotFound)
}

// DELETE localhost:8080/categories/{id}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	/// Parse ID from URL path e.g. /categories/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	// lookup category by ID with linear search, then remove the category
	for i, category := range categories {
		if category.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			// Create the new slice with data before index(es) and data after index(es)
			categories = append(categories[:i], categories[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "delete category success",
			})
			return
		}
	}

	// Return if category not found
	http.Error(w, "Category not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc(
		"/health",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "OK",
				"message": "API Running",
			})
		},
	)

	http.HandleFunc("/api/products",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				AddProduct(w, r)
			} else if r.Method == "GET" {
				GetProducts(w)
			} else {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		},
	)

	http.HandleFunc("/api/products/",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" {
				GetProductByID(w, r)
			} else if r.Method == "PUT" {
				UpdateProduct(w, r)
			} else if r.Method == "DELETE" {
				DeleteProduct(w, r)
			} else {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		},
	)

	http.HandleFunc("/categories",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "POST" {
				AddCategory(w, r)
			} else if r.Method == "GET" {
				GetCategories(w)
			} else {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		},
	)

	http.HandleFunc("/categories/",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "GET" {
				GetCategoryByID(w, r)
			} else if r.Method == "PUT" {
				UpdateCategory(w, r)
			} else if r.Method == "DELETE" {
				DeleteCategory(w, r)
			} else {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
		},
	)

	port := "8080"
	fmt.Println("running server on port", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("failed to run server")
	}
}
