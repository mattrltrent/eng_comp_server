package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/hn275/eng_comp_server/db"
	"github.com/hn275/eng_comp_server/lib"
)

func (d *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get user info
	token, err := GetUser(r)
	if err != nil {
		if errors.Is(err, ErrMissingToken) {
			w.WriteHeader(http.StatusForbidden)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		lib.Err(w, err.Error())
		return
	}

	user, err := d.GetUserInfo(token)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			lib.Err(w, "user not found")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		lib.Err(w, err.Error())
		return
	}

	// get body
	var post db.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.Err(w, err.Error())
		return
	}
	post.UserID = user.ID
	post.Course = strings.ToUpper(post.Course)

	// save to db
	if err := d.Table(db.TablePosts).Create(&post).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.Err(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}
