package handlers

import (
	"log"
	"net/http"
)

type Query struct {
	Role    string   `json:"role"`
	Courses []string `json:"courses"`
}

func (d *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	courses := r.URL.Query().Get("course")
	log.Println(courses)

}
