package router

import (
	"tommynurwantoro/mygoweb/internal/http/handler"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, h *handler.Handler) {
	e.GET("/healthz", h.General.HealthzHandler)
	e.GET("/", h.Home.Index)
	e.GET("/payme", h.Payme.Index)
}
