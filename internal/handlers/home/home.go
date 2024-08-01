package home

import (
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type HomeHandler interface {
	Home(e echo.Context) error
}

type (
	handler struct {
	}
)

func NewHandler() HomeHandler {
	return &handler{}
}

func Mount(e *echo.Echo, h HomeHandler) {
	e.GET("/", h.Home)
}

func (h handler) Home(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(views.HomePage(), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.HomePage(), nil), c, 200)
}
