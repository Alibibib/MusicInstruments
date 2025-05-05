package main

import (
	"github.com/gin-gonic/gin"
	"instrument-service/handler"
	"instrument-service/middleware"
)

func main() {
	router := gin.New()

	router.Use(middleware.LoggingMiddleware())

	router.GET("/instruments", handler.GetInstruments)
	router.POST("/instruments", handler.CreateInstrument)

	router.Run(":8082")
}
