package about

import (
	"net/http"
	"sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
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
	e.GET("/about", h.AboutPage).Name = handlerscommon.AboutPage.String()
}

func (h handler) AboutPage(c echo.Context) error {

	if !util.IsHTMX(c) {
		return util.RenderTempl(layouts.MainLayout(views.AboutPage(), nil), c, http.StatusOK)
	}
	return util.RenderTempl(views.AboutPage(), c, http.StatusOK)
}
