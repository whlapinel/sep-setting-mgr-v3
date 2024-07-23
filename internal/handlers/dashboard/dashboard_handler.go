package dashboard

import (
	"log"
	"net/http"
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/assignments"
	"sep_setting_mgr/internal/services/classes"
	"sep_setting_mgr/internal/services/rooms"
	"sep_setting_mgr/internal/services/students"
	testevents "sep_setting_mgr/internal/services/test_events"
	"sep_setting_mgr/internal/util"
	"time"

	"github.com/labstack/echo/v4"
)

type DashboardHandler interface {
	// GET /dashboard
	Redirect(c echo.Context) error

	// GET /dashboard/calendar
	DashboardCalendar(c echo.Context) error
}

type handler struct {
	classesService classes.ClassesService
	assignments    assignments.AssignmentsService
	testEvents     testevents.TestEventsService
	students       students.StudentsService
	rooms          rooms.RoomsService
}

func NewHandler(classes classes.ClassesService, assignments assignments.AssignmentsService, testEvents testevents.TestEventsService, students students.StudentsService, rooms rooms.RoomsService) DashboardHandler {
	return &handler{
		classes,
		assignments,
		testEvents,
		students,
		rooms,
	}
}

var router *echo.Echo

func Mount(e *echo.Echo, h DashboardHandler) {
	router = e
	common.DashboardGroup.Use(auth.AddCookieToHeader)
	common.DashboardGroup.Use(auth.JWTMiddleware)
	common.DashboardGroup.Use(auth.GetClaims)
	common.DashboardGroup.GET("", h.Redirect)
	common.DBCalendarGroup.GET("", h.DashboardCalendar).Name = string(common.DashboardCalendar)
}

func (h handler) Redirect(c echo.Context) error {
	return c.Redirect(303, router.Reverse(string(common.Classes)))
}

func (h handler) DashboardCalendar(c echo.Context) error {
	log.SetPrefix("ShowCalendar: ")
	teacherID := c.Get("id").(int)
	assignments, err := h.assignments.GetByTeacherID(teacherID)
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving assignments")
	}
	log.Println("len(assignments): ", len(assignments))
	assignmentsMap := assignments.MapForCalendar()
	rooms, err := h.rooms.ListRooms()
	if err != nil {
		return err
	}
	date := time.Now()
	calendar := views.CalendarComponent(date, assignmentsMap, rooms, false, router)
	if util.IsHTMX(c) {
		return util.RenderTempl(calendar, c, http.StatusOK)
	}
	return util.RenderTempl(layouts.MainLayout(calendar), c, http.StatusOK)
}
