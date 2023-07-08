package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hn275/eng_comp_server/db"
	"github.com/hn275/eng_comp_server/lib"
	"golang.org/x/crypto/bcrypt"
)

type AuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	AuthData
	jwt.RegisteredClaims
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
	hashed, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
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

func (d *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
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

	// pull user
	var user db.User
	result := d.Table(db.TableUsers).Where("email = ?", auth.Email).First(&user)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		lib.Err(w, "user not found")
		return
	}

	// verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		lib.Err(w, err.Error())
		return
	}

	// sign jwt
	claims := Token{
		auth,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "EduBuddy",
			Subject:   user.Email,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(lib.JwtSecret))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		lib.Err(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	lib.JSON(w, map[string]string{"token": token})
}
