package signin

import (
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type (
	Handler interface {
		// signin : GET /
		SignInHandler(e echo.Context) error
		// signin : POST /
		HxHandleSignin(e echo.Context) error
	}

	handler struct {
		service Service
	}
)

func NewHandler(svc Service) Handler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h Handler) {
	e.GET("/signin", h.SignInHandler)
	e.POST("/hx-signin", h.HxHandleSignin)
}

func (h handler) SignInHandler(c echo.Context) error {
	isSignedIn := auth.IsSignedIn(c)
	if util.IsHTMX(c) {
		return util.RenderTempl(SignInPage(isSignedIn), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(SignInPage(isSignedIn)), c, 200)
}

func (h handler) HxHandleSignin(c echo.Context) error {
	if !(util.IsHTMX(c)) {
		c.Redirect(303, "/signin")
	}
	email := c.FormValue("email")
	password := c.FormValue("password")
	isAuth, err := h.service.VerifyCredentials(email, password)
	if !isAuth || err != nil {
		return c.String(401, "Invalid credentials")
	}
	id := h.service.GetUserID(email)
	t, err := auth.IssueToken(email, id)
	if err != nil {
		return c.String(500, "Failed to issue token")
	}
	auth.WriteToken(c, t)
	return c.Redirect(303, "/dashboard")
}
