package dashboard

import (
	"log"

	"github.com/labstack/echo/v4"

	"sep_setting_mgr/internal/auth"
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
	r.GET("", h.DashboardHandler)
	r.GET("/calendar", h.ShowCalendar)
	r.DELETE("/students/:student-id", h.DeleteStudent)
	r.DELETE("/test-events/:test-event-id", h.DeleteTestEvent)
	classesGroup := r.Group("/classes")
	classesGroup.POST("", h.CreateClass)
	classIDgroup := classesGroup.Group("/:class-id")
	classIDgroup.DELETE("", h.DeleteClass)
	classIDgroup.GET("/students", h.Students)
	classIDgroup.GET("/test-events", h.TestEvents)
	classIDgroup.POST("/test-events", h.CreateTestEvent)
	classIDgroup.POST("/students", h.AddStudent)
}

func (h handler) DashboardHandler(c echo.Context) error {
	log.SetPrefix("DashboardHandler: ")
	teacherID := c.Get("id").(int)
	classes, err := h.service.List(teacherID)
	if err != nil {
		log.Println("Failed to list classes: ", err)
		return c.String(500, "Failed to list classes. See server logs for details.")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(components.DashboardPage(classes), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.DashboardPage(classes)), c, 200)
}

func (h handler) ShowCalendar(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(components.CalendarComponent(), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.CalendarComponent()), c, 200)
}
