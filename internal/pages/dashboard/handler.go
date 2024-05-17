package dashboard

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"

	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/components"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"
)

type (
	Handler interface {
		// Dashboard : GET /
		DashboardHandler(c echo.Context) error

		// Create : POST /classes
		CreateClass(c echo.Context) error

		// Students : GET /classes/:classID/students
		Students(c echo.Context) error
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

	classesGroup := r.Group("/classes")
	classesGroup.POST("", h.CreateClass)
	classIDgroup := classesGroup.Group("/:classid")
	classIDgroup.GET("/students", h.Students)
}

func (h handler) DashboardHandler(c echo.Context) error {
	teacherID := c.Get("id").(int)
	classes, err := h.service.List(teacherID)
	if err != nil {
		return c.String(500, "Failed to list classes. See server logs for details.")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(DashboardPage(classes), c)
	}
	return util.RenderTempl(layouts.MainLayout(DashboardPage(classes)), c)
}

func (h handler) CreateClass(c echo.Context) error {
	log.Println("Handler: Creating class")
	name := c.FormValue("name")
	block, err := strconv.Atoi(c.FormValue("block"))
	if err != nil {
		return c.String(400, "Invalid block")
	}
	teacherID := c.Get("id").(int)
	log.Println(teacherID)
	class, err := h.service.AddClass(name, block, teacherID)
	if err != nil {
		log.Println("Failed to create class:", err)
		return c.String(500, "Failed to create class. Error:"+err.Error())
	}

	switch util.IsHTMX(c) {
	case true:
		err := util.RenderTempl(components.ClassRowComponent(class), c)
		if err != nil {
			return c.String(500, "Failed to render class row component. See server logs for details.")
		}
	default:
		return c.Redirect(303, "/")
	}
	return nil
}

func (h handler) Students(c echo.Context) error {
	if !util.IsHTMX(c) {
		return c.String(400, "Invalid request")
	}
	classID := c.Param("classid")
	log.Println("classID: ", classID)
	class, err := h.service.FindClassByID(classID)
	if err != nil {
		return err
	}
	return util.RenderTempl(StudentTableComponent(class.Students), c)
}
