package services

import (
	"MusicInstruments/database"
	"MusicInstruments/models"
	"gorm.io/gorm"
	"log"
)

// CategoryService структура сервиса для категорий
type CategoryService struct {
	DB *gorm.DB
}

// NewCategoryService создает новый экземпляр сервиса для категорий
func NewCategoryService() *CategoryService {
	return &CategoryService{
		DB: database.GetDB(),
	}
}

// Create создает новую категорию
func (service *CategoryService) Create(category *models.Category) error {
	if err := service.DB.Create(category).Error; err != nil {
		log.Println("Ошибка при создании категории:", err)
		return err
	}
	return nil
}

// GetAll возвращает все категории
func (service *CategoryService) GetAll() ([]models.Category, error) {
	var categories []models.Category
	if err := service.DB.Find(&categories).Error; err != nil {
		log.Println("Ошибка при получении категорий:", err)
		return nil, err
	}
	return categories, nil
}

// GetByID возвращает категорию по ID
func (service *CategoryService) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := service.DB.First(&category, id).Error; err != nil {
		log.Println("Ошибка при получении категории:", err)
		return nil, err
	}
	return &category, nil
}

// Update обновляет категорию
func (service *CategoryService) Update(id uint, category *models.Category) error {
	if err := service.DB.Model(&models.Category{}).Where("id = ?", id).Updates(category).Error; err != nil {
		log.Println("Ошибка при обновлении категории:", err)
		return err
	}
	return nil
}

// Delete удаляет категорию
func (service *CategoryService) Delete(id uint) error {
	if err := service.DB.Delete(&models.Category{}, id).Error; err != nil {
		log.Println("Ошибка при удалении категории:", err)
		return err
	}
	return nil
}
