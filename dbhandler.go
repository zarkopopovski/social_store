package main

import (
	"database/sql"
	"fmt"
	"github.com/fzzy/radix/redis"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DBConnection struct {
	db     *sql.DB
	client *redis.Client
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
	client, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis Connection is Active")
	dbConnection.client = client

	return
}
