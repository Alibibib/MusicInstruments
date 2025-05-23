package handlers

import (
	"MusicInstruments/database"
	"MusicInstruments/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET /me
func MeHandler(c *gin.Context) {
	// Получаем userID из контекста
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID не найден в контексте"})
		return
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "userID неверного типа"})
		return
	}

	fmt.Println("В MeHandler userID:", userID)

	var user models.User
	err := database.GetDB().Preload("Role").First(&user, userID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		return
	}
	fmt.Println("MeHandler вызван с userID:", userID)

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role.Name,
	})
}
