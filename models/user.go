package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role" gorm:"default:user"`
}
