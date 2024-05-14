package urlhandler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/handlers/urlHandler/dto"
	"github.com/vinylSummer/microUrl/internal/services/v1"
	"github.com/vinylSummer/microUrl/pkg/logger"
	"io"
	"net/http"
)

type CreateShortURLRoute struct {
	urlService v1.URLService
	logger     logger.Interface
}

const MainPagePath = "./views/main_page.html"

func NewCreateShortURLRoute(router *mux.Router, urlService v1.URLService, logger logger.Interface) {
	route := &CreateShortURLRoute{
		urlService: urlService,
		logger:     logger,
	}

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, MainPagePath)
	}).Methods("GET", "OPTIONS")
	router.HandleFunc("/", route.createShortURL).Methods("POST", "OPTIONS")
}

func (route *CreateShortURLRoute) createShortURL(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		fmt.Println("only POST method is allowed")
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Couldn't read request body")
		return
	}

	var createShortURLRequest dto.CreateShortURLRequest
	err = json.Unmarshal(body, &createShortURLRequest)
	if err != nil {
		fmt.Println("Couldn't unmarshal request body")
		return
	}

	route.logger.Info("activated createShortURL handler with longURL %s", createShortURLRequest.LongURL)

	URLBinding, err := route.urlService.CreateShortURL(request.Context(), createShortURLRequest.ToModel())
	if err != nil {
		route.logger.Info("Couldn't create short url in createShortURLHandler, %v", err)
		urlCreationFailedData := dto.CreateShortURLResponse{
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

	urlCreationSuccessPageData := dto.CreateShortURLResponse{
		ErrorMessage: "",
		ShortURL:     URLBinding.ShortURL,
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
