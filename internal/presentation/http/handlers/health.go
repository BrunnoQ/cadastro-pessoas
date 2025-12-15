package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/BrunnoQ/cadastro-pessoas/internal/infrastructure/persistence/mongodb"
	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	dbConn *mongodb.Connection
}

// NewHealthHandler creates a new health handler
func NewHealthHandler(dbConn *mongodb.Connection) *HealthHandler {
	return &HealthHandler{
		dbConn: dbConn,
	}
}

// HealthResponse represents health check response
type HealthResponse struct {
	Status   string            `json:"status"`
	Services map[string]string `json:"services"`
}

// Check performs health check
func (h *HealthHandler) Check(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()

	response := HealthResponse{
		Status:   "healthy",
		Services: make(map[string]string),
	}

	// Check database connection
	if err := h.dbConn.HealthCheck(ctx); err != nil {
		response.Status = "unhealthy"
		response.Services["database"] = "unhealthy"
		c.JSON(http.StatusServiceUnavailable, response)
		return
	}

	response.Services["database"] = "healthy"
	c.JSON(http.StatusOK, response)
}
