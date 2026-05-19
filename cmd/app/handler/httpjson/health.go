package httpjson

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) setupHealthCheckRoutes(app *echo.Echo) {
	app.GET("/health", h.healthCheck)
}

func (h *Handler) healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
