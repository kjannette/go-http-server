package main

import (
	"context"

	"github.com/kjannette/go_http_server/cmd/app/handler"
	"github.com/kjannette/go_http_server/internal/config"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context, cfg *config.Config) error {
	lg := zap.L().With(zap.String("app_name", appName), zap.String("app_version", appVersion))

	eg, ctx := errgroup.WithContext(ctx)
	app, h := handler.New(lg)
	serverLog := lg.Named("run_server")

	eg.Go(func() error {
		serverLog.Info("starting server")
		defer serverLog.Info("terminated server")
		return h.RunServer(ctx, cfg, app)
	})

	return eg.Wait()
}
