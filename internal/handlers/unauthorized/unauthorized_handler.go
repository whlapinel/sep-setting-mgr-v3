package unauthorized

import (
	"net/http"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/components"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type UnauthorizedHandler interface {
	// redirect after middleware credential check fails
	Unauthorized(c echo.Context) error
}

type handler struct {
}

func NewHandler() UnauthorizedHandler {
	return &handler{}
}

func Mount(e *echo.Echo, h UnauthorizedHandler) {
	e.GET("/unauthorized", h.Unauthorized).Name = string(common.Unauthorized)
}

func (h handler) Unauthorized(c echo.Context) error {
	c.Response().Header().Set("HX-Retarget", "#page")
	c.Response().Header().Set("HX-Reswap", "innerHTML")
	if util.IsHTMX(c) {
		return util.RenderTempl(components.UnauthorizedPage(), c, http.StatusOK)
	}
	return util.RenderTempl(layouts.MainLayout(components.UnauthorizedPage()), c, http.StatusOK)
}
