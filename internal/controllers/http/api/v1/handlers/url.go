package handlers

import (
	"github.com/gorilla/mux"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/handlers/urlHandler"
	"github.com/vinylSummer/microUrl/internal/services/v1"
)

func NewURLRoutes(router *mux.Router, urlService v1.URLService) {
	urlhandler.NewCreateShortURLRoute(router, urlService)
	urlhandler.NewGetLongURLRoute(router, urlService)
}
