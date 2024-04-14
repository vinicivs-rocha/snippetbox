package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (a *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		a.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create some snippet..."))
}

func (a *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}

	files := []string{
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/pages/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		a.errLogger.Println(err.Error())
		a.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		a.errLogger.Println(err.Error())
		a.serverError(w, err)
		return
	}
}

func (a *application) viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		a.clientError(w, http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Displaying snippet with id %d", id)
}
