package main

import (
	"MusicInstruments/database"
	"MusicInstruments/handlers"
	"MusicInstruments/middleware"
	"MusicInstruments/models"
	"MusicInstruments/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	// Инициализация базы данных
	err := database.InitDB()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"}, // Добавь нужные фронтенд адреса
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Подключение статики и шаблонов
	r.Static("/static", "./static")    // Статические файлы (CSS, JS)
	r.LoadHTMLGlob("templates/*.html") // HTML-шаблоны

	// Отображение страниц
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.GET("/", func(c *gin.Context) {
		role := ""
		isLoggedIn := false

		if userID, exists := c.Get("userID"); exists {
			var user models.User
			// Получаем пользователя из БД по userID
			database.GetDB().Preload("Role").First(&user, userID)
			role = user.Role.Name
			isLoggedIn = true
		}

		instruments, _ := services.NewMusicalInstrumentService().GetAll(0)

		// Передаём данные в шаблон
		c.HTML(http.StatusOK, "index.html", gin.H{
			"IsLoggedIn":  isLoggedIn,
			"Role":        role,
			"Instruments": instruments,
		})
	})

	// Обработка регистрации и логина
	r.POST("/register", handlers.RegisterHandler)
	r.POST("/login", handlers.LoginHandler)
	r.GET("/categories", handlers.GetAllCategories)

	// Защищённые маршруты
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/users", handlers.GetAllUsersHandler)
		auth.GET("/users/:id", handlers.GetUserByIDHandler)
		auth.PUT("/users/:id", handlers.UpdateUserHandler)
		auth.DELETE("/users/:id", handlers.DeleteUserHandler)

		instrumentHandler := handlers.NewMusicalInstrumentHandler()
		auth.GET("/instruments", instrumentHandler.GetAllHandler)
		auth.GET("/instruments/:id", instrumentHandler.GetByIDHandler)
		auth.POST("/instruments", instrumentHandler.CreateHandler)
		auth.PUT("/instruments/:id", instrumentHandler.UpdateHandler)
		auth.DELETE("/instruments/:id", instrumentHandler.DeleteHandler)

		categoryHandler := handlers.NewCategoryHandler()
		auth.GET("/categories/:id", categoryHandler.GetByIDHandler)
		auth.POST("/categories", categoryHandler.CreateHandler)
		auth.PUT("/categories/:id", categoryHandler.UpdateHandler)
		auth.DELETE("/categories/:id", categoryHandler.DeleteHandler)

		auth.POST("/roles", handlers.CreateRoleHandler)
		auth.GET("/roles", handlers.GetAllRolesHandler)
		auth.GET("/roles/:id", handlers.GetRoleByIDHandler)
		auth.PUT("/roles/:id", handlers.UpdateRoleHandler)
		auth.DELETE("/roles/:id", handlers.DeleteRoleHandler)

		auth.GET("/me", handlers.MeHandler)

		cartHandler := handlers.NewCartHandler()
		auth.POST("/add-to-cart/:id", cartHandler.AddToCartHandler)
		auth.GET("/cart", cartHandler.ViewCartHandler)
	}

	log.Println("Сервер запущен на http://localhost:8080")
	log.Println("Сервер запущен на http://localhost:8080/register")
	err = r.Run()
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
