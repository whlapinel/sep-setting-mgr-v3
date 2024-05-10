package teacher_dashboard

import (
	"github.com/labstack/echo"

	"sep_setting_mgr/internal/templates/layouts"
	"sep_setting_mgr/internal/templates/pages"
)

type (
	Handler interface {
		// Dashboard : GET /
		Dashboard(e echo.Context) error
	}

	handler struct {
		service Service
	}
)

func NewHandler(svc Service) Handler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h Handler) {
	e.GET("/", h.Dashboard)
}

func (h handler) Dashboard(c echo.Context) error {
	classes, err := h.service.List()
	if err != nil {
		return c.String(500, "Failed to list classes. See server logs for details.")
	}
	return layouts.MainLayout(pages.Dashboard(&classes)).Render(c.Request().Context(), c.Response().Writer)
}
