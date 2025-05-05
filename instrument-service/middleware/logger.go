package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)

		c.Next()

		duration := time.Since(start)
		log.Printf("[%s] [RequestID: %s] %s %s - %d - Duration: %.3fms",
			time.Now().Format(time.RFC3339),
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			float64(duration.Microseconds())/1000.0,
		)
	}
}
