package http

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/vinylSummer/microUrl/docs"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/handlers"
	v1 "github.com/vinylSummer/microUrl/internal/services/v1"
)

// @title microURL API
// @version 1.0
// @description Shorten your URLs

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func NewRouter(router *mux.Router, urlService v1.URLService) {
	handlers.NewURLRoutes(router, urlService)
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
