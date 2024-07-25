package admin

import (
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/admin"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type AdminHandler interface {
	// GET /admin
	AdminHandler(c echo.Context) error
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
	common.AdminGroup.GET("", h.AdminHandler)
}

func (h handler) AdminHandler(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(views.AdminPage(views.AdminPageProps{R: router}), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage(views.AdminPageProps{R: router})), c, 200)

}
