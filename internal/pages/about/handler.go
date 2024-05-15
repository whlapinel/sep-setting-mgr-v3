package about

import (
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type (
	Handler interface {
		// redirect after middleware credential check fails
		AboutHandler(c echo.Context) error
	}
	handler struct {
	}
)

func NewHandler() Handler {
	return &handler{}
}

func Mount(e *echo.Echo, h Handler) {
	e.GET("/about", h.AboutHandler)
}

func (h handler) AboutHandler(c echo.Context) error {
	return util.RenderTempl(AboutPage(), c)
}
