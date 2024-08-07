package signup

import (
	"fmt"
	"log"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/signup"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type SignupHandler interface {
	// signup : GET /
	SignUpPage(c echo.Context) error

	// signup : POST /
	Signup(c echo.Context) error
}

type handler struct {
	service signup.SignupService
}

func NewHandler(svc signup.SignupService) SignupHandler {
	return &handler{service: svc}
}

var router *echo.Echo

func Mount(e *echo.Echo, h SignupHandler) {
	router = e
	e.GET("/signup", h.SignUpPage).Name = string(common.SignupPage)
	e.POST("/hx-signup", h.Signup).Name = string(common.Signup)
}

func (h handler) Signup(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	created, err := h.service.CreateUser(email, password)
	if err != nil {
		log.Println("SignUpHandler(): ", err)
		c.String(500, "Error creating user.")
		return err
	}
	if !created {
		fmt.Println("User not created.")
		return echo.ErrInternalServerError
	}
	return c.String(201, "User created.")
}

func (h handler) SignUpPage(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(views.SignUpPage(router), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.SignUpPage(router)), c, 200)
}
