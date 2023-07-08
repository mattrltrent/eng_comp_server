package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hn275/eng_comp_server/db"
	"github.com/hn275/eng_comp_server/lib"
)

type Post struct {
	Email  string `json:"email"`
	Role   string `json:"role"`
	Course string `json:"course"`
}

func (d *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get body
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.Err(w, err.Error())
		return
	}

	// save to db
	if err := d.Table(db.TablePosts).Create(&post).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.Err(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
