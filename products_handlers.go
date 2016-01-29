package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

	file, header, err := r.FormFile("file")

	if err != nil {
	}

	defer file.Close()

	timeNow := time.Now()
	nanoTime := timeNow.UnixNano()
	nanoInMillis := nanoTime / 100000

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(strconv.FormatInt(nanoInMillis, 10) + " " + token + " " + storeID + " " + header.Filename))
	sha1HashString := sha1Hash.Sum(nil)

	fileHash := fmt.Sprintf("%x", sha1HashString)

	hashedFileName := fileHash + "" + header.Filename

	out, err := os.Create("/Users/zarkopopovski/Documents/GOWorkspace/SocialStore/uploads/" + hashedFileName)

	if err != nil {
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {
	}

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

	file, header, err := r.FormFile("file")

	if err != nil {
	}

	defer file.Close()

	timeNow := time.Now()
	nanoTime := timeNow.UnixNano()
	nanoInMillis := nanoTime / 100000

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(strconv.FormatInt(nanoInMillis, 10) + " " + token + " " + storeID + " " + header.Filename))
	sha1HashString := sha1Hash.Sum(nil)

	fileHash := fmt.Sprintf("%x", sha1HashString)

	hashedFileName := fileHash + "" + header.Filename

	out, err := os.Create("/Users/zarkopopovski/Documents/GOWorkspace/SocialStore/uploads/" + hashedFileName)

	if err != nil {
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {
	}

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

}
