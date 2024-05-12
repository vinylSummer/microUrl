package app

import (
	"fmt"
	"github.com/gorilla/mux"
	cfg "github.com/vinylSummer/microUrl/config"
	http "github.com/vinylSummer/microUrl/internal/controllers/http/api/v1"
	sqliteRepo "github.com/vinylSummer/microUrl/internal/repositories/urlRepository/sqlite"
	"github.com/vinylSummer/microUrl/internal/services"
	"github.com/vinylSummer/microUrl/pkg/httpServer"
	log "github.com/vinylSummer/microUrl/pkg/logger"
	"github.com/vinylSummer/microUrl/pkg/sqlite"
)

func Run(config *cfg.Config) {
	logger := log.New(config.Log.Level)

	logger.Info(fmt.Sprintf("Starting MicroUrl with config:\n%#v", config))

	db, err := sqlite.New(config.SQLite.URL)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	urlRepo := sqliteRepo.New(db)

	urlService := services.NewURLService(urlRepo)

	handler := mux.NewRouter()
	http.NewRouter(handler, *urlService, logger)
	_ = httpServer.New(handler, httpServer.Port(config.HTTP.Port))
}
