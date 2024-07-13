package unauthorized

import (
	"net/http"
	domain "sep_setting_mgr/internal/domain/pages"
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
	c.Response().Header().Set("HX-Retarget", "#page")
	c.Response().Header().Set("HX-Reswap", "innerHTML")
	if util.IsHTMX(c) {
		return util.RenderTempl(UnauthorizedPage(), c, http.StatusOK)
	}
	return util.RenderTempl(layouts.MainLayout(UnauthorizedPage()), c, http.StatusOK)
}
