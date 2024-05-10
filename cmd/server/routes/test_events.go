package routes

import (
	"fmt"
	"log"
	"sep_setting_mgr/internal/domain"
	"sep_setting_mgr/internal/templates/components"
	"sep_setting_mgr/internal/templates/layouts"
	"sep_setting_mgr/internal/templates/pages"
	"strconv"
	"time"

	// import echo
	"github.com/labstack/echo"
)

func RegisterTestEventRoutes(e *echo.Echo, svc domain.TestEventService) {
	e.GET("/", func(c echo.Context) error {
		testEvents := svc.ListAll()
		if c.Request().Header.Get("Hx-Target") == "page" {
			log.Println("rendering page")
			return pages.CalendarPage(testEvents).Render(c.Request().Context(), c.Response().Writer)
		}
		return layouts.MainLayout(pages.CalendarPage(testEvents)).Render(c.Request().Context(), c.Response().Writer)
	})
	e.POST("/add", func(c echo.Context) error {
		newDate, err := time.Parse("2006-01-02", c.FormValue("date"))
		if err != nil {
			return c.String(400, "Invalid date")
		}
		block, err := strconv.Atoi(c.FormValue("block"))
		if err != nil {
			return c.String(400, "Invalid block")
		}
		// new code
		testEvent, err := svc.RegisterNewTestEvent(c.FormValue("name"), domain.Class{}, newDate, block)
		if err != nil {
			log.Println("RegisterNewTestEvent() error: ", err)
			return c.String(500, "Failed to regiser user. See server logs for details.")
		}
		fmt.Println("New testEvent registered: ", testEvent)
		testEvents := svc.ListAll()
		// return components.TestEventRow(testEvent).Render(c.Request().Context(), c.Response().Writer)
		return components.TestEventListComponent(testEvents).Render(c.Request().Context(), c.Response().Writer)
	})
}
