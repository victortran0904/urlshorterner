package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victortran0904/urlshortener/internal/service"
)

// DefaultLen is the number of characters requested by the password handler.
const DefaultLen = 10

// GenPass handles requests for generated passwords.
type GenPass interface {
	// GeneratePassword writes a 200 JSON response with a "password" field.
	// It requests DefaultLen characters and writes a generic 500 JSON error on failure.
	GeneratePassword(c *gin.Context)
}

type genPass struct {
	svc service.GenPass
}

// NewGenPass creates a handler that obtains passwords from svc.
func NewGenPass(svc service.GenPass) GenPass {
	return &genPass{svc: svc}
}

func (g *genPass) GeneratePassword(c *gin.Context) {

	pass, err := g.svc.GeneratePassword(DefaultLen)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"password": pass})

}
