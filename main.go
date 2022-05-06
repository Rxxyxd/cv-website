package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	if err := http.ListenAndServeTLS(":443", "./static/certs/certificate.crt", "./static/certs/private.key", nil); err != nil {
		fmt.Println(err)
	}
}
