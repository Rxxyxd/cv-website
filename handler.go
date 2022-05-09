package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Password string `json:"password", db:"password"`
	Username string `json:"username", db:"username"`
}

func signUp(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Salt and hash the password using bycrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	// add the username and hashed password into the database
	if _, err = db.Query("insert into users values ($1, $2)", creds.Username, string(hashedPassword)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := db.QueryRow("select password from users where username=$1", creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	storeCreds := &Credentials{}

	// store obtained password in 'storedCreds'
	err = result.Scan(&storeCreds.Password)
	if err != nil {
		//if an entry with the username does not exist, send an "Unauthorized" 401 Status
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If any other error type occurs, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password with the one that was recieved
	if err = bcrypt.CompareHashAndPassword([]byte(storeCreds.Password), []byte(creds.Password)); err != nil {
		// if the passwords don't match return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
	}
}
