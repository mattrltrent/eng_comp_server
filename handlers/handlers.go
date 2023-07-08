package handlers

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hn275/eng_comp_server/db"
	"github.com/hn275/eng_comp_server/lib"
	"gorm.io/gorm"
)

type Handler struct {
	*gorm.DB
}

func New() *Handler {
	return &Handler{db.New()}
}

func ParseJwt(t string) (*Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
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
