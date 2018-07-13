package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/devdavidkarlsson/rest-api/internal/model"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	client := &http.Client{}

	//test_user_id := "54455429"
	key := "IeCzEEdRf8p7FbB1A9ElEQ"
	e := endpoint{book: model.BookConfig{client}.Create(key), user: model.UserConfig{client}.Create(key)} //Create book-connector

	fmt.Print("Api started")
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/book", e.bookRouter())
		r.Mount("/user", e.userRouter())
	})x

	log.Fatal(http.ListenAndServe(":3000", r))
}
