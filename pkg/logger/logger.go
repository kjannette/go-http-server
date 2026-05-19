package logger

import (
    "log"
    "os"
    "go.uber.org/zap"
    "go.uber.go/zap/zapcore"
)

type Logger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	With(fields ...zap.Field) *zap.Logger
}

// Option is a function to provide the Logger creation with extra initialisation options.
type Option func(*zap.Config)

// init initializes zap's global loggers (i.e. zap.L() and zap.S()).
// Log level is retrieved from environment variable LOG_LEVEL, defaulting to "info" otherwise.
// To activate this facility, add a blank import to this package.
func init() {
	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "info"
	}

	initializeZapGlobals(level)
}

func initializeZapGlobals(level string) {
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		log.Fatalf("failed to parse logger level: %v", err)
	}

	log.Printf("setting logging level to '%s'.\n", zapLevel.String())

	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zapLevel)
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder

	logger, err := cfg.Build()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	zap.ReplaceGlobals(logger)
}
