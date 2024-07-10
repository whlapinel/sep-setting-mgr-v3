package dashboard

import (
	"log"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/pages"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/pages/dashboard/components"
	"sep_setting_mgr/internal/util"
)

type (
	handler struct {
		service pages.DashboardService
	}
)

func NewHandler(svc pages.DashboardService) pages.DashboardHandler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h pages.DashboardHandler) {
	r := e.Group("/dashboard")
	r.Use(auth.AddCookieToHeader)
	r.Use(auth.JWTMiddleware)
	r.Use(auth.GetClaims)
	r.GET("", h.Redirect)
	r.GET("/calendar", h.ShowCalendar)
	r.DELETE("/students/:student-id", h.DeleteStudent)
	r.DELETE("/test-events/:test-event-id", h.DeleteTestEvent)
	classesGroup := r.Group("/classes")
	classesGroup.GET("", h.DashboardHandler)
	classesGroup.POST("", h.CreateClass)
	classIDgroup := classesGroup.Group("/:class-id")
	classIDgroup.DELETE("", h.DeleteClass)
	classIDgroup.GET("/students", h.Students)
	classIDgroup.GET("/test-events", h.TestEvents)
	classIDgroup.POST("/test-events", h.CreateTestEvent)
	classIDgroup.POST("/students", h.AddStudent)
}

func (h handler) Redirect(c echo.Context) error {
	return c.Redirect(303, "/dashboard/classes")
}

func (h handler) DashboardHandler(c echo.Context) error {
	log.SetPrefix("DashboardHandler: ")
	teacherID := c.Get("id").(int)
	// if the current page is dashboard, don't return entire page, just the classes component
	prevUrl := c.Request().Header.Get("HX-Current-URL")
	host := c.Request().Host
	prevPath := strings.ReplaceAll(prevUrl, "http://"+host, "")
	log.Println("prevPath: ", prevPath)
	classes, err := h.service.List(teacherID)
	if err != nil {
		log.Println("Failed to list classes: ", err)
		return c.String(500, "Failed to list classes. See server logs for details.")
	}
	if strings.Contains(prevPath, "/dashboard") {
		log.Println("Returning classes table only")
		return util.RenderTempl(components.ClassesTable(classes), c, 200)
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(components.DashboardPage(classes), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.DashboardPage(classes)), c, 200)
}

func (h handler) ShowCalendar(c echo.Context) error {
	log.SetPrefix("ShowCalendar: ")
	teacherID := c.Get("id").(int)
	classes, err := h.service.List(teacherID)
	if err != nil {
		return err
	}
	var assignments models.Assignments
	for _, class := range classes {
		eventAssignments, err := h.service.GetAssignments(class.ID, time.Now(), time.Now().AddDate(0, 1, 0))
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
