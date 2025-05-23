package models

type Cart struct {
	ID     uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint       `json:"user_id"`
	User   User       `json:"user" gorm:"foreignKey:UserID"`
	Items  []CartItem `json:"items" gorm:"foreignKey:CartID"`
}
