package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/devdavidkarlsson/rest-api/internal/model"
)

type endpoint struct {
	book model.Book
	user model.User
}

// unless otherwise specified by error interface the error will result in a 500 with logging but no output
type handler func(http.ResponseWriter, *http.Request) (obj interface{}, err error)

func handle(fn handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj, err := fn(w, r)

		if err != nil {
			fmt.Printf(err.Error())
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		if obj != nil {
			bytes, err := json.Marshal(obj)

			if err != nil {
				fmt.Printf(err.Error())
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(bytes)
		}
	}
}
