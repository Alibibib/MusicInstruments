package handlers

import (
	"MusicInstruments/models"
	"MusicInstruments/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddUserHandler(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}

	if err := services.AddUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetAllUsersHandler(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный параметр page"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный параметр limit"})
		return
	}

	users, err := services.GetAllUsers(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUserByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	user, err := services.GetUserByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
		return
	}

	if err := services.UpdateUser(uint(id), user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь успешно обновлён"})
}

func DeleteUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	if err := services.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления"})
		return
	}

	c.Status(http.StatusNoContent)
}
