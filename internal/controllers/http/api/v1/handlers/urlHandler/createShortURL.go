package urlhandler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/handlers/urlHandler/dto"
	"github.com/vinylSummer/microUrl/internal/services/v1"
	"io"
	"net/http"
)

type CreateShortURLRoute struct {
	urlService v1.URLService
}

func NewCreateShortURLRoute(router *mux.Router, urlService v1.URLService) {
	route := &CreateShortURLRoute{
		urlService: urlService,
	}

	router.HandleFunc("/", route.createShortURL).Methods(http.MethodPost, http.MethodOptions)
}

func (route *CreateShortURLRoute) createShortURL(writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't read request body")
		return
	}

	var createShortURLRequest dto.CreateShortURLRequest
	err = json.Unmarshal(body, &createShortURLRequest)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't unmarshal request body")
		return
	}

	log.Trace().Msgf("activated createShortURL handler with longURL %s", createShortURLRequest.LongURL)

	URLBinding, err := route.urlService.CreateShortURL(request.Context(), createShortURLRequest.ToModel())
	if err != nil {
		log.Error().Err(err).Msg("Couldn't create short url in createShortURLHandler")
		urlCreationFailedData := dto.CreateShortURLResponse{
			ErrorMessage: "Service error, please try again later :(",
			ShortURL:     "",
		}
		err = json.NewEncoder(writer).Encode(urlCreationFailedData)
		if err != nil {
			log.Error().Err(err).Msg("Couldn't send json encoded response")
			return
		}
		return
	}

	urlCreationSuccessPageData := dto.CreateShortURLResponse{
		ErrorMessage: "",
		ShortURL:     URLBinding.ShortURL,
	}
	err = json.NewEncoder(writer).Encode(urlCreationSuccessPageData)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't send json encoded response")
		return
	}
}
