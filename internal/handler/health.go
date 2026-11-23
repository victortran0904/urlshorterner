package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/victortran0904/urlshortener/internal/service"
)

type Health interface {
	CheckHealth(c *gin.Context)
}

type health struct {
	svc service.Health
}

func NewHealth(svc service.Health) Health {
	return &health{}
}

func (h *health) CheckHealth(c *gin.Context){
	message, err := h.svc.CheckHealth()

}

