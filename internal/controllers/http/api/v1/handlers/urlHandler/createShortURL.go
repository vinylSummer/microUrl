package urlhandler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vinylSummer/microUrl/internal/services"
	"github.com/vinylSummer/microUrl/pkg/logger"
	"net/http"
)

type CreateShortURLRoute struct {
	urlService services.URLService
	logger     logger.Interface
}

const MainPagePath = "./views/main_page.html"

func NewCreateShortURLRoute(router *mux.Router, urlService services.URLService, logger logger.Interface) {
	route := &CreateShortURLRoute{
		urlService: urlService,
		logger:     logger,
	}

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, MainPagePath)
	}).Methods("GET")
	router.HandleFunc("/", route.createShortURL).Methods("POST", "OPTIONS")
}

func (route *CreateShortURLRoute) createShortURL(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		route.logger.Info("Couldn't parse form in createShortURLHandler, %v", err)
		return
	}

	longURL := request.PostFormValue("longURL")

	route.logger.Info("activated createShortURL handler with longURL %s", longURL)

	type createShortURLData struct {
		ErrorMessage string `json:"error_message"`
		ShortURL     string `json:"short_url"`
	}

	shortURL, err := route.urlService.CreateShortURL(longURL)
	if err != nil {
		route.logger.Info("Couldn't create short url in createShortURLHandler, %v", err)
		urlCreationFailedData := createShortURLData{
			ErrorMessage: "Service error, please try again later :(",
			ShortURL:     "",
		}
		err = json.NewEncoder(writer).Encode(urlCreationFailedData)
		if err != nil {
			fmt.Printf(
				"Couldn't send json encoded response %v in createShortURL handler, %v",
				urlCreationFailedData,
				err,
			)
			return
		}
		return
	}

	urlCreationSuccessPageData := createShortURLData{
		ErrorMessage: "",
		ShortURL:     shortURL,
	}
	err = json.NewEncoder(writer).Encode(urlCreationSuccessPageData)
	if err != nil {
		fmt.Printf(
			"Couldn't send json encoded response %v in createShortURL handler, %v",
			urlCreationSuccessPageData,
			err,
		)
		return
	}
}
