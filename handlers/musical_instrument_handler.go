package handlers

import (
	"MusicInstruments/database"
	"MusicInstruments/models"
	"MusicInstruments/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MusicalInstrumentHandler struct {
	Service *services.MusicalInstrumentService
}

func NewMusicalInstrumentHandler() *MusicalInstrumentHandler {
	return &MusicalInstrumentHandler{
		Service: services.NewMusicalInstrumentService(),
	}
}

func (h *MusicalInstrumentHandler) CreateHandler(c *gin.Context) {
	var instrument models.MusicalInstrument
	if err := c.ShouldBindJSON(&instrument); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(&instrument); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать инструмент"})
		return
	}
	c.JSON(http.StatusOK, instrument)
}

func (h *MusicalInstrumentHandler) GetAllHandler(c *gin.Context) {
	// Получаем параметр category_id из запроса
	categoryID, _ := strconv.Atoi(c.DefaultQuery("category_id", "0"))

	// Получаем инструменты с фильтрацией по категории
	instruments, err := h.Service.GetAll(categoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить инструменты"})
		return
	}
	c.JSON(http.StatusOK, instruments)
}

func (h *MusicalInstrumentHandler) GetByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	instrument, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Инструмент не найден"})
		return
	}
	c.JSON(http.StatusOK, instrument)
}

func (h *MusicalInstrumentHandler) UpdateHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var instrument models.MusicalInstrument
	if err := c.ShouldBindJSON(&instrument); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Update(uint(id), &instrument); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить инструмент"})
		return
	}
	c.JSON(http.StatusOK, instrument)
}

func (h *MusicalInstrumentHandler) DeleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить инструмент"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Инструмент удалён"})
}

func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	err := database.GetDB().Find(&categories).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения категорий"})
		return
	}
	c.JSON(http.StatusOK, categories)
}
