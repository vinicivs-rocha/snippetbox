package main

import (
	"log"
	"net/http"
	"snippetbox.vinicivs-rocha.com/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/snippet/view", handlers.ViewSnippet)
	mux.HandleFunc("/snippet/create", handlers.CreateSnippet)

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
