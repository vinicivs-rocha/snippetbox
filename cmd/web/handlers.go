package main

import (
	"errors"
	"fmt"
	// "html/template"
	"net/http"
	"snippetbox.vinicivs-rocha.com/internal/models"
	"strconv"
)

func (a *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		a.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := a.snippetsRepo.Insert(title, content, expires)
	if err != nil {
		a.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}

func (a *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}

	snps, err := a.snippetsRepo.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}

	for _, snp := range snps {
		fmt.Fprintf(w, "%+v\n", snp)
	}
}

func (a *application) viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id <= 0 {
		a.clientError(w, http.StatusBadRequest)
		return
	}

	snp, err := a.snippetsRepo.Get(id)

	if err != nil && errors.Is(err, models.ErrNoRecord) {
		a.notFound(w)
		return
	}

	if err != nil {
		a.serverError(w, err)
		return
	}

	fmt.Fprintf(w, "%+v", snp)
}
