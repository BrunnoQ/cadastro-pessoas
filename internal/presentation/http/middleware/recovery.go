package middleware

import (
	"net/http"

	"github.com/BrunnoQ/cadastro-pessoas/internal/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RecoveryMiddleware recovers from panics and returns 500 error
func RecoveryMiddleware(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic
				log.Error("Panic recovered",
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
					zap.String("method", c.Request.Method),
				)

				// Return 500 error
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    "INTERNAL_ERROR",
					"message": "An internal error occurred",
				})
			}
		}()

		c.Next()
	}
}
