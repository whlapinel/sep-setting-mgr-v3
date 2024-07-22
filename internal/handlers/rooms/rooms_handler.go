package rooms

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/rooms"
	"sep_setting_mgr/internal/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RoomsHandler interface {
	// GET /admin/rooms
	Rooms(c echo.Context) error

	// GET /admin/rooms/add
	ShowAddRoomForm(c echo.Context) error

	// GET /admin/rooms/:room-id/edit
	ShowEditRoomForm(c echo.Context) error

	// POST /admin/rooms
	CreateRoom(c echo.Context) error

	// DELETE /admin/rooms/:room-id
	DeleteRoom(c echo.Context) error

	// POST /admin/rooms/:room-id
	EditRoom(c echo.Context) error
}

type handler struct {
	service rooms.RoomsService
}

func NewHandler(service rooms.RoomsService) RoomsHandler {
	return &handler{
		service,
	}
}

var router *echo.Echo

func Mount(e *echo.Echo, h RoomsHandler) {
	router = e
	common.RoomsGroup.GET("", h.Rooms).Name = string(common.Rooms)
	common.RoomsGroup.GET("/add", h.ShowAddRoomForm).Name = string(common.ShowAddRoomForm)
	common.RoomsGroup.POST("", h.CreateRoom).Name = string(common.CreateRoom)
	common.RoomsIDGroup.GET("/edit", h.ShowEditRoomForm).Name = string(common.ShowEditRoomForm)
	common.RoomsIDGroup.POST("", h.EditRoom).Name = string(common.EditRoom)
	common.RoomsIDGroup.DELETE("", h.DeleteRoom).Name = string(common.DeleteRoom)
}

func (h handler) Rooms(c echo.Context) error {
	log.SetPrefix("AdminHandler: Rooms()")
	rooms, err := h.service.ListRooms()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving rooms")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(views.RoomsTableComponent(rooms, router), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage(router)), c, 200)
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
	return util.RenderTempl(views.RoomsRowComponent(&room, router), c, 201)
}

func (h handler) EditRoom(c echo.Context) error {
	log.SetPrefix("AdminHandler: EditRoom()")
	roomID, err := strconv.Atoi(c.Param("room-id"))
	if err != nil {
		return c.String(400, "Invalid room ID")
	}
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
	room.ID = roomID
	err = h.service.UpdateRoom(&room)
	if err != nil {
		return c.String(500, "Error updating room")
	}
	return util.RenderTempl(views.RoomsRowComponent(&room, router), c, 200)
}

func (h handler) DeleteRoom(c echo.Context) error {
	log.SetPrefix("AdminHandler: DeleteRoom()")
	roomID, err := strconv.Atoi(c.Param("room-id"))
	if err != nil {
		return c.String(400, "Invalid room ID")
	}
	err = h.service.DeleteRoom(roomID)
	if err != nil {
		return c.String(500, "Failed to delete room")
	}
	return c.NoContent(200)
}

func (h handler) ShowAddRoomForm(c echo.Context) error {
	log.SetPrefix("AdminHandler: ShowAddRoomForm()")
	if util.IsHTMX(c) {
		return util.RenderTempl(views.AddRoomForm(false, &models.Room{}), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage(router)), c, 200)
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
		return util.RenderTempl(views.AddRoomForm(true, room), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage(router)), c, 200)
}
