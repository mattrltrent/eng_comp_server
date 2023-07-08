package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hn275/eng_comp_server/handlers"
	"github.com/hn275/eng_comp_server/lib"
)

func main() {
	r := chi.NewMux()

	r.Use(middleware.Logger)

	h := handlers.New()

	r.Handle("/auth", http.HandlerFunc(h.SignUp))
	r.Handle("/signin", http.HandlerFunc(h.SignIn))

	r.Route("/", func(r chi.Router) {
		r.Handle("/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := handlers.GetUser(r)
			if err != nil {
				if errors.Is(err, handlers.ErrMissingToken) {
					w.WriteHeader(http.StatusForbidden)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}
				lib.Err(w, err.Error())
				return
			}

			log.Println(user)
			w.WriteHeader(200)
		}))
	})

	log.Println("PORT 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
