package handlers

import (
	"MusicInstruments/models"
	"MusicInstruments/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CategoryHandler структура для хендлеров категорий
type CategoryHandler struct {
	Service *services.CategoryService
}

// NewCategoryHandler создает новый экземпляр хендлера для категорий
func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		Service: services.NewCategoryService(),
	}
}

// CreateHandler создаёт новую категорию
func (h *CategoryHandler) CreateHandler(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Create(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать категорию"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// GetAllHandler возвращает все категории
func (h *CategoryHandler) GetAllHandler(c *gin.Context) {
	categories, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить категории"})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GetByIDHandler возвращает категорию по ID
func (h *CategoryHandler) GetByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найдена"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// UpdateHandler обновляет категорию
func (h *CategoryHandler) UpdateHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Service.Update(uint(id), &category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить категорию"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// DeleteHandler удаляет категорию
func (h *CategoryHandler) DeleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить категорию"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Категория удалена"})
}
