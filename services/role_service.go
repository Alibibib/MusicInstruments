// services/role_service.go
package services

import (
	"MusicInstruments/database"
	"MusicInstruments/models"
)

func CreateRole(role models.Role) error {
	return database.GetDB().Create(&role).Error
}

func GetAllRoles() ([]models.Role, error) {
	var roles []models.Role
	err := database.GetDB().Find(&roles).Error
	return roles, err
}

func GetRoleByID(id uint) (models.Role, error) {
	var role models.Role
	err := database.GetDB().First(&role, id).Error
	return role, err
}

func UpdateRole(id uint, updated models.Role) error {
	var role models.Role
	if err := database.GetDB().First(&role, id).Error; err != nil {
		return err
	}
	role.Name = updated.Name
	return database.GetDB().Save(&role).Error
}

func DeleteRole(id uint) error {
	return database.GetDB().Delete(&models.Role{}, id).Error
}
