package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

var instruments = []map[string]string{
	{"id": "1", "name": "Guitar"},
	{"id": "2", "name": "Piano"},
}

func GetInstruments(c *gin.Context) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&[]map[string]string{}).
		Get("http://localhost:8081/users")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch users"})
		return
	}

	users := resp.Result().(*[]map[string]string)

	c.JSON(http.StatusOK, gin.H{
		"users":       users,
		"instruments": instruments,
	})
}

func CreateInstrument(c *gin.Context) {
	var newInstrument map[string]string

	// Преобразуем тело запроса в карту
	if err := c.ShouldBindJSON(&newInstrument); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}

	newInstrument["id"] = string(len(instruments) + 1)

	instruments = append(instruments, newInstrument)

	c.JSON(http.StatusCreated, newInstrument)
}
