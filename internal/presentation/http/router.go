package http

import (
	"net/http"

	"github.com/BrunnoQ/cadastro-pessoas/internal/presentation/http/handlers"
	"github.com/gin-gonic/gin"
)

// Router manages HTTP routes and middleware
type Router struct {
	engine        *gin.Engine
	healthHandler *handlers.HealthHandler
	personHandler *handlers.PersonHandler
}

// NewRouter creates a new HTTP router
func NewRouter(healthHandler *handlers.HealthHandler, personHandler *handlers.PersonHandler) *Router {
	// Set Gin mode based on environment
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	return &Router{
		engine:        engine,
		healthHandler: healthHandler,
		personHandler: personHandler,
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
		v1.GET("/health", r.healthHandler.Check)

		// Person endpoints
		v1.GET("/persons", r.personHandler.List)
		v1.POST("/persons", r.personHandler.Create)
		v1.GET("/persons/:id", r.personHandler.GetByID)
		v1.PUT("/persons/:id", r.personHandler.Update)
	}
}

// ServeHTTP implements http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.engine.ServeHTTP(w, req)
}
