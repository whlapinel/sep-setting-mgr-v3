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

	// if not signed in show alert "you are already signed out"
	if !auth.IsSignedIn(c) {
		return c.String(200, "You are already signed out")
	}
	// delete cookie and render signout page
	auth.WriteToken(c, "")
	return util.RenderTempl(SignedOutPage(), c, 200)
}
