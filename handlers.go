package main

import (
	"fmt"
	"net/http"
)

type ApiConnection struct {
	dbConnection *DBConnection
	uHandlers    *UsersHandlers
}

func CreateApiConnection() *ApiConnection {
	API := &ApiConnection{
		dbConnection: OpenConnectionSession(),
		uHandlers:    &UsersHandlers{},
	}
	API.uHandlers.dbConnection = API.dbConnection

	return API
}

func (c *ApiConnection) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
