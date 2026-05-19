package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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
