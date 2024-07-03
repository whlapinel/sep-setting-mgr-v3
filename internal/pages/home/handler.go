package home

import (
	"sep_setting_mgr/internal/domain"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type (
	handler struct {
	}
)

func NewHandler() domain.HomeHandler {
	return &handler{}
}

func Mount(e *echo.Echo, h domain.HomeHandler) {
	e.GET("/", h.Home)
}

func (h handler) Home(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(HomePage(), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(HomePage()), c, 200)
}
