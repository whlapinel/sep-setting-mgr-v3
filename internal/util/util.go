package util

import (
	"errors"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var ErrNotAssigned = errors.New("one or more students not assigned to a room")

func SetMessage(c echo.Context, message string) {
	c.Response().Header().Set("HX-Trigger", "{\"showMessage\":\""+message+"\"}")
}

func IsHTMX(e echo.Context) bool {
	// Check for "HX-Request" header
	return e.Request().Header.Get("Hx-Request") != ""
}

func RenderTempl(component templ.Component, c echo.Context, statusCode int) error {
	c.Response().WriteHeader(statusCode)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func ParseDate(date string) (*time.Time, error) {
	// Parse date string
	// Return time.Time and error
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	return &parsedDate, nil
}
