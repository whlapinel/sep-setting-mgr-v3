package applications

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

type ApplicationsHandler interface {
	// GET /applications
	ApplicationsPage(c echo.Context) error

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
	common.ApplicationsGroup.GET("", h.ApplicationsPage).Name = common.ApplicationsPage.String()
	common.ApplicationsGroup.POST("/:userID/:role", h.ApplyForRole).Name = common.ApplyForRole.String()

}

func (h handler) ApplyForRole(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return util.RenderTempl(views.AppliedFailure(views.AppliedFailureProps{Role: "unknown", Reason: "error parsing id"}), c, 200)
	}
	roleString := c.Param("role")
	role := models.Role(roleString)
	err = h.applications.ApplyForRole(userID, role)
	if err != nil {
		return util.RenderTempl(views.AppliedFailure(views.AppliedFailureProps{Role: role, Reason: "error creating application"}), c, 200)
	}
	return util.RenderTempl(views.AppliedSuccess(views.AppliedSuccessProps{Role: role}), c, 200)
}

func (h handler) ApplicationsPage(c echo.Context) error {
	userID, ok := c.Get("id").(int)
	if !ok {
		return c.String(400, "Invalid user ID")
	}
	hasTeacherRole, err := h.applications.HasRole(userID, "teacher")
	if err != nil {
		log.Println(err)
		return c.String(500, "Error checking for teacher role")
	}
	hasAdminRole, err := h.applications.HasRole(userID, "admin")
	if err != nil {
		log.Println(err)
		return c.String(500, "Error checking for admin role")
	}
	appliedTeacher, err := h.applications.HasApplied(userID, "teacher")
	if err != nil {
		log.Println(err)
		return c.String(500, "Error checking if applied for teacher role")
	}
	appliedAdmin, err := h.applications.HasApplied(userID, "admin")
	if err != nil {
		log.Println(err)
		return c.String(500, "Error checking if applied for admin role")
	}
	log.Println("Applied for teacher:", appliedTeacher)
	log.Println("Applied for admin:", appliedAdmin)
	template := views.ApplicationsPage(views.ApplicationsPageProps{
		R:                 router,
		UserID:            userID,
		AppliedForTeacher: appliedTeacher,
		AppliedForAdmin:   appliedAdmin,
		HasTeacherRole:    hasTeacherRole,
		HasAdminRole:      hasAdminRole,
	})
	if util.IsHTMX(c) {
		return util.RenderTempl(template, c, 200)
	}
	user, err := models.NewUser(c.Get("first").(string), c.Get("last").(string), c.Get("email").(string), c.Get("picture").(string))
	if err != nil {
		return err
	}
	return util.RenderTempl(layouts.MainLayout(template, user), c, 200)

}
