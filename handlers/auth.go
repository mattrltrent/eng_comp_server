package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/hn275/eng_comp_server/db"
	"golang.org/x/crypto/bcrypt"
)

type AuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (d *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get body
	var auth AuthData
	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		fmt.Fprint(os.Stderr, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// hash password
	cost := 10
	hashed, err := bcrypt.GenerateFromPassword([]byte(auth.Password), cost)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	auth.Password = string(hashed)

	// save to db
	if err := d.Table(db.TableUsers).Create(&auth).Error; err != nil {
		fmt.Fprint(os.Stderr, err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "duplicate email"})
		return
	}

	w.WriteHeader(http.StatusCreated)
}
