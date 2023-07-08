package handlers

import (
	"log"
	"net/http"
)

func (d *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	courses := r.URL.Query().Get("course")
	log.Println(courses)
	//
	// token, err := GetUser(r)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	lib.Err(w, err.Error())
	// 	return
	// }
	//
	// _, err = d.GetUserInfo(token)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	lib.Err(w, err.Error())
	// 	return
	// }
}
