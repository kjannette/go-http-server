package main

import (
	"context"

	"github.com/kjannette/go-http-server/cmd/app/handler/httpjson"
	"github.com/kjannnette/go-http-server/internal/config"
	"github.com/kjannette/go-http-server/pkg/logger"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func run(ctx context.Context, cfg *config.Config) error {
	lg := zap.L().With(zap.String("app_name", appName), zap.String("app_version", appVersion))

	eg, ctx := errgroup.WithContext(ctx)
	app, h := httpjson.New(lg)

	eg.Go(func(l logger.Logger) func() error {
		return func() error {
			l.Info("starting server")
			defer l.Info("terminated server")
			return h.RunServer(ctx, cfg, app)
		}
	}(lg.Named("run_server")))

	return eg.Wait()
}
