package users

import (
	"log"
	"sep_setting_mgr/internal/services/users"

	handlerscommon "sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/pages/admin/components"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type UsersHandler interface {
	// GET /admin/users
	Users(c echo.Context) error
}

type handler struct {
	service users.UsersService
}

func NewHandler(service users.UsersService) UsersHandler {
	return &handler{service}

}

func Mount(e *echo.Echo, h UsersHandler) {
	handlerscommon.AdminGroup.GET("/users", h.Users)
}

func (h handler) Users(c echo.Context) error {
	log.SetPrefix("AdminHandler: Users()")
	users, err := h.service.ListUsers()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving users")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(components.UsersTableComponent(users), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.AdminPage()), c, 200)
}
