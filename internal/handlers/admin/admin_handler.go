package admin

import (
	"log"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/components"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/services/admin"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type AdminHandler interface {
	// GET /admin
	AdminHandler(c echo.Context) error

	// Middleware for /admin/* routes
	Authorization(next echo.HandlerFunc) echo.HandlerFunc
}

type handler struct {
	adminService admin.AdminService
}

func NewHandler(adminService admin.AdminService) AdminHandler {
	return &handler{adminService}
}

func Mount(e *echo.Echo, h AdminHandler) {
	common.AdminGroup.Use(h.Authorization)
	common.AdminGroup.GET("", h.AdminHandler)
}

func (h handler) AdminHandler(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(components.AdminPage(), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.AdminPage()), c, 200)

}

func (h handler) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("id").(int)
		log.Println("User ID: ", userID)
		isAdmin := h.adminService.IsAdmin(userID)
		if !isAdmin {
			return util.RenderTempl(components.UnauthorizedPage(), c, 200)
		}
		return next(c)
	}
}
