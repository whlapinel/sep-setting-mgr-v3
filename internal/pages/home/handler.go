package home

import (
	"sep_setting_mgr/internal/templates/layouts"

	"github.com/labstack/echo"
)

type (
	Handler interface {
		// Dashboard : GET /
		Home(e echo.Context) error
	}

	handler struct {
		service Service
	}
)

func NewHandler(svc Service) Handler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h Handler) {
	e.GET("/", h.Home)
}

func (h handler) Home(c echo.Context) error {
	return layouts.MainLayout(HomePage()).Render(c.Request().Context(), c.Response().Writer)
}
