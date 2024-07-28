package admin

import (
	common "sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/components/componentscommon"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/admin"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type AdminHandler interface {
	// GET /admin
	AdminPage(c echo.Context) error
}

type handler struct {
	adminService admin.AdminService
}

func NewHandler(adminService admin.AdminService) AdminHandler {
	return &handler{adminService}
}

var router *echo.Echo

func Mount(e *echo.Echo, h AdminHandler) {
	router = e
	common.AdminGroup.GET("", h.AdminPage).Name = common.AdminPage.String()
}

func (h handler) AdminPage(c echo.Context) error {

	template := componentscommon.Templify(views.NewAdminPage(router))
	if util.IsHTMX(c) {
		return util.RenderTempl((template), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(template), c, 200)

}
