package schemas

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Role     string
	Password string
}
