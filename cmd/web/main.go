package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", viewSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
