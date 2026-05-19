package httpjson

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kjannette/go-http-server/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Handler has services that are required by the handler.
type Handler struct {
	logger *zap.Logger
}

// New creates a new Handler instance.
func New(logger *zap.Logger) (*echo.Echo, Handler) {
	h := Handler{
		logger: logger,
	}

	e := echo.New()
	h.setupHealthCheckRoutes(e)

	return e, h
}

// RunServer runs http server and listens until context is canceled.
func (h *Handler) RunServer(ctx context.Context, conf *config.Config, e *echo.Echo) error {
	go func() {
		<-ctx.Done()

		timeout, cancel := context.WithTimeout(context.Background(), time.Duration(conf.API.ReadTimeout)*time.Second)
		defer cancel()

		_ = e.Shutdown(timeout) // try gracefully first
		_ = e.Close()           // immediate termination
	}()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.API.Port),
		ReadTimeout:  time.Duration(conf.API.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(conf.API.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(conf.API.IdleTimeout) * time.Second,
	}

	if err := e.StartServer(server); !errors.Is(err, http.ErrServerClosed) {
		return errors.Wrap(err, "server error")
	}

	return nil
}
