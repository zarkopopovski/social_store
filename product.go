package main

import (
	"crypto/sha1"
	"fmt"
	"log"
)

type Product struct {
	id             string
	credentials_id string
	store_id       string
	product_code   string
	category       string
	name           string
	description    string
	price          string
	currency       string
}

func (product *Product) CreateNewProduct(dbConnection *DBConnection) bool {

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(product.credentials_id + " " + product.name))
	sha1HashString := sha1Hash.Sum(nil)

	productID := fmt.Sprintf("%x", sha1HashString)

	if err := dbConnection.db.Ping(); err != nil {
		log.Fatal(err)
		return false
	}

	query := "INSERT INTO products(id, credentials_id, stores,id, product_code, category, name, description, price, currency, deleted, date_created, date_modified) " +
		"VALUES('" + productID + "','" + product.credentials_id + "','" + product.store_id + "','" + product.product_code + "', " + product.category + ", '" + product.name + "', '" + product.description + "', '" +
		product.price + "', '" + product.currency + "', 0, NOW(), NOW())"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
