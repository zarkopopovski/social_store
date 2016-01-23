package main

import (
	"crypto/sha1"
	"fmt"
	"log"
)

type User struct {
	id        string
	nick_name string
	email     string
	password  string
	source    string
	status    string
}

type UserDetails struct {
	credentials_id string
	first_name     string
	last_name      string
	tel1           string
	tel2           string
	avatar         string
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

	query := "INSERT INTO credentials(id, nick_name, email, password, source, status, suspended, deleted, date_created, date_modified) " +
		"VALUES('" + credentialsID + "', '" + user.nick_name + "', '" + user.email + "', '" + user.password + "', '" + user.source + "', 'ENABLE', 'NO', 0, NOW(), NOW())"

	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (user *User) LoginWithCredentials(dbConnection *DBConnection, email string, password string) *User {

	query := "SELECT id, nick_name, email FROM credentials WHERE email='" + email + "' " +
		"AND password='" + password + "' AND status='ENABLE'"

	err := dbConnection.db.QueryRow(query).Scan(&user.id, &user.nick_name, &user.email)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return user
}

func (user *User) UpdateUserProfile(dbConnection *DBConnection, userDetails *UserDetails) bool {

	testUserDetails := user.ReadUserProfile(dbConnection, userDetails.credentials_id)

	query := "UPDATE user_details SET first_name='" + userDetails.first_name + "', last_name='" + userDetails.last_name + "', tel1='" + userDetails.tel1 + "', tel2='" + userDetails.tel2 + "', avatar='" + userDetails.avatar + "', date_modified=NOW() WHERE credentials_id='" + userDetails.credentials_id + "'"

	if testUserDetails == nil {
		query = "INSERT INTO user_details(credentials_id, first_name, last_name, tel1, tel2, avatar, date_created, date_modified) " +
			"VALUES('" + userDetails.credentials_id + "', '" + userDetails.first_name + "','" + userDetails.last_name + "','" + userDetails.tel1 + "','" + userDetails.tel2 + "','" + userDetails.avatar + "',NOW(), NOW())"
	}
	_, err := dbConnection.db.Exec(query)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (user *User) ReadUserProfile(dbConnection *DBConnection, token string) *UserDetails {

	userDetails := UserDetails{}

	query := "SELECT first_name, last_name, tel1, tel2, avatar FROM user_details WHERE credentials_id='" + token + "'"

	err := dbConnection.db.QueryRow(query).Scan(&userDetails.first_name, &userDetails.last_name, &userDetails.tel1, &userDetails.tel2, &userDetails.avatar)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &userDetails
}

type Users []User
