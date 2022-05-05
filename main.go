package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
