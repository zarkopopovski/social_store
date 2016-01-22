package main

import (
	"crypto/sha1"
	"fmt"
	"log"
)

type User struct {
	id           string
	nick_name    string
	email        string
	password     string
	source       string
	status       string
	date_created string
}

func (user *User) CreateNewUser(dbConnection *DBConnection) bool {

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(user.email + " " + user.password))
	sha1HashString := sha1Hash.Sum(nil)

	credentialsID := fmt.Sprintf("%x", sha1HashString)

	if err := dbConnection.db.Ping(); err != nil {
		log.Fatal(err)
		return false
	}

	query := "INSERT INTO credentials(id, nick_name, email, password, source, status, suspended, deleted, date_created, date_modified) VALUES('" + credentialsID + "', '" + user.nick_name + "', '" + user.email + "', '" + user.password + "', '" + user.source + "', 'ENABLE', 'NO', 0, NOW(), NOW())"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (user *User) LoginWithCredentials(dbConnection *DBConnection, email string, password string) *User {

	query := "SELECT id, nick_name, email FROM credentials WHERE email='" + email + "' AND password='" + password + "' AND status='ENABLE'"

	var err = dbConnection.db.QueryRow(query).Scan(&user.id, &user.nick_name, &user.email)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return user
}

type Users []User
