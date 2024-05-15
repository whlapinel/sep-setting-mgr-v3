package signin

import (
	"log"
	"net/http"
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/templates/layouts"
	"sep_setting_mgr/internal/util"
	"time"

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
	if util.IsHTMX(c) {
		return util.RenderTempl(SignInPage(), c)
	}
	return util.RenderTempl(layouts.MainLayout(SignInPage()), c)
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
	writeToken(c, t)
	return c.String(200, "Authenticated")
}

func writeToken(c echo.Context, t string) {
	writeCookie(c, t)
	c.Response().Header().Set("Authorization", "Bearer "+t)
}

func writeCookie(c echo.Context, t string) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = t
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(5 * time.Minute)
	log.Println("Setting cookie: ", cookie)
	c.SetCookie(cookie)
}
