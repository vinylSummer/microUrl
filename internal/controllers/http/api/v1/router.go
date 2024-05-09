package http

import (
	"github.com/gorilla/mux"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/handlers"
	"github.com/vinylSummer/microUrl/internal/services"
	"github.com/vinylSummer/microUrl/pkg/logger"
)

func NewRouter(router *mux.Router, urlService services.URLService, logger logger.Interface) {
	handlers.NewURLRoutes(router, urlService, logger)
}
