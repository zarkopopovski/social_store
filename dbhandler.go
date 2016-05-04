package main

import (
	"database/sql"
	"fmt"
	"github.com/fzzy/radix/extra/pool"
	_ "github.com/go-sql-driver/mysql"
)

type DBConnection struct {
	db   *sql.DB
	pool *pool.Pool
}

func OpenConnectionSession() (dbConnection *DBConnection) {
	dbConnection = new(DBConnection)
	dbConnection.createNewDBConnection()
	dbConnection.createNewCacheConnection()

	return
}

func (dbConnection *DBConnection) createNewDBConnection() (err error) {
	connection, err := sql.Open("mysql", "root@/social_store?charset=utf8")
	if err != nil {
		panic(err)
	}

	fmt.Println("MySQL Connection is Active")
	dbConnection.db = connection

	return
}

func (dbConnection *DBConnection) createNewCacheConnection() (err error) {
	pool, err := pool.NewPool("tcp", "127.0.0.1:6379", 10)

	fmt.Println("Redis Connection is Active")
	dbConnection.pool = pool

	return
}
