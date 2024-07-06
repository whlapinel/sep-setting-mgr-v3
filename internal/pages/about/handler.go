package about

import (
	"net/http"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type (
	Handler interface {
		// redirect after middleware credential check fails
		AboutHandler(c echo.Context) error
	}
	handler struct {
	}
)

func NewHandler() Handler {
	return &handler{}
}

func Mount(e *echo.Echo, h Handler) {
	e.GET("/about", h.AboutHandler)
}

func (h handler) AboutHandler(c echo.Context) error {
	if !util.IsHTMX(c) {
		return util.RenderTempl(layouts.MainLayout(AboutPage()), c, http.StatusOK)
	}
	return util.RenderTempl(AboutPage(), c, http.StatusOK)
}
