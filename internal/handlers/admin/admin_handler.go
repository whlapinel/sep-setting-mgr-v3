package admin

import (
	"sep_setting_mgr/internal/domain/models"
	common "sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/components/componentscommon"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/admin"
	"sep_setting_mgr/internal/services/assignments"
	"sep_setting_mgr/internal/services/rooms"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type AdminHandler interface {
	// GET /admin
	AdminPage(c echo.Context) error
}

type handler struct {
	adminService admin.AdminService
	assignments  assignments.AssignmentsService
	rooms        rooms.RoomsService
}

func NewHandler(adminService admin.AdminService,
	assignments assignments.AssignmentsService,
	rooms rooms.RoomsService) AdminHandler {
	return &handler{adminService, assignments, rooms}
}

var router *echo.Echo

func Mount(e *echo.Echo, h AdminHandler) {
	router = e
	common.AdminGroup.GET("", h.AdminPage).Name = common.AdminPage.String()
}

func (h handler) AdminPage(c echo.Context) error {
	// get all assignments
	assignments, err := h.assignments.All()
	if err != nil {
		return err
	}
	assignmentsMap := assignments.MapForCalendar()
	// get all rooms
	rooms, err := h.rooms.ListRooms()
	if err != nil {
		return err
	}
	template := componentscommon.Templify(views.NewAdminPage(router, assignmentsMap, rooms))
	if util.IsHTMX(c) {
		return util.RenderTempl((template), c, 200)
	}
	user, err := models.NewUser(c.Get("first").(string), c.Get("last").(string), c.Get("email").(string), c.Get("picture").(string))
	if err != nil {
		return err
	}
	return util.RenderTempl(layouts.MainLayout(template, user), c, 200)

}
