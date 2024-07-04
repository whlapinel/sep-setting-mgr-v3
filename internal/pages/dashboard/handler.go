package dashboard

import (
	"github.com/labstack/echo/v4"

	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/domain"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"
)

type (
	handler struct {
		service domain.DashboardService
	}
)

func NewHandler(svc domain.DashboardService) domain.DashboardHandler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h domain.DashboardHandler) {
	r := e.Group("/dashboard")
	r.Use(auth.AddCookieToHeader)
	r.Use(auth.JWTMiddleware)
	r.Use(auth.GetClaims)
	r.GET("", h.DashboardHandler)
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
	teacherID := c.Get("id").(int)
	classes, err := h.service.List(teacherID)
	if err != nil {
		return c.String(500, "Failed to list classes. See server logs for details.")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(DashboardPage(classes), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(DashboardPage(classes)), c, 200)
}
