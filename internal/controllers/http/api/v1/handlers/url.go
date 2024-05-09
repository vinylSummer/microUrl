package handlers

import (
	"github.com/gorilla/mux"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/handlers/urlHandler"
	"github.com/vinylSummer/microUrl/internal/services"
	"github.com/vinylSummer/microUrl/pkg/logger"
)

func NewURLRoutes(router *mux.Router, urlService services.URLService, logger logger.Interface) {
	urlhandler.NewCreateShortURLRoute(router, urlService, logger)
	urlhandler.NewGetLongURLRoute(router, urlService, logger)
}
