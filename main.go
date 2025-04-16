package main

import (
	"MusicInstruments/database"
	"MusicInstruments/handlers"
	"MusicInstruments/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := database.InitDB()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	r := gin.Default()

	// Auth routes
	r.POST("/register", handlers.RegisterHandler)
	r.POST("/login", handlers.LoginHandler)

	// Защищённые
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/users", handlers.GetAllUsersHandler)
		auth.GET("/users/:id", handlers.GetUserByIDHandler)
		auth.PUT("/users/:id", handlers.UpdateUserHandler)
		auth.DELETE("/users/:id", handlers.DeleteUserHandler)

		// Эндпоинты для музыкальных инструментов
		instrumentHandler := handlers.NewMusicalInstrumentHandler()
		auth.GET("/instruments", instrumentHandler.GetAllHandler)
		auth.GET("/instruments/:id", instrumentHandler.GetByIDHandler)
		auth.POST("/instruments", instrumentHandler.CreateHandler)
		auth.PUT("/instruments/:id", instrumentHandler.UpdateHandler)
		auth.DELETE("/instruments/:id", instrumentHandler.DeleteHandler)

		// Эндпоинты для категорий
		categoryHandler := handlers.NewCategoryHandler()
		auth.GET("/categories", categoryHandler.GetAllHandler)
		auth.GET("/categories/:id", categoryHandler.GetByIDHandler)
		auth.POST("/categories", categoryHandler.CreateHandler)
		auth.PUT("/categories/:id", categoryHandler.UpdateHandler)
		auth.DELETE("/categories/:id", categoryHandler.DeleteHandler)
	}

	log.Println("Сервер запущен на http://localhost:8080")
	r.Run(":8080")
}
