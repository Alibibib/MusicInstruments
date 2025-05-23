// handlers/role_handler.go
package handlers

import (
	"MusicInstruments/models"
	"MusicInstruments/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateRoleHandler(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}
	if err := services.CreateRole(role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания роли"})
		return
	}
	c.JSON(http.StatusCreated, role)
}

func GetAllRolesHandler(c *gin.Context) {
	roles, err := services.GetAllRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения ролей"})
		return
	}
	c.JSON(http.StatusOK, roles)
}

func GetRoleByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	role, err := services.GetRoleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Роль не найдена"})
		return
	}
	c.JSON(http.StatusOK, role)
}

func UpdateRoleHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}
	if err := services.UpdateRole(uint(id), role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления роли"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Роль обновлена"})
}

func DeleteRoleHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteRole(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления роли"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Роль удалена"})
}
