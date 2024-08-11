package calendar

import (
	"errors"
	"log"
	"sep_setting_mgr/internal/domain/models"
	common "sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/views"
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

	// GET /admin/calendar/:date/:block
	// GET /admin/calendar/:date/:block/:room-id
	AdminCalendarDetails(c echo.Context) error

	// GET /admin/calendar/auto-assign/:date/:block
	AutoAssign(c echo.Context) error

	// POST /admin/calendar/auto-assign/:date/:block
	AcceptAutoAssignments(c echo.Context) error

	// GET /dashboard/calendar/:date/:block
	// GET /dashboard/calendar/:date/:block/:room-id
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
	common.AdminCalendarGroup.GET("/:date", h.Calendar).Name = common.AdminCalendar.String()
	common.AdminDayDetailsGroup.GET("", h.AdminCalendarDetails).Name = common.AdminCalendarDetails.String()
	common.AdminCalendarGroup.GET("/:date/:block", h.AdminCalendarDetails).Name = common.AdminCalDetailsWithoutRoomID.String()
	common.DBDayDetailsGroup.GET("", h.DBCalendarDetails).Name = common.DBCalendarDetails.String()
	common.DBCalendarGroup.GET("/:date/:block", h.DBCalendarDetails).Name = common.DBCalDetailsWithoutRoomID.String()
	common.AutoAssignGroup.GET("", h.AutoAssign).Name = common.AutoAssign.String()
	common.AutoAssignGroup.POST("", h.AcceptAutoAssignments).Name = common.AcceptAutoAssignments.String()
	common.AssignRoomGroup.GET("", h.ShowAssignRoomForm).Name = common.ShowAssignRoomForm.String()
	common.AssignRoomGroup.POST("", h.AssignRoom).Name = common.AssignRoom.String()

}

func (h handler) Calendar(c echo.Context) error {
	log.SetPrefix("AdminHandler: Calendar()")
	assignments, err := h.assignments.All()
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
	date, err := time.Parse(views.ButtonDateFormat, dateString)
	if err != nil {
		log.Println(err)
		return c.String(500, "Error parsing date")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(views.CalendarComponent(date, assignmentsMap, rooms, true, router), c, 200)
	}
	return c.Redirect(303, router.Reverse(common.AdminPage.String()))
}

func (h handler) AdminCalendarDetails(c echo.Context) error {
	date, block, roomID, err := h.parseCalendarParams(c)
	if err != nil {
		log.Println(err)
		return err
	}
	assignments, err := h.assignments.All()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving assignments")
	}
	assignmentsMap := assignments.MapForCalendar()
	rooms, err := h.rooms.ListRooms()
	if err != nil {
		log.Println(err)
		return err
	}
	var selectedRoom *models.Room
	for _, room := range rooms {
		log.Println("room.ID: ", room.ID)
		if room.ID == roomID {
			selectedRoom = room
		}
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(views.DayComponent(date, assignmentsMap[date.Format("2006-01-02")], block, rooms, selectedRoom, true, router), c, 200)
	}
	return c.Redirect(303, router.Reverse(common.AdminPage.String()))
}

func (h handler) AutoAssign(c echo.Context) error {
	log.Println("AutoAssign() handler running...")
	dateParam := c.Param("date")
	date, err := time.Parse("2006-01-02", dateParam)
	log.Println("date: ", date)
	if err != nil {
		log.Println(err)
		return err
	}
	block, err := strconv.Atoi(c.Param("block"))
	log.Println("block: ", block)
	if err != nil {
		log.Println(err)
		return err
	}
	autoAssignments, err := h.assignments.CreateAutoAssignments(date, block)
	log.Println("len(autoAssignments): ", len(autoAssignments))
	if err != nil {
		log.Println(err)
		return err
	}
	return util.RenderTempl(views.AutoAssignments(autoAssignments, date, block, router), c, 200)
}

func (h handler) parseCalendarParams(c echo.Context) (time.Time, int, int, error) {
	dateParam := c.Param("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		log.Println(err)
		return time.Time{}, 0, 0, err
	}
	block, err := strconv.Atoi(c.Param("block"))
	if err != nil {
		log.Println(err)
		return time.Time{}, 0, 0, err
	}
	if block == 0 {
		log.Println("Invalid block")
		return time.Time{}, 0, 0, errors.New("invalid block")
	}
	var roomID int
	if c.Param("room-id") == "" {
		roomID = 0
	} else {
		roomID, err = strconv.Atoi(c.Param("room-id"))
		if err != nil {
			log.Println(err)
			return time.Time{}, 0, 0, err
		}
	}
	return date, block, roomID, nil
}

func (h handler) DBCalendarDetails(c echo.Context) error {
	date, block, roomID, err := h.parseCalendarParams(c)
	if err != nil {
		log.Println(err)
		return err
	}
	assignments, err := h.assignments.All()
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
	var selectedRoom *models.Room
	for _, room := range rooms {
		if room.ID == roomID {
			log.Println("room.ID: ", room.ID)
			selectedRoom = room
		}
	}

	if util.IsHTMX(c) {
		return util.RenderTempl(views.DayComponent(date, assignmentsMap[date.Format("2006-01-02")], block, rooms, selectedRoom, false, router), c, 200)
	}
	return c.Redirect(303, router.Reverse(common.AdminPage.String()))
}

func (h handler) ShowAssignRoomForm(c echo.Context) error {
	idParam := c.Param("assignment-id")
	assignmentID, err := strconv.Atoi(idParam)
	if err != nil {
		return err
	}
	assignment, err := h.assignments.FindByID(assignmentID)
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
	err = h.assignments.UpdateRoom(assignmentID, roomID)
	if err != nil {
		log.Println("Error updating room")
		return err
	}
	assignments, err := h.assignments.All()
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
		return util.RenderTempl(views.CalendarComponent(date, assignmentsMap, rooms, true, router), c, 201)
	}
	return c.Redirect(303, router.Reverse(common.AdminPage.String()))
}

func (h handler) AcceptAutoAssignments(c echo.Context) error {
	log.SetPrefix("AdminHandler: AcceptAutoAssignments()")
	dateParam := c.Param("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		log.Println(err)
		return err
	}
	block, err := strconv.Atoi(c.Param("block"))
	if err != nil {
		log.Println(err)
		return err
	}
	err = c.Request().ParseForm()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Form: ", c.Request().Form)
	assignmentIDs := c.Request().Form["assignment-id"]
	roomIDs := c.Request().Form["room-id"]
	log.Println("len(assignmentIDs): ", len(assignmentIDs))
	for i, id := range assignmentIDs {
		assignmentID, err := strconv.Atoi(id)
		if err != nil {
			log.Println(err)
			return err
		}
		roomID, err := strconv.Atoi(roomIDs[i])
		if err != nil {
			log.Println(err)
			return err
		}
		err = h.assignments.AcceptAutoAssignment(assignmentID, roomID)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return util.RenderTempl(views.AutoAssignmentsAccepted(date, block, router), c, 201)
}
