package calendar

import (
	"log"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/assignments"
	"sep_setting_mgr/internal/services/students"
	testevents "sep_setting_mgr/internal/services/test_events"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type CalendarHandler interface {
	// GET /admin/calendar
	Calendar(c echo.Context) error
}

type handler struct {
	assignments assignments.AssignmentsService
	testEvents  testevents.TestEventsService
	students    students.StudentsService
}

func NewHandler(assignments assignments.AssignmentsService, testEvents testevents.TestEventsService, students students.StudentsService) CalendarHandler {
	return &handler{
		assignments, testEvents, students,
	}
}

func Mount(e *echo.Echo, h CalendarHandler) {
	common.CalendarGroup.GET("", h.Calendar)
}

func (h handler) Calendar(c echo.Context) error {
	log.SetPrefix("AdminHandler: Calendar()")
	assignments, err := h.assignments.ListAll()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving assignments")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(views.AdminCalendarComponent(assignments), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage()), c, 200)
}
