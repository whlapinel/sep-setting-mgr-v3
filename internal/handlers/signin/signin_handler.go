package signin

import (
	"log"
	"sep_setting_mgr/internal/auth"
	common "sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/signin"
	"sep_setting_mgr/internal/util"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type SigninHandler interface {
	// GET /signin
	SignInPage(e echo.Context) error

	// POST /signin
	GoogleSignin(e echo.Context) error
}

type handler struct {
	service signin.SigninService
}

func NewHandler(svc signin.SigninService) SigninHandler {
	return &handler{service: svc}
}

var router *echo.Echo

func Mount(e *echo.Echo, h SigninHandler) {
	router = e
	e.GET("/signin", h.SignInPage).Name = string(common.SigninPage)
	e.POST("/signin", h.GoogleSignin).Name = string(common.GoogleSignin)
}

func (h handler) SignInPage(c echo.Context) error {
	isSignedIn := auth.IsSignedIn(c)
	if util.IsHTMX(c) {
		return util.RenderTempl(views.SignInPage(isSignedIn, false, router, ""), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.SignInPage(isSignedIn, false, router, "")), c, 200)
}

func (h handler) GoogleSignin(c echo.Context) error {

	payload, err := auth.GoogleAuth(c)
	if err != nil {
		log.Println(err)
		return c.String(500, "Failed to authenticate")
	}
	email := payload.Claims["email"].(string)
	// check for user in database
	id, err := h.service.GetUserID(email)
	if err != nil {
		log.Println(err)
		return c.String(500, err.Error())
	}
	t, err := auth.IssueToken(email, id)
	if err != nil {
		return c.String(500, "Failed to issue token")
	}
	auth.WriteToken(c, t)
	expirationTime := time.Now().Add(auth.SessionLifeSpan).UnixMilli()
	expyString := strconv.Itoa(int(expirationTime))
	return util.RenderTempl(layouts.MainLayout(views.SignInPage(true, true, router, expyString)), c, 200)
}
