package signup

import (
	"fmt"
	"log"
	"sep_setting_mgr/internal/templates/layouts"
	"sep_setting_mgr/internal/util"

	"github.com/labstack/echo/v4"
)

type (
	Handler interface {
		// signin : GET /
		SignUpHandler(c echo.Context) error
		// signin : POST /
		HxHandleSignUp(c echo.Context) error
	}

	handler struct {
		service Service
	}
)

func NewHandler(svc Service) Handler {
	return &handler{service: svc}
}

func Mount(e *echo.Echo, h Handler) {
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
		return util.RenderTempl(SignUpPage(), c)
	}
	return util.RenderTempl(layouts.MainLayout(SignUpPage()), c)
}
