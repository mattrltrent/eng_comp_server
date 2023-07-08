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
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string `gorm:""`
}

func init() {
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func New() *gorm.DB {
	return db
}
