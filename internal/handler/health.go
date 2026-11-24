// Package handler implements HTTP handlers for health checks and password generation.
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victortran0904/urlshortener/internal/service"
)

// HealthCheck handles requests for the service's health status and identity.
type HealthCheck interface {
	// CheckHealth writes the health response as JSON with status 200.
	// It writes a generic JSON error with status 500 if the service returns an error.
	CheckHealth(c *gin.Context)
}

type healthCheck struct {
	svc service.HealthCheck
}

// NewHealthCheck creates a handler that obtains health responses from svc.
func NewHealthCheck(svc service.HealthCheck) HealthCheck {
	return &healthCheck{svc: svc}
}

func (h *healthCheck) CheckHealth(c *gin.Context) {
	body, err := h.svc.CheckHealth()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, body)

}
