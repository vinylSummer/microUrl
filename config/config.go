package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Log      `yaml:"logger"`
		Postgres `yaml:"postgres"`
		SQLite   `yaml:"sqlite"`
	}

	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	Postgres struct {
		URL string `env-required:"true" yaml:"url" env:"PG_URL"`
	}

	SQLite struct {
		URL string `env-required:"true" yaml:"url" env:"SQLITE_URL"`
	}
)

const DefaultPath = "./config/config.yml"

func NewConfig() (*Config, error) {
	config := new(Config)

	// reads BOTH .yml and environ, env variables overwrite .yml ones
	err := cleanenv.ReadConfig(DefaultPath, config)
	if err != nil {
		return nil, fmt.Errorf("couldn't get configurations from .yml file or environ: %w", err)
	}

	return config, nil
}
