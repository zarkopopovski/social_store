package main

import (
	"fmt"
	"net/http"
)

type ApiConnection struct {
	dbConnection *DBConnection
	uHandlers    *UsersHandlers
	sHandlers    *StoresHandlers
}

func CreateApiConnection() *ApiConnection {
	API := &ApiConnection{
		dbConnection: OpenConnectionSession(),
		uHandlers:    &UsersHandlers{},
		sHandlers:    &StoresHandlers{},
	}
	API.uHandlers.dbConnection = API.dbConnection
	API.sHandlers.dbConnection = API.dbConnection

	return API
}

func (c *ApiConnection) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}
