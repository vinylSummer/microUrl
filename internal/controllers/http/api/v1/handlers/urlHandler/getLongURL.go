package urlhandler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/handlers/urlHandler/dto"
	"github.com/vinylSummer/microUrl/internal/repositories/urlRepository/sqlite"
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

// @Summary Redirect to long URL
// @Description When a user sends a request with a short URL, they should be redirected to the corresponding long URL
// @Tags URLs
// @Accept json
// @Param url path string true "Short URL"
// @Success 301
// @Header 301 {string} Location "Long URL for the provided short URL"
// @Failure 404 {object} dto.GetLongURLResponse
// @Failure 500 {object} dto.GetLongURLResponse
// @Router /{url} [get]
func (route *GetLongURLRoute) getLongURL(writer http.ResponseWriter, request *http.Request) {
	path := mux.Vars(request)["path"]
	log.Trace().Msgf("activated getLongURL handler with path %s", path)

	getLongURLRequest := dto.GetLongURLRequest{
		ShortURL: path,
	}

	longURL, err := route.urlService.GetLongURL(request.Context(), getLongURLRequest.ToModel())
	var urlNotFoundErr *sqlite.ErrURLNotFound
	if errors.As(err, &urlNotFoundErr) {
		log.Warn().Msgf("long url from %s is not present in the database", path)
		urlRetrievalFailedData := dto.GetLongURLResponse{
			ErrorMessage: "No such URL",
			LongURL:      "",
		}
		writer.WriteHeader(http.StatusNotFound)
		err = json.NewEncoder(writer).Encode(urlRetrievalFailedData)
		if err != nil {
			log.Error().Err(err).Msg("Couldn't send json encoded response")
			return
		}
		return
	}
	if err != nil {
		log.Error().Err(err).Msgf("couldn't get long url from %s", path)
		urlRetrievalFailedData := dto.GetLongURLResponse{
			ErrorMessage: "Service error, please try again later :(",
			LongURL:      "",
		}
		writer.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(writer).Encode(urlRetrievalFailedData)
		if err != nil {
			log.Error().Err(err).Msg("Couldn't send json encoded response")
			return
		}
		return
	}

	log.Info().Msgf("Resolved %s -> %s", path, longURL)

	http.Redirect(writer, request, longURL, http.StatusMovedPermanently)
}
