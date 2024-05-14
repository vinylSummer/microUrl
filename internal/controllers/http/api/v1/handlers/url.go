package handlers

import (
	"github.com/gorilla/mux"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/handlers/urlHandler"
	"github.com/vinylSummer/microUrl/internal/services/v1"
	"github.com/vinylSummer/microUrl/pkg/logger"
)

func NewURLRoutes(router *mux.Router, urlService v1.URLService, logger logger.Interface) {
	urlhandler.NewCreateShortURLRoute(router, urlService, logger)
	urlhandler.NewGetLongURLRoute(router, urlService, logger)
}
