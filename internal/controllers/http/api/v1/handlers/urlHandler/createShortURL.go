package urlhandler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vinylSummer/microUrl/internal/services"
	"github.com/vinylSummer/microUrl/pkg/logger"
	"html/template"
	"net/http"
)

type CreateShortURLRoute struct {
	urlService services.URLService
	logger     logger.Interface
}

func NewCreateShortURLRoute(router *mux.Router, urlService services.URLService, logger logger.Interface) {
	route := &CreateShortURLRoute{
		urlService: urlService,
		logger:     logger,
	}

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./internal/controllers/http/api/v1/templates/main_page.html"))
		err := tmpl.Execute(w, nil)
		if err != nil {
			route.logger.Error(fmt.Errorf("couldn't execute guess_quest_piece template, %v", err))
		}
	}).Methods("GET")
	router.HandleFunc("/micrify", route.createShortURL).Methods("POST")
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

	createShortURLTemplate := template.Must(template.ParseFiles(
		"./internal/controllers/http/api/v1/templates/main_page.html"),
	)
	type createShortURLPageData struct {
		ErrorMessage string
		ShortURL     string
	}

	shortURL, err := route.urlService.CreateShortURL(longURL)
	if err != nil {
		route.logger.Info("Couldn't create short url in createShortURLHandler, %v", err)
		urlCreationFailedPageData := createShortURLPageData{
			ErrorMessage: "Service error, please try again later :(",
			ShortURL:     "",
		}
		err = createShortURLTemplate.Execute(writer, urlCreationFailedPageData)
		if err != nil {
			route.logger.Error("Couldn't render url creation failed template, %v", err)
		}
		return
	}

	urlCreationSuccessPageData := createShortURLPageData{
		ErrorMessage: "",
		ShortURL:     shortURL,
	}
	err = createShortURLTemplate.Execute(writer, urlCreationSuccessPageData)
	if err != nil {
		route.logger.Error("Couldn't render url creation success template, %v", err)
	}

}
