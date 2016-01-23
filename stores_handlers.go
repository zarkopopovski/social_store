package main

import (
	"encoding/json"
	"net/http"
)

type StoresHandlers struct {
	dbConnection *DBConnection
}

func (sHandlers *StoresHandlers) CreateStore(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	name := r.FormValue("name")
	address := r.FormValue("address")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	country := r.FormValue("country")
	tel := r.FormValue("tel")

	photo := r.FormValue("photo")

	store := &Store{credentials_id: token, name: name, address: address, city: city, zip: zip, country: country, tel: tel, photo: photo}

	result := store.CreateNewStore(sHandlers.dbConnection)

	//result := sHandlers.dbConnection.CreateNewStore(store)

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

func (sHandlers *StoresHandlers) UpdateStore(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	storeId := r.FormValue("store_id")
	name := r.FormValue("name")
	address := r.FormValue("address")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	country := r.FormValue("country")
	tel := r.FormValue("tel")
	photo := r.FormValue("photo")

	store := &Store{id: storeId, credentials_id: token, name: name, address: address, city: city, zip: zip, country: country, tel: tel, photo: photo}

	result := store.UpdateStoreDetails(sHandlers.dbConnection)

	//result := sHandlers.dbConnection.UpdateStoreDetails(store)

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