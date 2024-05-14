package main

import (
	"github.com/rs/zerolog/log"
	cfg "github.com/vinylSummer/microUrl/config"
	"github.com/vinylSummer/microUrl/internal/app"
)

func main() {
	config, err := cfg.NewConfig()
	if err != nil {
		log.Error().Err(err).Msg("Could not load configurations")
	}
	app.Run(config)
}
