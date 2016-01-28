package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
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

	latitude := r.FormValue("lat")
	longitude := r.FormValue("lon")

	file, header, err := r.FormFile("file")

	if err != nil {
	}

	defer file.Close()

	out, err := os.Create("/Users/zarkopopovski/Documents/GOWorkspace/receive/uploads/" + header.Filename)

	if err != nil {
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {
	}

	store := &Store{credentials_id: token, name: name, address: address, city: city, zip: zip, country: country, tel: tel, photo: header.Filename, lat: latitude, lon: longitude}

	result := store.CreateNewStore(sHandlers.dbConnection)

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
	storeID := r.FormValue("store_id")
	name := r.FormValue("name")
	address := r.FormValue("address")
	city := r.FormValue("city")
	zip := r.FormValue("zip")
	country := r.FormValue("country")
	tel := r.FormValue("tel")

	store := &Store{id: storeID, credentials_id: token, name: name, address: address, city: city, zip: zip, country: country, tel: tel}

	result := store.UpdateStoreDetails(sHandlers.dbConnection)

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

func (sHandlers *StoresHandlers) UpdateStorePhoto(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	storeID := r.FormValue("store_id")

	file, header, err := r.FormFile("file")

	if err != nil {
	}

	defer file.Close()

	out, err := os.Create("/Users/zarkopopovski/Documents/GOWorkspace/receive/uploads/" + header.Filename)

	if err != nil {
	}

	defer out.Close()

	_, err = io.Copy(out, file)

	if err != nil {
	}

	store := &Store{id: storeID, credentials_id: token, photo: header.Filename}

	result := store.UpdateStorePhoto(sHandlers.dbConnection)

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

func (sHandlers *StoresHandlers) DeleteStore(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	storeID := r.FormValue("store_id")

	store := &Store{id: storeID, credentials_id: token}

	result := store.DeleteExistingStore(sHandlers.dbConnection)

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

func (sHandlers *StoresHandlers) ListPersonalStores(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")

	store := &Store{credentials_id: token}
	stores := store.ListPersonalExistingStores(sHandlers.dbConnection)

	if stores != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(stores); err != nil {
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

func (sHandlers *StoresHandlers) ListStores(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	pageID, _ := strconv.Atoi(r.FormValue("pageID"))

	store := &Store{credentials_id: token}
	stores := store.ListStoresByPages(sHandlers.dbConnection, pageID)

	if stores != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(stores); err != nil {
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
