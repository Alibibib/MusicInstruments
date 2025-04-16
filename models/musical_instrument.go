package models

// Музыкальный инструмент
type MusicalInstrument struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Type        string  `json:"type"` // Тип инструмента (например, "струнный", "духовой", и т.д.)
	Description string  `json:"description"`
	Price       float64 `json:"price"`

	// Категория, к которой принадлежит инструмент
	CategoryID uint     `json:"category_id"` // Внешний ключ к категории
	Category   Category `json:"category"`    // Ссылка на саму категорию
}
