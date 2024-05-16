package home

import (
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type (
	Handler interface {
		// Dashboard : GET /
		Home(e echo.Context) error
	}

	handler struct {
	}
)

func NewHandler() Handler {
	return &handler{}
}

func Mount(e *echo.Echo, h Handler) {
	e.GET("/", h.Home)
}

func (h handler) Home(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(HomePage(), c)
	}
	return util.RenderTempl(layouts.MainLayout(HomePage()), c)
}
