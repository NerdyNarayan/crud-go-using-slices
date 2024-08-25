package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Schema
type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Seller      *Seller `json:"seller"`
}
type Seller struct {
	ID          string `json:"id"`
	Store_Name  string `json:"store-name"`
	Description string `json:"description"`
}

// Slice array
var products []Product

// getProducts
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// deleteProduct
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, product := range products {
		if product.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(products)
}

// getProduct
func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, product := range products {
		if product.ID == params["id"] {
			json.NewEncoder(w).Encode(product)
		}
	}
}

// addProduct
func addProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(10000))
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

// updateProduct
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for index, product := range products {
		if product.ID == id {
			products = append(products[:index], products[index+1:]...)
		}
	}
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = id
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

func main() {
	//Dummy Data
	products = append(products, Product{ID: "1", Name: "Product 1", Price: 10.99, Description: "This is product 1", Seller: &Seller{ID: "1", Store_Name: "Seller 1", Description: "This is seller 1"}})
	products = append(products, Product{ID: "2", Name: "Product 2", Price: 9.99, Description: "This is product 2", Seller: &Seller{ID: "2", Store_Name: "Seller 2", Description: "This is seller 2"}})

	r := mux.NewRouter()
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/products", addProduct).Methods("POST")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products/{id}", updateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	fmt.Printf("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
