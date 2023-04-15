package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
)

var globalLogger *zap.Logger

func Init(dev bool) {
	var err error
	globalLogger, err = New(dev)

	if err != nil {
		log.Fatalf("init logger: %s", err)
	}
}

func New(dev bool) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error
	if dev {
		logger, err = zap.NewDevelopment()
	} else {
		cfg := zap.NewProductionConfig()
		cfg.DisableCaller = true
		cfg.DisableStacktrace = true
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		logger, err = cfg.Build()
	}

	if err != nil {
		return nil, errors.Wrap(err, "create logger")
	}

	return logger, nil
}

func Debug(msg string, fields ...zap.Field) {
	globalLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	globalLogger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	globalLogger.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	globalLogger.Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	globalLogger.Fatal(msg, fields...)
}
