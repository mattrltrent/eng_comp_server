package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/hn275/eng_comp_server/db"
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

	var post Post
	// get body
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		fmt.Fprint(os.Stderr, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// save to db

	if err := d.Table(db.TablePosts).Create(&post).Error; err != nil {
		fmt.Fprint(os.Stderr, err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "duplicate email"})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
