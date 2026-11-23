package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victortran0904/urlshortener/internal/service"
)

const DefaultLen = 10

type GenPass interface {
	GeneratePassword(c *gin.Context)
}

type genPass struct {
	svc service.GenPass
}

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
