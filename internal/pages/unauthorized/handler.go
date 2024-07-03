package unauthorized

import (
	"net/http"
	"sep_setting_mgr/internal/domain"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type handler struct {
}

func NewHandler() domain.UnauthorizedHandler {
	return &handler{}
}

func Mount(e *echo.Echo, h domain.UnauthorizedHandler) {
	e.GET("/unauthorized", h.UnauthorizedHandler)
}

func (h handler) UnauthorizedHandler(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(UnauthorizedPage(), c, http.StatusForbidden)
	}
	return util.RenderTempl(layouts.MainLayout(UnauthorizedPage()), c, http.StatusForbidden)
}
