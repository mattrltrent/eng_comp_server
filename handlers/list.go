package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/hn275/eng_comp_server/lib"
)

type Query struct {
	Role    string   `json:"role"`
	Courses []string `json:"courses"`
}

type TutorEntry struct {
	ID          uint   `json:"id"`
	Role        string `json:"role"`
	Course      string `json:"course"`
	Description string `json:"description"`
	User        string `json:"user"`
}

func (d *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	tok, err := GetUser(r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		lib.Err(w, err.Error())
		return
	}
	// parse query
	var q Query
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.Err(w, err.Error())
		return
	}

	var tutorEntries []TutorEntry
	for _, c := range q.Courses {
		var _ []TutorEntry
		d.Table("posts").
			Joins("JOIN users ON posts.user_id=user.id").
			Where("course = ? AND users.id=?", strings.ToUpper(c), tok.Email)

	}
	log.Println(tutorEntries)
}
