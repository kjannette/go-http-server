package main

import (
	"context"

	"github.com/kjannette/go_http_server/internal/config"
	infrasignal "github.com/kjannette/go_http_server/internal/infrastructure/os"
	_ "github.com/kjannette/go_http_server/pkg/logger"

	"go.uber.org/zap"
)

// Globals that are set at build time.
var (
	appName    = "go-http-server"
	appVersion = "dev"
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
		logger.Error("error creating config", zap.Error(err))
		return
	}

	logger.Info("starting app")
	infrasignal.SignalListener(logger, cancel)

	if err = run(ctx, cfg); err != nil {
		logger.Error("command failed", zap.Error(err))
		return
	}

	logger.Info("command completed")
}
