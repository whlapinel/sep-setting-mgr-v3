package home

import (
	"sep_setting_mgr/internal/handlers/components"
	"sep_setting_mgr/internal/layouts"
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
		return util.RenderTempl(components.HomePage(), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.HomePage()), c, 200)
}
