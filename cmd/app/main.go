package main

import (
	cfg "github.com/vinylSummer/microUrl/config"
	"github.com/vinylSummer/microUrl/internal/app"
	"log"
)

func main() {
	config, err := cfg.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(config)
}
