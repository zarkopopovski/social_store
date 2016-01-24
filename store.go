package main

import (
	"crypto/sha1"
	"fmt"
	"log"
)

type Store struct {
	id             string
	credentials_id string
	name           string
	address        string
	city           string
	zip            string
	country        string
	tel            string
	photo          string
	lat            string
	lon            string
}

func (store *Store) CreateNewStore(dbConnection *DBConnection) bool {

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(store.credentials_id + " " + store.name))
	sha1HashString := sha1Hash.Sum(nil)

	storeID := fmt.Sprintf("%x", sha1HashString)

	if err := dbConnection.db.Ping(); err != nil {
		log.Fatal(err)
		return false
	}

	query := "INSERT INTO stores(id, credentials_id, name, address, city, zip, country, tel, photo, lat, lon, deleted, date_created, date_modified) " +
		"VALUES('" + storeID + "', '" + store.credentials_id + "', '" + store.name + "', '" + store.address + "', '" + store.city + "', '" + store.zip + "', " +
		store.country + ", '" + store.tel + "', '" + store.photo + "', '" + store.lat + "', '" + store.lon + "', 0, NOW(), NOW())"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (store *Store) UpdateStoreDetails(dbConnection *DBConnection) bool {

	query := "UPDATE store SET name='" + store.name + "', address='" + store.address + "', city='" + store.city + "', zip='" + store.zip + "', country=" + store.country +
		", tel='" + store.tel + "', photo='" + store.photo + "', lat='" + store.lat + "', lon='" + store.lon + "', date_modified=NOW() WHERE id='" + store.id + "' AND credentials_id='" + store.credentials_id + "'"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (store *Store) DeleteExistingStore(dbConnection *DBConnection) bool {

	query := "UPDATE store SET deleted=1, date_modified=NOW() WHERE id='" + store.id + "' AND credentials_id='" + store.credentials_id + "'"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (store *Store) ListPersonalExistingStores(dbConnection *DBConnection) []*Store {

	query := "SELECT id, credentials_id, name, address, city, zip, country, tel, photo FROM stores WHERE credentials_id='" + store.credentials_id + "' AND deleted=0"

	rows, err := dbConnection.db.Query(query)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer rows.Close()

	stores := make([]*Store, 0)

	for rows.Next() {

		newStore := new(Store)

		err := rows.Scan(&newStore.id, &newStore.credentials_id, &newStore.name, &newStore.address, &newStore.city, &newStore.zip, &newStore.country, &newStore.tel, &newStore.photo)

		if err != nil {
			log.Fatal(err)
			return nil
		}

		stores = append(stores, newStore)

	}

	return stores
}

type Stores []Store
