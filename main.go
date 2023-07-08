package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hn275/eng_comp_server/handlers"
)

func main() {
	r := chi.NewMux()

	r.Use(middleware.Logger)

	h := handlers.New()

	r.Handle("/auth", http.HandlerFunc(h.SignUp))

	log.Println("PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
