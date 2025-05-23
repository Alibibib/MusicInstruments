package handlers

import (
	"MusicInstruments/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	cartService services.CartService
	// можно добавить в память временную гостевую корзину, если нужно
	// guestCarts map[string][]uint // map sessionID -> list of instrument IDs
}

func NewCartHandler() *CartHandler {
	return &CartHandler{
		cartService: services.NewCartService(),
		// guestCarts: make(map[string][]uint),
	}
}

// POST /add-to-cart/:id
func (h *CartHandler) AddToCartHandler(c *gin.Context) {
	instrumentIDStr := c.Param("id")
	instrumentID64, err := strconv.ParseUint(instrumentIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid instrument ID"})
		return
	}
	instrumentID := uint(instrumentID64)

	userIDValue, exists := c.Get("userID")
	if exists {
		// Пользователь авторизован — добавляем в корзину по userID
		userID, ok := userIDValue.(uint)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
			return
		}

		err = h.cartService.AddToCart(userID, instrumentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Товар добавлен в корзину"})
		return
	}

	// Гость — например, можно сохранить временно в сессии или куках
	// Здесь просто выводим ID товара в консоль и возвращаем сообщение
	fmt.Println("Гость добавил товар с ID:", instrumentID)

	// TODO: Реализовать логику гостевой корзины (например, хранение в сессии или Redis)

	c.JSON(http.StatusOK, gin.H{"message": "Товар добавлен в корзину (гостевой режим)"})
}

// GET /cart — можно тоже сделать поддержку гостевой корзины
func (h *CartHandler) ViewCartHandler(c *gin.Context) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		// Гость — пока просто вернуть пустую корзину или сообщение
		c.HTML(http.StatusOK, "cart.html", gin.H{
			"Items": []interface{}{}, // пустой список
			"Guest": true,
		})
		return
	}
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	items, err := h.cartService.GetCartItems(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart items"})
		return
	}

	c.HTML(http.StatusOK, "cart.html", gin.H{
		"Items": items,
		"Guest": false,
	})
}
