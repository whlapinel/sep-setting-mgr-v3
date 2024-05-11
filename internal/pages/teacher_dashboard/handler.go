package teacher_dashboard

import (
	"github.com/labstack/echo"

	"sep_setting_mgr/internal/templates/layouts"
)

type (
	Handler interface {
		// Dashboard : GET /
		DashboardHandler(e echo.Context) error
	}

	handler struct {
		service Service
	}
)

func NewHandler(svc Service) Handler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h Handler) {
	e.GET("/dashboard", h.DashboardHandler)
}

func (h handler) DashboardHandler(c echo.Context) error {
	classes, err := h.service.List()
	if err != nil {
		return c.String(500, "Failed to list classes. See server logs for details.")
	}
	return layouts.MainLayout(DashboardPage(&classes)).Render(c.Request().Context(), c.Response().Writer)
}
