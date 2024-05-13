package classes

import (
	"log"
	"sep_setting_mgr/internal/templates/components"
	"strconv"

	"github.com/labstack/echo"
)

type (
	Handler interface {
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
	classesGroup := e.Group("/hx-classes")
	classesGroup.POST("", h.Create)

	classIDgroup := classesGroup.Group("/:hx-classid")
	classIDgroup.GET("", h.Details)
}

func (h handler) Create(c echo.Context) error {
	log.Println("Handler: Creating class")
	name := c.FormValue("name")
	block, err := strconv.Atoi(c.FormValue("block"))
	if err != nil {
		return c.String(400, "Invalid block")
	}
	class, err := h.service.Add(name, block)
	if err != nil {
		log.Println("Failed to create class:", err)
		return c.String(500, "Failed to create class. Error:"+err.Error())
	}

	switch isHTMX(c) {
	case true:
		err := components.ClassRowComponent(*class).Render(c.Request().Context(), c.Response().Writer)
		if err != nil {
			return c.String(500, "Failed to render class row component. See server logs for details.")
		}
	default:
		return c.Redirect(303, "/")
	}
	return nil
}

func (h handler) Details(e echo.Context) error {
	classID := e.Param("classid")
	log.Println(classID)
	class, err := h.service.FindByID(classID)
	if err != nil {
		return err
	}
	return e.String(200, "Class name: "+class.Name)
}

func isHTMX(e echo.Context) bool {
	// Check for "HX-Request" header
	return e.Request().Header.Get("Hx-Request") != ""
}
