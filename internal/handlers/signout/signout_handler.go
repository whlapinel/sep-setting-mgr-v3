package signout

import (
	"sep_setting_mgr/internal/auth"
	common "sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type SignoutHandler interface {
	// signout : POST /
	Signout(c echo.Context) error
}

type handler struct {
}

func NewHandler() SignoutHandler {
	return &handler{}
}

func Mount(e *echo.Echo, h SignoutHandler) {
	e.POST("/hx-signout", h.Signout).Name = string(common.Signout)
}

func (h handler) Signout(c echo.Context) error {

	// if not signed in show alert "you are already signed out"
	if !auth.IsSignedIn(c) {
		return c.String(200, "You are already signed out")
	}
	// delete cookie and render signout page
	auth.WriteToken(c, "")
	return util.RenderTempl(views.SignedOutPage(), c, 200)
}
