package signout

import (
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type (
	Handler interface {
		// signout : POST /
		HxHandleSignOut(c echo.Context) error
	}

	handler struct {
	}
)

func NewHandler() Handler {
	return &handler{}
}

func Mount(e *echo.Echo, h Handler) {
	e.POST("/hx-signout", h.HxHandleSignOut)
}

func (h handler) HxHandleSignOut(c echo.Context) error {
	// delete cookie and render signout page
	auth.WriteToken(c, "")
	return util.RenderTempl(SignedOutPage(), c)
}
