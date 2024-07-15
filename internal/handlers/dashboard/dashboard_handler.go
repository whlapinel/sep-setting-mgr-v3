package dashboard

import (
	"log"
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/components"
	"sep_setting_mgr/internal/services/assignments"
	"sep_setting_mgr/internal/services/classes"
	"sep_setting_mgr/internal/util"
	"time"

	"github.com/labstack/echo/v4"
)

type DashboardHandler interface {
	// GET /dashboard
	Redirect(c echo.Context) error

	// GET /dashboard/calendar
	ShowCalendar(c echo.Context) error
}

type handler struct {
	classesService     classes.ClassesService
	assignmentsService assignments.AssignmentsService
}

func NewHandler(classes classes.ClassesService, assignments assignments.AssignmentsService) DashboardHandler {
	return &handler{
		classes,
		assignments,
	}
}

var router *echo.Echo

func Mount(e *echo.Echo, h DashboardHandler) {
	router = e
	common.DashboardGroup.Use(auth.AddCookieToHeader)
	common.DashboardGroup.Use(auth.JWTMiddleware)
	common.DashboardGroup.Use(auth.GetClaims)
	common.DashboardGroup.GET("", h.Redirect)
	common.DashboardGroup.GET("/calendar", h.ShowCalendar).Name = string(common.Calendar)
}

func (h handler) Redirect(c echo.Context) error {
	return c.Redirect(303, router.Reverse(string(common.Classes)))
}

func (h handler) ShowCalendar(c echo.Context) error {
	log.SetPrefix("ShowCalendar: ")
	teacherID := c.Get("id").(int)
	classes, err := h.classesService.List(teacherID)
	if err != nil {
		return err
	}
	var assignments models.Assignments
	for _, class := range classes {
		eventAssignments, err := h.assignmentsService.GetAssignments(class.ID, time.Now(), time.Now().AddDate(0, 1, 0))
		log.Println("eventAssignments: ", eventAssignments)
		log.Println("len(eventAssignments): ", len(eventAssignments))
		if err != nil {
			log.Println("Failed to get assignments: ", err)
			return c.String(500, "Failed to get assignments. See server logs for details.")
		}
		assignments = append(assignments, eventAssignments...)
	}
	log.Println("len(assignments): ", len(assignments))
	calendar := util.RenderTempl(components.CalendarComponent(assignments), c, 200)
	if util.IsHTMX(c) {
		return calendar
	}
	return calendar
}
