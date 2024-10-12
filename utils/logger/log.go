package logger

import (
	"github.com/juanMaAV92/user-auth-api/config"
	"github.com/rs/zerolog"
	slogzerolog "github.com/samber/slog-zerolog/v2"
	"log/slog"
	"os"
)

const (
	serviceTag  = "service"
	versionTag  = "version"
	traceTag    = "trace_id"
	stepTag     = "step"
	fileTag     = "file"
	functionTag = "function"
)

type Log struct {
	log slog.Logger
}

type Logger interface {
	Info(traceID, message string, options ...Opts)
}

func NewLogger(cfg config.Config) *Log {

	zeroLogger := zerolog.New(os.Stdout).With().Logger()
	logger := slog.New(
		slogzerolog.Option{
			Level:  slog.LevelDebug,
			Logger: &zeroLogger,
		}.NewZerologHandler())

	logger = logger.
		With(serviceTag, cfg.AppName).
		With(versionTag, "1.0.0")

	return &Log{
		log: *logger,
	}
}

func (l *Log) Info(traceID, message string, options ...Opts) {
	file, function := getCaller()
	logEntry := l.log.With(traceTag, traceID).With(fileTag, file).With(functionTag, function)
	attributes := make(map[string]interface{})
	for _, opt := range options {
		for key, value := range opt.field {
			attributes[key] = value
		}
	}

	logEntry = logEntry.With("attributes", attributes)
	logEntry.Info(message)
}
