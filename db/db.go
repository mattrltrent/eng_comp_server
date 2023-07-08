package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error

	TableUsers = "users"
	TablePosts = "posts"

	BannerJSON = "data.json"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string `gorm:""`
}

type Post struct {
	ID          uint   `gorm:"primaryKey" json:",omitempty"`
	UserID      uint   `json:"-"`
	Role        string `json:"role"`
	Course      string `json:"course"`
	Description string `json:"description"`
}

func init() {
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
}

func New() *gorm.DB {
	return db
}
