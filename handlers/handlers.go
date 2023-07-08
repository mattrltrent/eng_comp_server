package handlers

import (
	"github.com/hn275/eng_comp_server/db"
	"gorm.io/gorm"
)

type Handler struct {
	*gorm.DB
}

func New() *Handler {
	return &Handler{db.New()}
}
