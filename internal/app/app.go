package app

import (
	"fmt"
	"github.com/gorilla/mux"
	cfg "github.com/vinylSummer/microUrl/config"
	http "github.com/vinylSummer/microUrl/internal/controllers/http/api/v1"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/middleware"
	sqliteRepo "github.com/vinylSummer/microUrl/internal/repositories/urlRepository/sqlite"
	"github.com/vinylSummer/microUrl/internal/services"
	"github.com/vinylSummer/microUrl/pkg/httpServer"
	log "github.com/vinylSummer/microUrl/pkg/logger"
	"github.com/vinylSummer/microUrl/pkg/sqlite"
	"os"
	sig "os/signal"
	"syscall"
)

func Run(config *cfg.Config) {
	logger := log.New(config.Log.Level)

	fmt.Println(fmt.Sprintf("Starting MicroUrl with config:\n%#v", config))

	db, err := sqlite.New(config.SQLite.URL)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	urlRepo := sqliteRepo.New(db)

	urlService := services.NewURLService(urlRepo)

	handler := mux.NewRouter()
	handler.Use(middleware.CORS)

	http.NewRouter(handler, *urlService, logger)
	server := httpServer.New(
		handler,
		httpServer.Port(config.HTTP.Port),
	)
	fmt.Println("Serving at http://127.0.0.1:8080")

	interrupt := make(chan os.Signal, 1)
	sig.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case signal := <-interrupt:
		fmt.Println("Got signal: " + signal.String())
	case err = <-server.Notify():
		fmt.Println("Got error: " + err.Error())
	}

	fmt.Println("Shutting down the server..")
	err = server.Shutdown()
	if err != nil {
		logger.Error(fmt.Errorf("error shutting down http server: %w", err))
	}
}
