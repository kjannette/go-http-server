package os

import (
	"context"
	stdlib "os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

// SignalListener listens for SIGINT and SIGTERM and cancels the root context.
func SignalListener(logger *zap.Logger, cancel context.CancelFunc) {
	sigCh := make(chan stdlib.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		logger.Info("received shutdown signal", zap.String("signal", sig.String()))
		cancel()
	}()
}
