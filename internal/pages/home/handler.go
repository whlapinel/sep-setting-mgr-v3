package home

import (
	"sep_setting_mgr/internal/domain/pages"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type (
	handler struct {
	}
)

func NewHandler() pages.HomeHandler {
	return &handler{}
}

func Mount(e *echo.Echo, h pages.HomeHandler) {
	e.GET("/", h.Home)
}

func (h handler) Home(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(HomePage(), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(HomePage()), c, 200)
}
