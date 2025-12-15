package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// CompressionMiddleware adds gzip compression to responses
// Compresses responses larger than 1KB
func CompressionMiddleware() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions([]string{".png", ".jpg", ".jpeg", ".gif", ".pdf"}))
}
