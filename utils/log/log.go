package log

import (
	"context"
	"github.com/rs/zerolog"
	"os"
	"time"
)

const (
	serviceTag  = "service"
	versionTag  = "version"
	traceTag    = "trace_id"
	stepTag     = "step"
	fileTag     = "file"
	functionTag = "function"
)

type Level uint32

const (
	ErrorLevel Level = iota
	WarningLevel
	InfoLevel
	DebugLevel
)

var levels = map[Level]zerolog.Level{
	ErrorLevel:   zerolog.ErrorLevel,
	WarningLevel: zerolog.WarnLevel,
	InfoLevel:    zerolog.InfoLevel,
	DebugLevel:   zerolog.DebugLevel,
}

type Entry = zerolog.Event
type Log struct {
	log   zerolog.Logger
	level Level
}

type Logger interface {
	Info(ctx context.Context, traceID, step, message string, options ...Opts)
	Error(ctx context.Context, traceID, step, message string, options ...Opts)
	Warning(ctx context.Context, traceID, step, message string, options ...Opts)
	Debug(ctx context.Context, traceID, step, message string, options ...Opts)
}

func New(serviceName string, cf ...Config) *Log {
	config := applyConfig(cf...)
	logger := newZeroLogger(levels[config.level]).With().Str(serviceTag, serviceName).Logger()
	return &Log{
		log:   logger,
		level: config.level,
	}
}

func newZeroLogger(level zerolog.Level) *zerolog.Logger {
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(os.Stdout).Level(level).With().Timestamp().Logger()
	return &logger
}

func (l *Log) Info(ctx context.Context, traceID, step, message string, options ...Opts) {
	l.logMessage(zerolog.InfoLevel, ctx, traceID, step, message, options...)
}

func (l *Log) Warning(ctx context.Context, traceID, step, message string, options ...Opts) {
	l.logMessage(zerolog.WarnLevel, ctx, traceID, step, message, options...)
}

func (l *Log) Error(ctx context.Context, traceID, step, message string, options ...Opts) {
	l.logMessage(zerolog.ErrorLevel, ctx, traceID, step, message, options...)
}

func (l *Log) Debug(ctx context.Context, traceID, step, message string, options ...Opts) {
	l.logMessage(zerolog.DebugLevel, ctx, traceID, step, message, options...)
}

func (l *Log) logMessage(level zerolog.Level, ctx context.Context, traceID, step, message string, options ...Opts) {
	file, function := getCaller()
	logEntry := l.log.WithLevel(level).
		Str(traceTag, traceID).
		Str(fileTag, file).
		Str(functionTag, function).
		Str(stepTag, step)

	// Unir todos los campos de los Opts en un solo mapa
	attributes := make(map[string]interface{})
	for _, opt := range options {
		for key, value := range opt.field {
			attributes[key] = value
		}
	}

	// Agregar los atributos si hay alguno presente
	if len(attributes) > 0 {
		logEntry = logEntry.Interface("attributes", attributes)
	}

	logEntry.Msg(message)
}
