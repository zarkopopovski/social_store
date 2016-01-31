package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
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

	imageName := r.FormValue("image_name")

	timeNow := time.Now()
	nanoTime := timeNow.UnixNano()
	nanoInMillis := nanoTime / 100000

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(strconv.FormatInt(nanoInMillis, 10) + "" + token + "" + name + "" + imageName))
	sha1HashString := sha1Hash.Sum(nil)

	fileHash := fmt.Sprintf("%x", sha1HashString)

	hashedFileName := fileHash + "" + imageName

	store := &Store{credentials_id: token, name: name, address: address, city: city, zip: zip, country: country, tel: tel, photo: hashedFileName, lat: latitude, lon: longitude}

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

	imageName := r.FormValue("image_name")

	timeNow := time.Now()
	nanoTime := timeNow.UnixNano()
	nanoInMillis := nanoTime / 100000

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(strconv.FormatInt(nanoInMillis, 10) + "" + token + "" + storeID + "" + imageName))
	sha1HashString := sha1Hash.Sum(nil)

	fileHash := fmt.Sprintf("%x", sha1HashString)

	hashedFileName := fileHash + "" + imageName

	store := &Store{id: storeID, credentials_id: token, photo: hashedFileName}

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
