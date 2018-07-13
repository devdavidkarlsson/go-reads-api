package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (e endpoint) userRouter() (r chi.Router) {
	b := userEndpoint{e}
	r = chi.NewRouter()
	r.Get("/{id}", handle(b.get))
	return
}

type userEndpoint struct {
	endpoint
}

func (e userEndpoint) get(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return e.user.Get(r.Context(), chi.URLParam(r, "id"))
}
