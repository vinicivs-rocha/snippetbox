package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", viewSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
