package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
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

	imageName := r.FormValue("image_name")

	timeNow := time.Now()
	nanoTime := timeNow.UnixNano()
	nanoInMillis := nanoTime / 100000

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(strconv.FormatInt(nanoInMillis, 10) + "" + token + "" + storeID + "" + imageName))
	sha1HashString := sha1Hash.Sum(nil)

	fileHash := fmt.Sprintf("%x", sha1HashString)

	hashedFileName := fileHash + "" + imageName

	product := &Product{credentials_id: token, store_id: storeID, product_code: productCode, category: category, name: name, description: description, price: price, currency: currency, image_name: hashedFileName}

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
	token := r.FormValue("token")
	productID := r.FormValue("product_id")
	storeID := r.FormValue("store_id")
	productCode := r.FormValue("product_code")
	category := r.FormValue("category")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	currency := r.FormValue("currency")

	product := &Product{id: productID, credentials_id: token, store_id: storeID, product_code: productCode, category: category, name: name, description: description, price: price, currency: currency}

	result := product.UpdateExistingProduct(pHandlers.dbConnection)

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

func (pHandlers *ProductsHandlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	productID := r.FormValue("product_id")

	product := &Product{id: productID, credentials_id: token}

	result := product.DeleteExistingProduct(pHandlers.dbConnection)

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

func (pHandlers *ProductsHandlers) CreateProductPhoto(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	productID := r.FormValue("product_id")
	storeID := r.FormValue("store_id")

	imageName := r.FormValue("image_name")

	timeNow := time.Now()
	nanoTime := timeNow.UnixNano()
	nanoInMillis := nanoTime / 100000

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(strconv.FormatInt(nanoInMillis, 10) + "" + token + "" + storeID + "" + imageName))
	sha1HashString := sha1Hash.Sum(nil)

	fileHash := fmt.Sprintf("%x", sha1HashString)

	hashedFileName := fileHash + "" + imageName

	product := &Product{id: productID, credentials_id: token, store_id: storeID, image_name: hashedFileName}

	result := product.CreateNewProductPhoto(pHandlers.dbConnection)

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

func (pHandlers *ProductsHandlers) ListProductsByStore(w http.ResponseWriter, r *http.Request) {
	storeID := r.FormValue("store_id")

	product := &Product{store_id: storeID}
	products := product.ListAllProductsByStore(pHandlers.dbConnection)

	if products != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(products); err != nil {
			panic(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (pHandlers *ProductsHandlers) SetLikeProduct(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	productID := r.FormValue("product_id")

	product := &Product{id: productID, credentials_id: token}
	result := product.SetLikeToProduct(pHandlers.dbConnection)

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

func (pHandlers *ProductsHandlers) RemoveLikeProduct(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	productID := r.FormValue("product_id")

	product := &Product{id: productID, credentials_id: token}
	result := product.RemoveLikeFromProduct(pHandlers.dbConnection)

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

func (pHandlers *ProductsHandlers) ReadProductLikes(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	productID := r.FormValue("product_id")

	product := &Product{id: productID, credentials_id: token}
	result := product.ReadProductsLikes(pHandlers.dbConnection)

	if result != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(result); err != nil {
			panic(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
