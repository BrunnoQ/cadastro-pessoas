package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router manages HTTP routes and middleware
type Router struct {
	engine *gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter() *Router {
	// Set Gin mode based on environment
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	return &Router{
		engine: engine,
	}
}

// Engine returns the underlying Gin engine
func (r *Router) Engine() *gin.Engine {
	return r.engine
}

// SetupRoutes configures all application routes
func (r *Router) SetupRoutes() {
	// API v1 group
	v1 := r.engine.Group("/api/v1")
	{
		// Health check endpoint
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "healthy",
			})
		})
	}
}

// ServeHTTP implements http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.engine.ServeHTTP(w, req)
}
