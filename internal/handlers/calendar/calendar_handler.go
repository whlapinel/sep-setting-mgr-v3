package calendar

import (
	"log"
	common "sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/assignments"
	"sep_setting_mgr/internal/services/rooms"
	"sep_setting_mgr/internal/services/students"
	testevents "sep_setting_mgr/internal/services/test_events"
	"sep_setting_mgr/internal/util"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type CalendarHandler interface {
	// GET /admin/calendar
	Calendar(c echo.Context) error

	// GET /admin/calendar/:date/details
	AdminCalendarDetails(c echo.Context) error

	// GET /dashboard/calendar/:date/details
	DBCalendarDetails(c echo.Context) error

	// GET /admin/calendar/assign-room/:assignment-id
	ShowAssignRoomForm(c echo.Context) error

	// POST /admin/calendar/assign-room/:assignment-id
	AssignRoom(c echo.Context) error
}

type handler struct {
	assignments assignments.AssignmentsService
	rooms       rooms.RoomsService
	testEvents  testevents.TestEventsService
	students    students.StudentsService
}

func NewHandler(assignments assignments.AssignmentsService, rooms rooms.RoomsService, testEvents testevents.TestEventsService, students students.StudentsService) CalendarHandler {
	return &handler{
		assignments, rooms, testEvents, students,
	}
}

var router *echo.Echo

func Mount(e *echo.Echo, h CalendarHandler) {
	router = e
	common.CalendarGroup.GET("/:date", h.Calendar).Name = string(common.AdminCalendar)
	common.DayDetailsGroup.GET("", h.AdminCalendarDetails).Name = string(common.AdminCalendarDetails)
	common.DBDayDetailsGroup.GET("", h.DBCalendarDetails).Name = string(common.DBCalendarDetails)
	common.AssignRoomGroup.GET("", h.ShowAssignRoomForm).Name = string(common.ShowAssignRoomForm)
	common.AssignRoomGroup.POST("", h.AssignRoom).Name = string(common.AssignRoom)
}

func (h handler) Calendar(c echo.Context) error {
	log.SetPrefix("AdminHandler: Calendar()")
	assignments, err := h.assignments.ListAll()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving assignments")
	}
	rooms, err := h.rooms.ListRooms()
	if err != nil {
		return err
	}
	assignmentsMap := assignments.MapForCalendar()
	dateString := c.Param("date")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		log.Println(err)
		return c.String(500, "Error parsing date")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(views.CalendarComponent(date, assignmentsMap, rooms, true, router), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage(views.AdminPageProps{R: router})), c, 200)
}

func (h handler) AdminCalendarDetails(c echo.Context) error {
	dateParam := c.Param("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		log.Println(err)
		return c.String(500, "Error parsing date param")
	}
	assignments, err := h.assignments.ListAll()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving assignments")
	}
	assignmentsMap := assignments.MapForCalendar()
	rooms, err := h.rooms.ListRooms()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving rooms")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(views.DayComponent(date, assignmentsMap[date.Format("2006-01-02")], rooms, true, router), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage(views.AdminPageProps{R: router})), c, 200)
}

func (h handler) DBCalendarDetails(c echo.Context) error {
	dateParam := c.Param("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		log.Println(err)
		return c.String(500, "Error parsing date param")
	}
	assignments, err := h.assignments.ListAll()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving assignments")
	}
	assignmentsMap := assignments.MapForCalendar()
	rooms, err := h.rooms.ListRooms()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving rooms")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(views.DayComponent(date, assignmentsMap[date.Format("2006-01-02")], rooms, false, router), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage(views.AdminPageProps{R: router})), c, 200)
}

func (h handler) ShowAssignRoomForm(c echo.Context) error {
	idParam := c.Param("assignment-id")
	assignmentID, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	assignment, err := h.assignments.GetByAssignmentID(assignmentID)
	log.Println("assignment.ID", assignment.ID)
	if err != nil {
		return c.String(500, err.Error())
	}
	rooms, err := h.rooms.ListRooms()
	if err != nil {
		return c.String(500, err.Error())
	}
	return util.RenderTempl(views.AssignRoomForm(assignment, rooms, router), c, 200)
}

func (h handler) AssignRoom(c echo.Context) error {
	idParam := c.Param("assignment-id")
	assignmentID, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("Error converting assignment ID to int")
		log.Println(err)
		return err
	}
	roomID, err := strconv.Atoi(c.FormValue("room-id"))
	log.Println("room-id: ", c.FormValue("room-id"))
	if err != nil {
		log.Println("Error converting room ID to int")
		return err
	}
	dateParam := c.FormValue("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		log.Println(err)
		return c.String(500, "Error parsing date")
	}
	_, err = h.assignments.UpdateRoom(assignmentID, roomID)
	if err != nil {
		log.Println("Error updating room")
		return err
	}
	assignments, err := h.assignments.ListAll()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving assignments")
	}
	assignmentsMap := assignments.MapForCalendar()
	rooms, err := h.rooms.ListRooms()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving rooms")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(views.DayComponent(date, assignmentsMap[date.Format("2006-01-02")], rooms, true, router), c, 201)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage(views.AdminPageProps{R: router})), c, 200)
}
