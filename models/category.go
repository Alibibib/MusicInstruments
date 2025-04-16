package models

// Category модель для категории музыкальных инструментов
type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
}
