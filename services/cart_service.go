package services

import (
	"MusicInstruments/database"
	"MusicInstruments/models"
	"errors"
	"gorm.io/gorm"
)

type CartService interface {
	AddToCart(userID uint, instrumentID uint) error
	GetCartItems(userID uint) ([]models.CartItem, error)
}

// Реализация CartService
type cartService struct{}

func NewCartService() CartService {
	return &cartService{}
}

func (s *cartService) AddToCart(userID uint, instrumentID uint) error {
	db := database.GetDB()

	// Получаем корзину пользователя
	var cart models.Cart
	err := db.Where("user_id = ?", userID).Preload("Items").First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Если корзина не найдена, создаём новую
			cart = models.Cart{
				UserID: userID,
			}
			if err := db.Create(&cart).Error; err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Проверяем, есть ли уже такой товар в корзине
	for _, item := range cart.Items {
		if item.MusicalInstrumentID == instrumentID {
			item.Quantity++
			return db.Save(&item).Error
		}
	}

	// Если товара нет, добавляем новый элемент
	cartItem := models.CartItem{
		CartID:              cart.ID,
		MusicalInstrumentID: instrumentID,
		Quantity:            1,
	}

	return db.Create(&cartItem).Error
}

func (s *cartService) GetCartItems(userID uint) ([]models.CartItem, error) {
	db := database.GetDB()

	// Получаем корзину пользователя с предзагрузкой элементов и инструментов
	var cart models.Cart
	err := db.Where("user_id = ?", userID).Preload("Items.MusicalInstrument").First(&cart).Error
	if err != nil {
		return nil, err
	}

	return cart.Items, nil
}
