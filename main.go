package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
