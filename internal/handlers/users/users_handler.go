package users

import (
	"log"
	common "sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/users"
	"sep_setting_mgr/internal/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UsersHandler interface {
	// GET /admin/users
	Users(c echo.Context) error

	// GET /admin/users/:user-id/edit
	ShowEditUserForm(c echo.Context) error

	// POST /admin/users/:user-id/
	EditUser(c echo.Context) error
}

type handler struct {
	service users.UsersService
}

func NewHandler(service users.UsersService) UsersHandler {
	return &handler{service}

}

var router *echo.Echo

func Mount(e *echo.Echo, h UsersHandler) {
	router = e
	common.UsersGroup.GET("", h.Users).Name = string(common.Users)
	common.UserIDGroup.GET("/edit", h.ShowEditUserForm).Name = string(common.ShowEditUserForm)
	common.UserIDGroup.POST("", h.EditUser).Name = string(common.EditUser)
}

func (h handler) Users(c echo.Context) error {
	log.SetPrefix("AdminHandler: Users()")
	users, err := h.service.ListUsers()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving users")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(views.UsersTableComponent(users, router), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage(views.AdminPageProps{R: router})), c, 200)
}

// GET /admin/users/:user-id/edit
func (h handler) ShowEditUserForm(c echo.Context) error {
	adminID := c.Get("id").(int)
	userID, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		log.Println(err)
		return c.String(400, "Invalid user ID")
	}
	user, err := h.service.FindUserByID(userID)
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving user")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(views.UserFormComponent(adminID, user, router), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage(views.AdminPageProps{R: router})), c, 200)

}

func (h handler) EditUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		log.Println(err)
		return c.String(400, "Invalid user ID")
	}
	user, err := h.service.FindUserByID(userID)
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving user")
	}
	user.Email = c.FormValue("email")
	if c.FormValue("admin") != "" {
		user.Admin = c.FormValue("admin") == "yes"
	}
	err = h.service.UpdateUser(user)
	if err != nil {
		log.Println(err)
		return c.String(500, "Error updating user")
	}
	return util.RenderTempl(views.UsersRowComponent(user, router), c, 200)

}
