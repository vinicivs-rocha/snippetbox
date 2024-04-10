package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func ViewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Displaying snippet with id %d", id)
}
