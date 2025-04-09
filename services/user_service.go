package services

import (
	"MusicInstruments/database"
	"MusicInstruments/models"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func AddUser(user models.User) error {
	db := database.GetDB()

	var existingUser models.User
	result := db.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		log.Println("Ошибка при проверке существующего пользователя:", result.Error)
		return result.Error
	}
	if existingUser.ID != 0 {
		return fmt.Errorf("Пользователь с таким именем или email уже существует")
	}

	result = db.Create(&user)
	if result.Error != nil {
		log.Println("Ошибка при добавлении пользователя в базу данных:", result.Error)
		return result.Error
	}

	return nil
}

func GetAllUsers() ([]models.User, error) {
	db := database.GetDB()
	var users []models.User
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (models.User, error) {
	db := database.GetDB()
	var user models.User
	result := db.First(&user, id)
	return user, result.Error
}

func UpdateUser(id uint, updated models.User) error {
	db := database.GetDB()
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return err
	}

	user.Username = updated.Username
	user.Email = updated.Email
	user.Password = updated.Password

	return db.Save(&user).Error
}

func DeleteUser(id uint) error {
	db := database.GetDB()
	return db.Delete(&models.User{}, id).Error
}
