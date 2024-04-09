package handlers

import "net/http"

func ViewSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display some snippet..."))
}
