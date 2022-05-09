package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/signup", signUp)
	initDB()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
