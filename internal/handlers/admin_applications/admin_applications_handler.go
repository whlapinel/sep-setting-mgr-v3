package adminapplications

import (
	"log"
	"sep_setting_mgr/internal/domain/models"
	common "sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/services/applications"
	"sep_setting_mgr/internal/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AdminApplicationsHandler interface {
	Applications(c echo.Context) error

	AdjudicateApplication(c echo.Context) error
}

type handler struct {
	applications applications.ApplicationsService
}

func NewHandler(applications applications.ApplicationsService) AdminApplicationsHandler {
	return &handler{applications}
}

var router *echo.Echo

func Mount(e *echo.Echo, h AdminApplicationsHandler) {
	router = e
	common.AdminApplicationsGroup.GET("", h.Applications).Name = common.Applications.String()
	common.AdminApplicationsGroup.POST("/:app-id/:action", h.AdjudicateApplication).Name = common.AdjudicateApplication.String()
}

func (h handler) Applications(c echo.Context) error {

	apps, err := h.applications.All()
	if err != nil {
		log.Println(err)
		return c.String(500, "Error fetching applications")
	}
	template := views.ApplicationsTable(views.ApplicationsTableProps{
		R:    router,
		Apps: apps,
	})
	if util.IsHTMX(c) {
		return util.RenderTempl(template, c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(template), c, 200)
}

func (h handler) AdjudicateApplication(c echo.Context) error {
	appID, err := strconv.Atoi(c.Param("app-id"))
	if err != nil {
		return c.String(400, "Invalid user ID")
	}
	actionString := c.Param("action")
	action := models.Action(actionString)

	err = h.applications.AdjudicateApplication(appID, action)
	if err != nil {
		return c.String(500, "Error adjudicating application")
	}
	return c.String(200, "Application adjudicated")
}
