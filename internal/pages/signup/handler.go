package signup

import (
	"fmt"
	"log"
	"sep_setting_mgr/internal/domain"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service domain.SignupService
}

func NewHandler(svc domain.SignupService) domain.SignupHandler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h domain.SignupHandler) {
	e.GET("/signup", h.SignUpHandler)
	e.POST("/hx-signup", h.HxHandleSignUp)
}

func (h handler) HxHandleSignUp(c echo.Context) error {
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

func (h handler) SignUpHandler(c echo.Context) error {
	if util.IsHTMX(c) {
		return util.RenderTempl(SignUpPage(), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(SignUpPage()), c, 200)
}
