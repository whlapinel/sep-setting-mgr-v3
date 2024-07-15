package about

import (
	"net/http"
	"sep_setting_mgr/internal/handlers/components"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type AboutPageHandler interface {
	// GET : /about
	AboutPage(c echo.Context) error
}

type handler struct {
}

func NewHandler() AboutPageHandler {
	return &handler{}
}

func Mount(e *echo.Echo, h AboutPageHandler) {
	e.GET("/about", h.AboutPage)
}

func (h handler) AboutPage(c echo.Context) error {

	if !util.IsHTMX(c) {
		return util.RenderTempl(layouts.MainLayout(components.AboutPage()), c, http.StatusOK)
	}
	return util.RenderTempl(components.AboutPage(), c, http.StatusOK)
}
