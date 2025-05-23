package models

type CartItem struct {
	ID uint `json:"id" gorm:"primaryKey;autoIncrement"`

	CartID uint `json:"cart_id"`
	Cart   Cart `json:"cart" gorm:"foreignKey:CartID"`

	MusicalInstrumentID uint              `json:"musical_instrument_id"`
	MusicalInstrument   MusicalInstrument `json:"musical_instrument" gorm:"foreignKey:MusicalInstrumentID"`

	Quantity uint `json:"quantity"`
}
