package services

import (
	"MusicInstruments/database"
	"MusicInstruments/models"
	"gorm.io/gorm"
	"log"
)

// MusicalInstrumentService структура сервиса для музыкальных инструментов
type MusicalInstrumentService struct {
	DB *gorm.DB
}

// NewMusicalInstrumentService создает новый экземпляр сервиса для музыкальных инструментов
func NewMusicalInstrumentService() *MusicalInstrumentService {
	return &MusicalInstrumentService{
		DB: database.GetDB(),
	}
}

// Create музыкальный инструмент
func (service *MusicalInstrumentService) Create(instrument *models.MusicalInstrument) error {
	if err := service.DB.Create(instrument).Error; err != nil {
		log.Println("Ошибка при создании инструмента:", err)
		return err
	}
	return nil
}

// GetAll возвращает все музыкальные инструменты
func (service *MusicalInstrumentService) GetAll() ([]models.MusicalInstrument, error) {
	var instruments []models.MusicalInstrument
	if err := service.DB.Preload("Category").Find(&instruments).Error; err != nil {
		log.Println("Ошибка при получении инструментов:", err)
		return nil, err
	}
	return instruments, nil
}

// GetByID возвращает инструмент по ID
func (service *MusicalInstrumentService) GetByID(id uint) (*models.MusicalInstrument, error) {
	var instrument models.MusicalInstrument
	if err := service.DB.Preload("Category").First(&instrument, id).Error; err != nil {
		log.Println("Ошибка при получении инструмента:", err)
		return nil, err
	}
	return &instrument, nil
}

// Update обновляет музыкальный инструмент
func (service *MusicalInstrumentService) Update(id uint, instrument *models.MusicalInstrument) error {
	if err := service.DB.Model(&models.MusicalInstrument{}).Where("id = ?", id).Updates(instrument).Error; err != nil {
		log.Println("Ошибка при обновлении инструмента:", err)
		return err
	}
	return nil
}

// Delete удаляет музыкальный инструмент
func (service *MusicalInstrumentService) Delete(id uint) error {
	if err := service.DB.Delete(&models.MusicalInstrument{}, id).Error; err != nil {
		log.Println("Ошибка при удалении инструмента:", err)
		return err
	}
	return nil
}
