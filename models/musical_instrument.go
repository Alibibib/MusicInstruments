package models

type MusicalInstrument struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	PhotoURL    string  `json:"photo_url"`

	CategoryID uint     `json:"category_id"`
	Category   Category `json:"category"`
}
