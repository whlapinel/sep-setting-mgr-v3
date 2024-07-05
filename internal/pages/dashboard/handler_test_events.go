package dashboard

import (
	"log"
	"sep_setting_mgr/internal/pages/dashboard/components"
	"sep_setting_mgr/internal/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h handler) TestEvents(c echo.Context) error {
	if !util.IsHTMX(c) {
		return c.String(400, "Invalid request")
	}
	classID, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		return c.String(400, "Invalid class ID")
	}
	testEvents, err := h.service.ListAllTestEvents(classID)
	if err != nil {
		return c.String(500, "Failed to list test events. See server logs for details.")
	}
	return util.RenderTempl(components.TestEventsTableComponent(testEvents, classID), c, 200)
}

func (h handler) CreateTestEvent(c echo.Context) error {
	log.SetPrefix("Handler: ")
	log.Println("Creating test event")
	log.Println("Class ID: ", c.Param("class-id"))
	classID, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		return c.String(400, "Invalid class ID")
	}
	testName := c.FormValue("test-name")
	testDate := c.FormValue("test-date")
	testEvent, err := h.service.CreateTestEvent(classID, testName, testDate)
	if err != nil {
		log.Println("Failed to create test event:", err)
		return c.String(500, "Failed to create test event. See server logs for details.")
	}
	return util.RenderTempl(components.TestEventRowComponent(testEvent), c, 201)
}

func (h handler) DeleteTestEvent(c echo.Context) error {
	if !util.IsHTMX(c) {
		return c.String(400, "Invalid request")
	}
	testEventID, err := strconv.Atoi(c.Param("test-event-id"))
	if err != nil {
		return c.String(400, "Invalid test event ID")
	}
	err = h.service.DeleteTestEvent(testEventID)
	if err != nil {
		return c.String(500, "Failed to delete test event. See server logs for details.")
	}
	return c.NoContent(200)
}
