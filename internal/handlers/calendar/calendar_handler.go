package calendar

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/services"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type CalendarHandler interface {
	// GET /admin/calendar
	Calendar(c echo.Context) error
}

type handler struct {
	service services.AssignmentsService
}

func NewHandler(service services.AssignmentsService) CalendarHandler {
	return &handler{
		service: service,
	}
}

func Mount(e *echo.Echo, h CalendarHandler) {
	common.CalendarGroup.GET("", h.Calendar)
}

func (h handler) Calendar(c echo.Context) error {
	var assignments models.Assignments
	assignments, err := h.service.GetAllAssignments()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error retrieving assignments")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(views.AdminCalendarComponent(assignments), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.AdminPage()), c, 200)
}
