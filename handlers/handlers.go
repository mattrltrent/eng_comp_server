package handlers

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hn275/eng_comp_server/db"
	"github.com/hn275/eng_comp_server/lib"
	"gorm.io/gorm"
)

var (
	ErrMissingToken = errors.New("missing auth token")
)

type Handler struct {
	*gorm.DB
}

func New() *Handler {
	return &Handler{db.New()}
}

// TODO: fix this
func ParseJwt(t string) (*Token, error) {
	token, err := jwt.ParseWithClaims(t, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(lib.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GetUser(r *http.Request) (*Token, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, ErrMissingToken
	}
	return ParseJwt(authHeader)
}
