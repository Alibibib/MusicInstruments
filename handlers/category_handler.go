package handlers

import (
	"MusicInstruments/models"
	"MusicInstruments/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	Service *services.CategoryService
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{
		Service: services.NewCategoryService(),
	}
}

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

func (h *CategoryHandler) GetAllHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	categories, err := h.Service.GetAll(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить категории"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
		"page":       page,
		"limit":      limit,
	})
}

func (h *CategoryHandler) GetByIDHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Категория не найдена"})
		return
	}
	c.JSON(http.StatusOK, category)
}

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

func (h *CategoryHandler) DeleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить категорию"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Категория удалена"})
}
