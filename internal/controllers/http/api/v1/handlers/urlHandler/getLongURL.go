package urlhandler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/handlers/urlHandler/dto"
	"github.com/vinylSummer/microUrl/internal/services"
	"github.com/vinylSummer/microUrl/pkg/logger"
	"net/http"
)

type GetLongURLRoute struct {
	urlService services.URLService
	logger     logger.Interface
}

func NewGetLongURLRoute(router *mux.Router, urlService services.URLService, logger logger.Interface) {
	route := &GetLongURLRoute{
		urlService: urlService,
		logger:     logger,
	}

	router.HandleFunc("/{path:[a-zA-Z0-9]+}", route.getLongURL).Methods("GET")
}

func (route *GetLongURLRoute) getLongURL(writer http.ResponseWriter, request *http.Request) {
	path := mux.Vars(request)["path"]
	route.logger.Info("activated getLongURL handler with path %s", path)

	getLongURLRequest := dto.GetLongURLRequest{
		ShortURL: path,
	}
	longURL, err := route.urlService.GetLongURL(getLongURLRequest.ToModel())
	if err != nil || longURL == "" {
		route.logger.Error(fmt.Sprintf("couldn't get long url from %s because %v", path, err))
		return
	}

	route.logger.Info(fmt.Sprintf("%s -> %s", path, longURL))

	http.Redirect(writer, request, longURL, http.StatusPermanentRedirect)
}
