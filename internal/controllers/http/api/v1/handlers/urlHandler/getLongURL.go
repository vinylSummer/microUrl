package urlhandler

import (
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/handlers/urlHandler/dto"
	"github.com/vinylSummer/microUrl/internal/services/v1"
	"net/http"
)

type GetLongURLRoute struct {
	urlService v1.URLService
}

func NewGetLongURLRoute(router *mux.Router, urlService v1.URLService) {
	route := &GetLongURLRoute{
		urlService: urlService,
	}

	router.HandleFunc("/{path:[a-zA-Z0-9]+}", route.getLongURL).Methods(http.MethodGet, http.MethodOptions)
}

func (route *GetLongURLRoute) getLongURL(writer http.ResponseWriter, request *http.Request) {
	path := mux.Vars(request)["path"]
	log.Trace().Msgf("activated getLongURL handler with path %s", path)

	getLongURLRequest := dto.GetLongURLRequest{
		ShortURL: path,
	}
	longURL, err := route.urlService.GetLongURL(request.Context(), getLongURLRequest.ToModel())
	if err != nil || longURL == "" {
		log.Error().Err(err).Msgf("couldn't get long url from %s", path)
		return
	}

	log.Info().Msgf("Resolved %s -> %s", path, longURL)

	http.Redirect(writer, request, longURL, http.StatusPermanentRedirect)
}
