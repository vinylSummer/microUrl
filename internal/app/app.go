package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	cfg "github.com/vinylSummer/microUrl/config"
	http "github.com/vinylSummer/microUrl/internal/controllers/http/api/v1"
	"github.com/vinylSummer/microUrl/internal/controllers/http/api/v1/middleware"
	sqliteRepo "github.com/vinylSummer/microUrl/internal/repositories/urlRepository/sqlite"
	v1 "github.com/vinylSummer/microUrl/internal/services/v1"
	"github.com/vinylSummer/microUrl/pkg/httpServer"
	"github.com/vinylSummer/microUrl/pkg/logger"
	"github.com/vinylSummer/microUrl/pkg/sqlite"
	"os"
	sig "os/signal"
	"syscall"
)

func Run(config *cfg.Config) {
	logger.NewLogger(config)

	prettyConfig, _ := json.MarshalIndent(config, "", "  ")
	log.Info().Msgf("Starting MicroUrl with config:\n%+v", string(prettyConfig))

	db, err := sqlite.New(config.SQLite.URL)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to database")
	}
	defer db.Close()
	log.Info().Msg("Successfully connected to the database")

	urlRepo := sqliteRepo.New(db)

	urlService := v1.NewURLService(urlRepo)

	handler := mux.NewRouter()
	handler.Use(middleware.CORS)

	http.NewRouter(handler, *urlService)
	server := httpServer.New(
		handler,
		httpServer.Port(config.HTTP.Port),
	)

	log.Info().Msgf("Serving at http://127.0.0.1:%s", config.HTTP.Port)

	interrupt := make(chan os.Signal, 1)
	sig.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case signal := <-interrupt:
		log.Info().Msgf("Caught signal %v", signal)
	case err = <-server.Notify():
		log.Error().Err(err).Msg("An error occurred while serving requests")
	}

	log.Info().Msg("Shutting down the server..")

	err = server.Shutdown()
	if err != nil {
		log.Error().Err(err).Msg("Could not shut down the server")
	}
}
