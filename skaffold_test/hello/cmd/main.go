package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", urlHandler)

	log.Println("Listening on port 8081")
	http.ListenAndServe(":8081", nil)
}

func urlHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "got request to /hello")
}
