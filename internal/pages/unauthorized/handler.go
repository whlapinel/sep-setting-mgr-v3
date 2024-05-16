package unauthorized

import (
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type (
	Handler interface {
		// redirect after middleware credential check fails
		UnauthorizedHandler(c echo.Context) error
	}
	handler struct {
	}
)

func NewHandler() Handler {
	return &handler{}
}

func Mount(e *echo.Echo, h Handler) {
	e.GET("/unauthorized", h.UnauthorizedHandler)
}

func (h handler) UnauthorizedHandler(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(UnauthorizedPage(), c)
	}
	return util.RenderTempl(layouts.MainLayout(UnauthorizedPage()), c)
}
