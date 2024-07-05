package admin

import (
	"sep_setting_mgr/internal/domain/pages"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/pages/admin/components"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service pages.AdminService
}

func NewHandler(svc pages.AdminService) pages.AdminHandler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h pages.AdminHandler) {
	r := e.Group("/admin")
	r.GET("", h.AdminHandler)
}

func (h handler) AdminHandler(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(components.AdminPage(), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.AdminPage()), c, 200)

}
