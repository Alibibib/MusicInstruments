//package middleware
//
//import (
//	"MusicInstruments/services"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strings"
//)
//
//func AuthMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		authHeader := c.GetHeader("Authorization")
//		fmt.Println("Authorization header:", authHeader) // Лог заголовка
//
//		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Нет токена"})
//			c.Abort()
//			return
//		}
//
//		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
//		claims, err := services.ValidateToken(tokenStr)
//		if err != nil {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
//			c.Abort()
//			return
//		}
//		fmt.Println("Пользователь прошёл аутентификацию, userID =", claims.UserID)
//
//		c.Set("userID", uint(claims.UserID)) // если claims.UserID — int, конвертируй в uint
//		c.Next()
//	}
//}

package middleware

import (
	"MusicInstruments/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из cookie
		tokenStr, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Нет токена"})
			c.Abort()
			return
		}

		// Валидация токена
		claims, err := services.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
			c.Abort()
			return
		}

		// Сохраняем userID в контекст
		c.Set("userID", uint(claims.UserID))
		c.Next()
	}
}
