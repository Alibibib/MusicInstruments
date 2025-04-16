package handlers

import (
	"MusicInstruments/models"
	"MusicInstruments/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// MusicalInstrumentHandler структура для хендлеров музыкальных инструментов
type MusicalInstrumentHandler struct {
	Service *services.MusicalInstrumentService
}

// NewMusicalInstrumentHandler создает новый экземпляр хендлера для музыкальных инструментов
func NewMusicalInstrumentHandler() *MusicalInstrumentHandler {
	return &MusicalInstrumentHandler{
		Service: services.NewMusicalInstrumentService(),
	}
}

// CreateHandler создаёт новый музыкальный инструмент
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

// GetAllHandler возвращает все музыкальные инструменты
func (h *MusicalInstrumentHandler) GetAllHandler(c *gin.Context) {
	instruments, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить инструменты"})
		return
	}
	c.JSON(http.StatusOK, instruments)
}

// GetByIDHandler возвращает музыкальный инструмент по ID
func (h *MusicalInstrumentHandler) GetByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	instrument, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Инструмент не найден"})
		return
	}
	c.JSON(http.StatusOK, instrument)
}

// UpdateHandler обновляет музыкальный инструмент
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

// DeleteHandler удаляет музыкальный инструмент
func (h *MusicalInstrumentHandler) DeleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить инструмент"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Инструмент удалён"})
}
