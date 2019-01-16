package main

import (
	"fmt"
	"html"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Implement REST API

	fmt.Println("Starting server:")
	r := mux.NewRouter()

	r.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	})

	// POST   /user/register
	// POST   /user/login
	// POST   /deck
	// GET    /deck
	// PUT    /deck
	// DELETE /deck
	// GET    /deck/deck_id
	// POST   /deck/share
	// POST   /deck/deck_id/card
	// PUT    /deck/deck_id/card/card_id
	// DELETE /deck/deck_id/card/card_id
	// GET    /deck/deck_id/study
	// POST   /deck/deck_id/study

	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)

}
