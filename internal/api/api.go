package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/victortran0904/urlshortener/internal/handler"
	"github.com/victortran0904/urlshortener/internal/service"
)

type Engine interface {
	Start() error
}

type engine struct {
	app *gin.Engine
}

func NewEngine() Engine {
	e := &engine{app: gin.Default()}
	e.initRoutes()
	return e
}

func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%d", 8080))
}

func (e *engine) initRoutes() {
	genPassSvc := service.NewGenPass()
	genPassHandler := handler.NewGenPass(genPassSvc)

	e.app.GET("/genpass", genPassHandler.GeneratePassword)

}
