package routes

import (
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

func RegisterTestEventRoutes(e *echo.Echo) {
	var testEvents domain.TestEvents
	for i := 0; i < 10; i++ {
		testEvents = append(testEvents, domain.NewTestEvent("Test "+strconv.Itoa(i), i, time.Now()))
	}
	e.GET("/", func(c echo.Context) error {
		return layouts.MainLayout(pages.CalendarPage(&testEvents)).Render(c.Request().Context(), c.Response().Writer)
	})
	e.POST("/add", func(c echo.Context) error {
		newDate, err := time.Parse("2006-01-02", c.FormValue("date"))
		if err != nil {
			return c.String(400, "Invalid date")
		}
		testEvent := domain.NewTestEvent(c.FormValue("name"), 0, newDate)
		log.Println("Adding new item")
		log.Println(c.FormValue("name"))
		
		return components.TestEventRow(testEvent).Render(c.Request().Context(), c.Response().Writer)
	})
}