package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// ValidationMiddleware validates request body against struct tags
func ValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validation is automatic with gin's binding
		// This middleware can be extended for custom validation logic
		c.Next()
	}
}

// BindJSON binds JSON body and validates
func BindJSON(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindJSON(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "VALIDATION_ERROR",
			"message": "Invalid request body",
			"details": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return err
	}
	return nil
}

// BindQuery binds query parameters and validates
func BindQuery(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindQuery(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "VALIDATION_ERROR",
			"message": "Invalid query parameters",
			"details": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return err
	}
	return nil
}

// BindURI binds URI parameters and validates
func BindURI(c *gin.Context, obj interface{}) error {
	if err := c.ShouldBindUri(obj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "VALIDATION_ERROR",
			"message": "Invalid URI parameters",
			"details": map[string]interface{}{
				"error": err.Error(),
			},
		})
		return err
	}
	return nil
}

// ValidateContentType validates request content type
func ValidateContentType(expectedType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodDelete {
			contentType := c.ContentType()
			if contentType != expectedType {
				c.AbortWithStatusJSON(http.StatusUnsupportedMediaType, gin.H{
					"code":    "INVALID_CONTENT_TYPE",
					"message": "Content-Type must be " + expectedType,
				})
				return
			}
		}
		c.Next()
	}
}

// ValidateJSONContentType validates JSON content type
func ValidateJSONContentType() gin.HandlerFunc {
	return ValidateContentType(binding.MIMEJSON)
}
