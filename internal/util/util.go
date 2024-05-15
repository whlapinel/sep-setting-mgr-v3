package util

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func IsHTMX(e echo.Context) bool {
	// Check for "HX-Request" header
	return e.Request().Header.Get("Hx-Request") != ""
}

func RenderTempl(component templ.Component, c echo.Context) error {
	return component.Render(c.Request().Context(), c.Response().Writer)
}
