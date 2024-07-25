package applications

import (
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/services/applications"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ApplicationsHandler interface {

	// POST /admin
	ApplyForRole(c echo.Context) error
}

type handler struct {
	applications applications.ApplicationsService
}

func NewHandler(applications applications.ApplicationsService) ApplicationsHandler {
	return &handler{applications}
}


var router *echo.Echo

func Mount(e *echo.Echo, h ApplicationsHandler) {
	router = e
	common.ApplicationsGroup.POST("/:userID", h.ApplyForRole).Name = string(common.ApplyForRole)
}

func (h handler) ApplyForRole(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.String(400, "Invalid user ID")
	}
	role := c.FormValue("role")
	err = h.applications.ApplyForRole(userID, role)
	if err != nil {
		return c.String(500, "Error applying for admin")
	}
	return c.Redirect(302, "/admin")
}
