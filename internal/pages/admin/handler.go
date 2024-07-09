package admin

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/pages"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/pages/admin/components"
	"sep_setting_mgr/internal/util"
	"strconv"

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
	r.GET("/calendar", h.Calendar)
	r.GET("/rooms", h.Rooms)
	r.POST("/rooms", h.CreateRoom)
}

func (h handler) AdminHandler(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(components.AdminPage(), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.AdminPage()), c, 200)

}

func (h handler) Calendar(c echo.Context) error {
	var assignments models.Assignments
	assignments, err := h.service.GetAllAssignments()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving assignments")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(components.AdminCalendarComponent(assignments), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.AdminPage()), c, 200)
}

func (h handler) Rooms(c echo.Context) error {
	log.SetPrefix("AdminHandler: Rooms()")
	rooms, err := h.service.ListRooms()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving rooms")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(components.RoomsTableComponent(rooms), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.AdminPage()), c, 200)
}

func (h handler) CreateRoom(c echo.Context) error {
	log.SetPrefix("AdminHandler: CreateRoom()")
	var room models.Room
	room.Name = c.FormValue("room-name")
	room.Number = c.FormValue("room-number")
	priority, err := strconv.Atoi(c.FormValue("priority"))
	if err != nil {
		return c.String(400, "Error getting priority value")
	}
	room.MaxCapacity, err = strconv.Atoi(c.FormValue("capacity"))
	if err != nil {
		return c.String(400, "Error getting capacity value")
	}
	room.Priority = priority
	id, err := h.service.AddRoom(&room)
	if err != nil {
		return c.String(500, "Error adding room")
	}
	room.ID = id
	return util.RenderTempl(components.RoomsRowComponent(&room), c, 201)
}
