package main

import (
	"context"

	"github.com/felipecoppola/go-boilerplate/internal/config"
	"github.com/felipecoppola/go-boilerplate/internal/infrastructure/os"

	"go.uber.org/zap"
)

// Globals that are set at build time.
var (
	appName    = ""
	appVersion = ""
)

func main() {
	defer func() { _ = zap.L().Sync() }()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := zap.L().With(
		zap.String("app_name", appName),
		zap.String("app_version", appVersion),
	)

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error("error creating config: %w", zap.Error(err))
		return
	}

	logger.Info("starting app")
	os.SignalListener(logger, cancel)

	if err = run(ctx, cfg); err != nil {
		logger.Error("command failed: %w", zap.Error(err))
		return
	}

	logger.Info("command completed")
}
