package services

import (
	"MusicInstruments/database"
	"MusicInstruments/models"
	"gorm.io/gorm"
	"log"
)

type MusicalInstrumentService struct {
	DB *gorm.DB
}

func NewMusicalInstrumentService() *MusicalInstrumentService {
	return &MusicalInstrumentService{
		DB: database.GetDB(),
	}
}

func (service *MusicalInstrumentService) Create(instrument *models.MusicalInstrument) error {
	if err := service.DB.Create(instrument).Error; err != nil {
		log.Println("Ошибка при создании инструмента:", err)
		return err
	}
	return nil
}

// Метод GetAll без пагинации с фильтрацией по категории
func (service *MusicalInstrumentService) GetAll(categoryID int) ([]models.MusicalInstrument, error) {
	var instruments []models.MusicalInstrument

	// Запрос с фильтрацией по категории
	query := service.DB.Preload("Category")
	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	if err := query.Find(&instruments).Error; err != nil {
		return nil, err
	}
	return instruments, nil
}

func (service *MusicalInstrumentService) GetByID(id uint) (*models.MusicalInstrument, error) {
	var instrument models.MusicalInstrument
	if err := service.DB.Preload("Category").First(&instrument, id).Error; err != nil {
		log.Println("Ошибка при получении инструмента:", err)
		return nil, err
	}
	return &instrument, nil
}

func (service *MusicalInstrumentService) Update(id uint, instrument *models.MusicalInstrument) error {
	if err := service.DB.Model(&models.MusicalInstrument{}).Where("id = ?", id).Updates(instrument).Error; err != nil {
		log.Println("Ошибка при обновлении инструмента:", err)
		return err
	}
	return nil
}

func (service *MusicalInstrumentService) Delete(id uint) error {
	if err := service.DB.Delete(&models.MusicalInstrument{}, id).Error; err != nil {
		log.Println("Ошибка при удалении инструмента:", err)
		return err
	}
	return nil
}
