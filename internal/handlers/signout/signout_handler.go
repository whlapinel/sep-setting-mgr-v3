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
	// delete cookie and render signout page
	auth.WriteToken(c, "")
	c.Response().Header().Set("Hx-Trigger", "userSignout")
	return util.RenderTempl(views.SignedOutPage(), c, 200)
}
