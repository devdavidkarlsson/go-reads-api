package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (e endpoint) bookRouter() (r chi.Router) {
	b := bookEndpoint{e}
	r = chi.NewRouter()
	r.Get("/{id}", handle(b.get))
	return
}

type bookEndpoint struct {
	endpoint
}

func (e bookEndpoint) get(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return e.book.Get(r.Context(), chi.URLParam(r, "id"))
}
