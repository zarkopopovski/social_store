package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DBConnection struct {
	db *sql.DB
}

func OpenConnectionSession() (dbConnection *DBConnection) {
	dbConnection = new(DBConnection)
	dbConnection.createNewDBConnection()

	return
}

func (dbConnection *DBConnection) createNewDBConnection() (err error) {
	connection, err := sql.Open("mysql", "root@/social_store?charset=utf8")
	if err != nil {
		panic(err)
	}

	dbConnection.db = connection

	return
}

func (dbConnection *DBConnection) CreateNewUser(user *User) bool {
	return user.CreateNewUser(dbConnection)
}

func (dbConnection *DBConnection) LoginWithCredentials(email string, password string) *User {
	user := &User{}
	return user.LoginWithCredentials(dbConnection, email, password)
}
