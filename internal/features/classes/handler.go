package classes

import (
	"sep_setting_mgr/internal/templates/components"
	"strconv"

	"github.com/labstack/echo"
)

type (
	Handler interface {
		// Create : POST /classes
		Create(c echo.Context) error
	}

	handler struct {
		service Service
	}
)

func NewHandler(svc Service) Handler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h Handler) {
	e.POST("/classes", h.Create)
}

func (h handler) Create(c echo.Context) error {
	name := c.FormValue("name")
	block, err := strconv.Atoi(c.FormValue("block"))
	if err != nil {
		return c.String(400, "Invalid block")
	}
	class, err := h.service.Add(name, block)
	if err != nil {
		return c.String(500, "Failed to create class. See server logs for details.")
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

func isHTMX(e echo.Context) bool {
	// Check for "HX-Request" header
	return e.Request().Header.Get("Hx-Request") != ""
}
