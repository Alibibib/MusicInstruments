package services

import (
	"MusicInstruments/database"
	"MusicInstruments/models"
	"gorm.io/gorm"
	"log"
)

type CategoryService struct {
	DB *gorm.DB
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		DB: database.GetDB(),
	}
}

func (service *CategoryService) Create(category *models.Category) error {
	if err := service.DB.Create(category).Error; err != nil {
		log.Println("Ошибка при создании категории:", err)
		return err
	}
	return nil
}

// Измененный метод GetAll с пагинацией
func (service *CategoryService) GetAll(limit int, offset int) ([]models.Category, error) {
	var categories []models.Category
	if err := service.DB.Limit(limit).Offset(offset).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (service *CategoryService) GetByID(id uint) (*models.Category, error) {
	var category models.Category
	if err := service.DB.First(&category, id).Error; err != nil {
		log.Println("Ошибка при получении категории:", err)
		return nil, err
	}
	return &category, nil
}

func (service *CategoryService) Update(id uint, category *models.Category) error {
	if err := service.DB.Model(&models.Category{}).Where("id = ?", id).Updates(category).Error; err != nil {
		log.Println("Ошибка при обновлении категории:", err)
		return err
	}
	return nil
}

func (service *CategoryService) Delete(id uint) error {
	if err := service.DB.Delete(&models.Category{}, id).Error; err != nil {
		log.Println("Ошибка при удалении категории:", err)
		return err
	}
	return nil
}
