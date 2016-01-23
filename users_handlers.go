package main

import (
	"encoding/json"
	"net/http"
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

	result := uHandlers.dbConnection.CreateNewUser(user)

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

	user := uHandlers.dbConnection.LoginWithCredentials(email, password)

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
	avatar := r.FormValue("avatar")

	userDetails := &UserDetails{credentials_id: token, first_name: firstName, last_name: lastName, tel1: tel1, tel2: tel2, avatar: avatar}

	result := uHandlers.dbConnection.UpdateUserProfile(userDetails)

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

	userDetails := uHandlers.dbConnection.ReadUserProfile(token)

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
