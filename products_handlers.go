package main

import (
	"encoding/json"
	"net/http"
)

type ProductsHandlers struct {
	dbConnection *DBConnection
}

func (pHandlers *ProductsHandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	storeID := r.FormValue("store_id")
	productCode := r.FormValue("product_code")
	category := r.FormValue("category")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	currency := r.FormValue("currency")

	product := &Product{credentials_id: token, store_id: storeID, product_code: productCode, category: category, name: name, description: description, price: price, currency: currency}

	result := product.CreateNewProduct(pHandlers.dbConnection)

	if result {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (pHandlers *ProductsHandlers) UpdateProduct(w http.ResponseWriter, r *http.Request) {

}

func (pHandlers *ProductsHandlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {

}

func (pHandlers *ProductsHandlers) ListProductsByStore(w http.ResponseWriter, r *http.Request) {

}
