package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

type Logger struct {
	logger *zerolog.Logger
}

var _ Interface = (*Logger)(nil)

func New(levelString string) *Logger {
	var level zerolog.Level

	switch strings.ToLower(levelString) {
	case "error":
		level = zerolog.ErrorLevel
	case "warn":
		level = zerolog.WarnLevel
	case "info":
		level = zerolog.InfoLevel
	case "debug":
		level = zerolog.DebugLevel
	default:
		level = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(level)

	skipFrameCount := 3
	newLogger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Logger{
		logger: &newLogger,
	}
}

func (logger *Logger) Debug(message any, args ...any) {
	logger.msg("debug", message, args...)
}

func (logger *Logger) Info(message string, args ...any) {
	logger.log(message, args...)
}

func (logger *Logger) Warn(message string, args ...any) {
	logger.log(message, args...)
}

func (logger *Logger) Error(message any, args ...any) {
	if logger.logger.GetLevel() == zerolog.DebugLevel {
		logger.Debug(message, args...)
	}

	logger.msg("error", message, args...)
}

func (logger *Logger) Fatal(message any, args ...any) {
	logger.msg("fatal", message, args...)

	os.Exit(1)
}

func (logger *Logger) log(message string, args ...any) {
	if len(args) == 0 {
		logger.logger.Info().Msg(message)
	} else {
		logger.logger.Info().Msgf(message, args...)
	}
}

func (logger *Logger) msg(level string, message any, args ...any) {
	switch msg := message.(type) {
	case error:
		logger.log(msg.Error(), args...)
	case string:
		logger.log(msg, args...)
	default:
		logger.log(fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
