package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func GetUsers(c *gin.Context) {
	users := []map[string]string{
		{"id": "1", "name": "John Doe"},
		{"id": "2", "name": "Jane Smith"},
	}

	c.JSON(http.StatusOK, users)
}

func GetInstruments(c *gin.Context) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&[]map[string]string{}).
		Get("http://localhost:8082/instruments")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch instruments"})
		return
	}

	instruments := resp.Result().(*[]map[string]string)

	c.JSON(http.StatusOK, gin.H{
		"users": []map[string]string{
			{"id": "1", "name": "John Doe"},
			{"id": "2", "name": "Jane Smith"},
		},
		"instruments": instruments,
	})
}
