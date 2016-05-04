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
	image_name     string
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

	query := "INSERT INTO products(id, credentials_id, stores_id, product_code, category, name, description, price, currency, deleted, date_created, date_modified) " +
		"VALUES('" + productID + "','" + product.credentials_id + "','" + product.store_id + "','" + product.product_code + "', " + product.category + ", '" + product.name + "', '" + product.description + "', '" +
		product.price + "', '" + product.currency + "', 0, NOW(), NOW())"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	sha1Hash.Write([]byte(product.credentials_id + " " + product.image_name))
	sha1HashString = sha1Hash.Sum(nil)

	productImageID := fmt.Sprintf("%x", sha1HashString)

	query = "INSERT INTO products_gallery(id, credentials_id, stores_id, products_id, image_name, main_photo, deleted, date_created, date_modified) " +
		"VALUES('" + productImageID + "','" + product.credentials_id + "','" + product.store_id + "','" + product.id + "','" + product.image_name + "','YES',0,NOW(), NOW())"

	_, err = dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true

}

func (product *Product) UpdateExistingProduct(dbConnection *DBConnection) bool {

	query := "UPDATE products SET product_code='" + product.product_code + "', category=" + product.category + ", name='" + product.name + "', description='" + product.description + "', price='" + product.price + "', currency='" + product.currency + "' " +
		"WHERE id='" + product.id + "'"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true

}

func (product *Product) CreateNewProductPhoto(dbConnection *DBConnection) bool {

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(product.credentials_id + " " + product.name + " " + product.image_name))
	sha1HashString := sha1Hash.Sum(nil)

	productImageID := fmt.Sprintf("%x", sha1HashString)

	if err := dbConnection.db.Ping(); err != nil {
		log.Fatal(err)
		return false
	}

	query := "INSERT INTO products_gallery(id, credentials_id, stores_id, products_id, image_name, main_photo, deleted, date_created, date_modified) " +
		"VALUES('" + productImageID + "','" + product.credentials_id + "','" + product.store_id + "','" + product.id + "','" + product.image_name + "','NO',0,NOW(), NOW())"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true

}

func (product *Product) DeleteProductPhoto(dbConnection *DBConnection, photoID string) bool {

	query := "UPDATE products_gallery SET deleted=1, date_modified=NOW() WHERE id='" + photoID + "'"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true

}

func (product *Product) DeleteExistingProduct(dbConnection *DBConnection) bool {

	query := "UPDATE products SET deleted=1 WHERE id='" + product.id + "'"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true

}

func (product *Product) ListAllProductsByStore(dbConnection *DBConnection) []*Product {

	query := "SELECT p.id, p.credentials_id, p.stores_id, p.product_code, p.category, p.name, p.description, p.price, p.currency, pg.image_name FROM products p LEFT JOIN products_gallery pg " +
		"ON p.stores_id='" + product.store_id + "' AND p.deleted=0 AND p.id=pg.products_id AND pg.main_photo='YES' AND p.deleted=0"

	rows, err := dbConnection.db.Query(query)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer rows.Close()

	products := make([]*Product, 0)

	for rows.Next() {

		newProduct := new(Product)

		err := rows.Scan(&newProduct.id, &newProduct.credentials_id, &newProduct.store_id, &newProduct.category, &newProduct.name, &newProduct.description, &newProduct.price, &newProduct.currency, &newProduct.image_name)

		if err != nil {
			log.Fatal(err)
			return nil
		}

		products = append(products, newProduct)

	}

	return products

}

func (product *Product) SetLikeToProduct(dbConnection *DBConnection) bool {
	client, err := dbConnection.cache.Get()

	if err != nil {
		log.Fatal(err)
		return false
	}

	defer dbConnection.cache.Put(client)

	r := client.Cmd("select", 8)

	if r.Err != nil {
		fmt.Println("Error selecting DB")
		log.Fatal(r.Err)
		return false
	}

	r = client.Cmd("sadd", "Product:Likes:"+product.id, product.credentials_id)

	if r.Err != nil {
		log.Fatal(r.Err)
		return false
	}

	return true
}

func (product *Product) RemoveLikeFromProduct(dbConnection *DBConnection) bool {
	client, err := dbConnection.cache.Get()

	if err != nil {
		log.Fatal(err)
		return false
	}

	defer dbConnection.cache.Put(client)

	r := client.Cmd("select", 8)

	if r.Err != nil {
		log.Fatal(r.Err)
		return false
	}

	r = client.Cmd("srem", "Product:Likes:"+product.id, product.credentials_id)

	if r.Err != nil {
		log.Fatal(r.Err)
		return false
	}

	return true
}

func (product *Product) ReadProductsLikes(dbConnection *DBConnection) []string {
	client, err := dbConnection.cache.Get()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer dbConnection.cache.Put(client)

	r := client.Cmd("select", 8)

	if r.Err != nil {
		log.Fatal(r.Err)
		return nil
	}

	r = client.Cmd("smembers", "Product:Likes:"+product.id)

	if r.Err != nil {
		log.Fatal(r.Err)
		return nil
	}

	vals, _ := r.List()

	return vals
}
