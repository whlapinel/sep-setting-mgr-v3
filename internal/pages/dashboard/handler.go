package dashboard

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"

	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/templates/components"
	"sep_setting_mgr/internal/templates/layouts"
	"sep_setting_mgr/internal/util"
)

type (
	Handler interface {
		// Dashboard : GET /
		DashboardHandler(c echo.Context) error

		// Create : POST /classes
		Create(c echo.Context) error

		// Details : GET /classes/:classID
		Details(c echo.Context) error
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

	classesGroup := r.Group("/hx-classes")
	classesGroup.POST("", h.Create)
	classIDgroup := classesGroup.Group("/:hx-classid")
	classIDgroup.GET("", h.Details)
}

func (h handler) DashboardHandler(c echo.Context) error {
	classes, err := h.service.List()
	if err != nil {
		return c.String(500, "Failed to list classes. See server logs for details.")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(DashboardPage(classes), c)
	}
	return util.RenderTempl(layouts.MainLayout(DashboardPage(classes)), c)
}

func (h handler) Create(c echo.Context) error {
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
		err := util.RenderTempl(components.ClassRowComponent(*class), c)
		if err != nil {
			return c.String(500, "Failed to render class row component. See server logs for details.")
		}
	default:
		return c.Redirect(303, "/")
	}
	return nil
}

func (h handler) Details(c echo.Context) error {
	classID := c.Param("classid")
	log.Println(classID)
	class, err := h.service.FindClassByID(classID)
	if err != nil {
		return err
	}
	return c.String(200, "Class name: "+class.Name)
}
