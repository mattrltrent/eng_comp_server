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

func (d *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	tok, err := GetUser(r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		lib.Err(w, err.Error())
		return
	}

	user, err := d.GetUserInfo(tok)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.Err(w, err.Error())
		return
	}
	log.Println(user)

	// parse query
	var q Query
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		lib.Err(w, err.Error())
		return
	}

	// query all posts with matching courses
	orClause := []string{}
	for idx, course := range q.Courses {
		q.Courses[idx] = strings.ToUpper(course)
		orClause = append(orClause, " course = ? ")
	}
	or := strings.Join(orClause, "OR")
	join := "JOIN users ON posts.user_id=user.id"

	sql := "SELECT posts.id,posts.role,posts.course,posts.description,users.user FROM posts"
	sql += " WHERE post_user"
	sql += or
	sql += join
	log.Println(sql)
}
