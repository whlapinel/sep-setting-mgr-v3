package signin

import (
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/signin"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type SigninHandler interface {
	// signin : GET /
	SignInHandler(e echo.Context) error
	// signin : POST /
	HxSignin(e echo.Context) error
}

type handler struct {
	service signin.SigninService
}

func NewHandler(svc signin.SigninService) SigninHandler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h SigninHandler) {
	e.GET("/signin", h.SignInHandler).Name = string(common.SigninPage)
	e.POST("/hx-signin", h.HxSignin).Name = string(common.SigninPostRoute)
}

func (h handler) SignInHandler(c echo.Context) error {
	isSignedIn := auth.IsSignedIn(c)
	if util.IsHTMX(c) {
		return util.RenderTempl(views.SignInPage(isSignedIn), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.SignInPage(isSignedIn)), c, 200)
}

func (h handler) HxSignin(c echo.Context) error {
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
