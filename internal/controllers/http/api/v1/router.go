package http

import (
	"github.com/gorilla/mux"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/handlers"
	"github.com/vinylSummer/microUrl/internal/services/v1"
)

func NewRouter(router *mux.Router, urlService v1.URLService) {
	handlers.NewURLRoutes(router, urlService)
}
