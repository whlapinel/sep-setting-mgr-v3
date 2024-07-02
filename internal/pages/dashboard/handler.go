package dashboard

import (
	"github.com/labstack/echo/v4"

	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"
)

type (
	Handler interface {
		// Dashboard : GET /dashboard
		DashboardHandler(c echo.Context) error

		// Create : POST /dashboard/classes
		CreateClass(c echo.Context) error

		// Delete : DELETE /dashboard/classes/:classID
		DeleteClass(c echo.Context) error

		// Students : GET /dashboard/classes/:classID/students
		Students(c echo.Context) error

		// AddStudent : POST /dashboard/classes/:classID/students
		AddStudent(c echo.Context) error

		// DeleteStudent : DELETE /dashboard/students/:studentID
		DeleteStudent(c echo.Context) error
	}

	handler struct {
		service Service
	}
)

func NewHandler(svc Service) Handler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h Handler) {
	r := e.Group("/dashboard")
	r.Use(auth.AddCookieToHeader)
	r.Use(auth.JWTMiddleware)
	r.Use(auth.GetClaims)
	r.GET("", h.DashboardHandler)
	r.DELETE("/students/:studentid", h.DeleteStudent)
	classesGroup := r.Group("/classes")
	classesGroup.POST("", h.CreateClass)
	classIDgroup := classesGroup.Group("/:classid")
	classIDgroup.DELETE("", h.DeleteClass)
	classIDgroup.GET("/students", h.Students)
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
