package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/vinylSummer/microUrl/config"
	"os"
	"time"
)

const (
	debugLevel = "debug"
	infoLevel  = "info"
	warnLevel  = "warn"
	errorLevel = "error"
)

func NewLogger(cfg *config.Config) {
	var level zerolog.Level

	switch cfg.Log.Level {
	case debugLevel:
		level = zerolog.DebugLevel
	case infoLevel:
		level = zerolog.InfoLevel
	case warnLevel:
		level = zerolog.WarnLevel
	case errorLevel:
		level = zerolog.ErrorLevel
	default:
		level = zerolog.InfoLevel
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"})
	zerolog.TimeFieldFormat = time.RFC822
	log.Logger = log.With().Caller().Logger()

	zerolog.SetGlobalLevel(level)
}
