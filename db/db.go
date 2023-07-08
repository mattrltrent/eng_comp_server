package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error

	TableUsers = "users"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique"`
	Password string `gorm:""`
}

type CourseEntry struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint `gorm:"references"`
	CourseID uint `gorm:"references"`
	Role     string
}

func init() {
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&CourseEntry{})
}

func New() *gorm.DB {
	return db
}
