// Package api configures and serves the health and password HTTP endpoints.
package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victortran0904/urlshortener/internal/handler"
	"github.com/victortran0904/urlshortener/internal/service"
)

// Engine serves the API over HTTP.
type Engine interface {
	// Start serves HTTP on the configured application port and returns any server error.
	Start() error
	// ServeHTTP dispatches a request to the registered routes.
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type engine struct {
	app *gin.Engine
	cfg *Config
}

// NewEngine creates an engine with request logging, panic recovery, and GET routes
// for /health and /genpass.
func NewEngine(cfg *Config) Engine {
	e := &engine{app: gin.Default(), cfg: cfg}
	e.initRoutes()
	return e
}

func (e *engine) Start() error {
	return e.app.Run(fmt.Sprintf(":%d", e.cfg.AppPort))
}

func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.app.ServeHTTP(w, req)
}

func (e *engine) initRoutes() {
	genPassSvc := service.NewGenPass()
	genPassHandler := handler.NewGenPass(genPassSvc)
	e.app.GET("/genpass", genPassHandler.GeneratePassword)

	healthCheckSvc := service.NewHealthCheck(e.cfg.ServiceName, e.cfg.InstanceID)
	healthCheckHandler := handler.NewHealthCheck(healthCheckSvc)
	e.app.GET("/health", healthCheckHandler.CheckHealth)
}
