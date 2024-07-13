package admin

import (
	"log"
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/pages"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/pages/admin/components"
	"sep_setting_mgr/internal/pages/unauthorized"
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

var router *echo.Echo

func Mount(e *echo.Echo, h pages.AdminHandler) {
	router = e
	r := e.Group("/admin")
	r.Use(auth.AddCookieToHeader)
	r.Use(auth.JWTMiddleware)
	r.Use(auth.GetClaims)
	r.Use(h.Authorization)
	r.GET("", h.AdminHandler)
	r.GET("/calendar", h.Calendar)
	r.GET("/rooms", h.Rooms)
	r.GET("/rooms/add", h.ShowAddRoomForm)
	r.POST("/rooms", h.CreateRoom)
	r.GET("/users", h.Users)
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

func (h handler) ShowAddRoomForm(c echo.Context) error {
	log.SetPrefix("AdminHandler: ShowAddRoomForm()")
	if util.IsHTMX(c) {
		return util.RenderTempl(components.AddRoomForm(false, &models.Room{}), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.AdminPage()), c, 200)
}

func (h handler) ShowEditRoomForm(c echo.Context) error {
	log.SetPrefix("AdminHandler: ShowEditRoomForm()")
	roomID, err := strconv.Atoi(c.Param("room-id"))
	if err != nil {
		return c.String(400, "Invalid room ID")
	}
	room, err := h.service.FindRoomByID(roomID)
	if err != nil {
		return c.String(500, "Failed to get room. See server logs for details.")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(components.AddRoomForm(true, room), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.AdminPage()), c, 200)
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

func (h handler) Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Get("id").(int)
		isAdmin := h.service.IsAdmin(userID)
		if !isAdmin {
			return util.RenderTempl(unauthorized.UnauthorizedPage(), c, 200)
		}
		return next(c)
	}
}
