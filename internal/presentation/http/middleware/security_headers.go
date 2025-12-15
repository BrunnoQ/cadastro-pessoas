package middleware

import (
	"github.com/gin-gonic/gin"
)

// SecurityHeadersMiddleware adds security headers to all responses
func SecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prevent clickjacking attacks
		c.Header("X-Frame-Options", "DENY")

		// Prevent MIME type sniffing
		c.Header("X-Content-Type-Options", "nosniff")

		// Enable XSS protection (for older browsers)
		c.Header("X-XSS-Protection", "1; mode=block")

		// Content Security Policy
		// Adjust this based on your application needs
		c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; font-src 'self'; connect-src 'self'; frame-ancestors 'none'")

		// Referrer Policy
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		// Permissions Policy (formerly Feature Policy)
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		// HSTS (HTTP Strict Transport Security)
		// Only enable in production with HTTPS
		// Uncomment the line below when using HTTPS in production
		// c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		c.Next()
	}
}
