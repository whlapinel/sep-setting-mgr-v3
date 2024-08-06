package rooms

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	common "sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/views"
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

	// POST /admin/rooms/:room-id/promote
	PromoteRoom(c echo.Context) error
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
	common.RoomsGroup.GET("", h.Rooms).Name = common.Rooms.String()
	common.RoomsGroup.GET("/add", h.ShowAddRoomForm).Name = common.ShowAddRoomForm.String()
	common.RoomsGroup.POST("", h.CreateRoom).Name = common.CreateRoom.String()
	common.RoomsIDGroup.GET("/edit", h.ShowEditRoomForm).Name = common.ShowEditRoomForm.String()
	common.RoomsIDGroup.POST("", h.EditRoom).Name = common.EditRoom.String()
	common.RoomsIDGroup.DELETE("", h.DeleteRoom).Name = common.DeleteRoom.String()
	common.RoomsIDGroup.POST("/promote", h.PromoteRoom).Name = common.PromoteRoom.String()
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
	return c.Redirect(303, router.Reverse(common.AdminPage.String()))
}

func (h handler) CreateRoom(c echo.Context) error {
	log.SetPrefix("AdminHandler: CreateRoom()")
	var room models.Room
	room.Name = c.FormValue("room-name")
	room.Number = c.FormValue("room-number")
	maxCapacity, err := strconv.Atoi(c.FormValue("capacity"))
	if err != nil {
		return c.String(400, "Error getting capacity value")
	}
	room.MaxCapacity = maxCapacity
	room.Priority, err = h.service.GetNextPriority()
	if err != nil {
		return err
	}
	err = h.service.AddRoom(&room)
	if err != nil {
		return err
	}
	rooms, err := h.service.ListRooms()
	if err != nil {
		return err
	}
	return util.RenderTempl(views.RoomsTableComponent(rooms, router), c, 200)
}

func (h handler) PromoteRoom(c echo.Context) error {
	roomID, err := strconv.Atoi(c.Param("room-id"))
	if err != nil {
		return err
	}
	err = h.service.PromoteRoom(roomID)
	if err != nil {
		return err
	}
	rooms, err := h.service.ListRooms()
	if err != nil {
		return err
	}
	return util.RenderTempl(views.RoomsTableComponent(rooms, router), c, 200)
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
	room.MaxCapacity, err = strconv.Atoi(c.FormValue("capacity"))
	if err != nil {
		return c.String(400, "Error getting capacity value")
	}
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
	return c.Redirect(303, router.Reverse(common.AdminPage.String()))
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
	return c.Redirect(303, router.Reverse(common.AdminPage.String()))
}
