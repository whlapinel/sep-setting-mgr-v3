package unauthorized

import (
	"log"
	"net/http"
	"net/url"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"sep_setting_mgr/internal/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UnauthorizedHandler interface {
	// redirect after middleware credential check fails
	Unauthorized(c echo.Context) error
}

type handler struct {
}

func NewHandler() UnauthorizedHandler {
	return &handler{}
}

var router *echo.Echo

func Mount(e *echo.Echo, h UnauthorizedHandler) {
	router = e
	e.GET("/unauthorized", h.Unauthorized).Name = string(common.Unauthorized)
	e.GET("/unauthorized/:page", h.Unauthorized).Name = string(common.UnauthorizedWithPage)
	e.GET("/unauthorized/:page/:reason", h.Unauthorized).Name = string(common.UnauthorizedWithPageAndReason)
	e.GET("/unauthorized/:page/:reason/:user-id", h.Unauthorized).Name = string(common.UnauthorizedWithPageReasonAndUserID)
}

func (h handler) Unauthorized(c echo.Context) error {
	log.SetPrefix("UnauthorizedHandler: Unauthorized()")
	userID, err := strconv.Atoi(c.Param("user-id"))
	if err != nil {
		log.Println("Unauthorized: User ID not found")
		userID = 0
	}
	log.Println("Unauthorized user ID: ", userID)
	page := c.Param("page")
	log.Println("Unauthorized page: ", page)
	reasonEscaped := c.Param("reason")
	reasonUnescaped, err := url.QueryUnescape(reasonEscaped)
	if err != nil {
		log.Println("Error unescaping reason: ", err)
	}
	log.Println("Unauthorized reason: ", reasonUnescaped)
	c.Response().Header().Set("HX-Retarget", "#page")
	c.Response().Header().Set("HX-Reswap", "innerHTML")
	unAuthTemplate := views.UnauthorizedPage(views.UnauthorizedProps{Page: page, Reason: reasonUnescaped, R: router, UserID: userID})
	if util.IsHTMX(c) {
		return util.RenderTempl(unAuthTemplate, c, http.StatusOK)
	}
	return util.RenderTempl(layouts.MainLayout(unAuthTemplate), c, http.StatusOK)
}
