package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type UsersHandlers struct {
	dbConnection *DBConnection
}

func (uHandlers *UsersHandlers) SignIn(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	source := r.FormValue("source")

	user := &User{nick_name: name, email: email, password: password, source: source}

	result := user.CreateNewUser(uHandlers.dbConnection)

	if result {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (uHandlers *UsersHandlers) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := &User{}
	user = user.LoginWithCredentials(uHandlers.dbConnection, email, password)

	if user != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		userMap := make(map[string]string)
		userMap["token"] = user.id

		if err := json.NewEncoder(w).Encode(userMap); err != nil {
			panic(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (uHandlers *UsersHandlers) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	tel1 := r.FormValue("tel1")
	tel2 := r.FormValue("tel2")

	imageName := r.FormValue("image_name")

	timeNow := time.Now()
	nanoTime := timeNow.UnixNano()
	nanoInMillis := nanoTime / 100000

	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(strconv.FormatInt(nanoInMillis, 10) + "" + token + "" + firstName + lastName + "" + imageName))
	sha1HashString := sha1Hash.Sum(nil)

	fileHash := fmt.Sprintf("%x", sha1HashString)

	hashedFileName := fileHash + "" + imageName

	user := &User{}
	userDetails := &UserDetails{credentials_id: token, first_name: firstName, last_name: lastName, tel1: tel1, tel2: tel2, avatar: hashedFileName}

	result := user.UpdateUserProfile(uHandlers.dbConnection, userDetails)

	if result {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}

func (uHandlers *UsersHandlers) ReadUserProfile(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")

	user := &User{}
	userDetails := user.ReadUserProfile(uHandlers.dbConnection, token)

	if userDetails != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		userMap := map[string]string{
			"first_name": userDetails.first_name,
			"last_name":  userDetails.last_name,
			"tel1":       userDetails.tel1,
			"tel2":       userDetails.tel2,
			"avatar":     userDetails.avatar,
		}

		if err := json.NewEncoder(w).Encode(userMap); err != nil {
			panic(err)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}
