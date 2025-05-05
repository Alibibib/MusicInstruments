package main

import (
	"github.com/gin-gonic/gin"
	"user-service/handler"
	"user-service/middleware"
)

func main() {
	router := gin.New()

	router.Use(middleware.LoggingMiddleware())

	router.GET("/users", handler.GetUsers)
	router.GET("/users/instruments", handler.GetInstruments)

	router.Run(":8081")
}
