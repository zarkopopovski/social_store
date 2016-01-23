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

	query := "INSERT INTO stores(id, credentials_id, name, address, city, zip, country, tel, photo, deleted, date_created, date_modified) " +
		"VALUES('" + storeID + "', " + store.credentials_id + "', '" + store.name + "', '" + store.address + "', '" + store.city + "', '" + store.zip + "', " + string(store.country) + ", '" + store.tel + "', '" + store.photo + "', 0, NOW(), NOW())"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (store *Store) UpdateStoreDetails(dbConnection *DBConnection) bool {

	query := "UPDATE store SET name='" + store.name + "', address='" + store.address + "', city='" + store.city + "', zip='" + store.zip + "', country=" + string(store.country) + ", tel='" + store.tel + "', photo='" + store.photo + "', date_modified=NOW() " +
		"WHERE id='" + store.id + "' AND credentials_id='" + store.credentials_id + "'"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
