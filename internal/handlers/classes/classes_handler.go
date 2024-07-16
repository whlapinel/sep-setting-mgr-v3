package classes

import (
	"log"
	"net/http"
	"sep_setting_mgr/internal/domain/models"
	common "sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/services/classes"
	"sep_setting_mgr/internal/util"
	"sep_setting_mgr/internal/handlers/views/layouts"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClassesHandler interface {
	// GET /dashboard/classes/add
	ShowAddClassForm(c echo.Context) error

	// GET /dashboard/classes/:class-id/edit
	ShowEditClassForm(c echo.Context) error

	// POST /dashboard/classes
	CreateClass(c echo.Context) error

	// POST /dashboard/classes/:class-id/edit
	EditClass(c echo.Context) error

	// DELETE /dashboard/classes/:class-id
	DeleteClass(c echo.Context) error

	// GET /dashboard/classes
	Classes(c echo.Context) error

	// GET /dashboard/hx-classes
	HxClasses(c echo.Context) error
}

type (
	handler struct {
		classes classes.ClassesService
	}
)

func NewHandler(classes classes.ClassesService) ClassesHandler {
	return &handler{
		classes,
	}
}

var router *echo.Echo

func Mount(e *echo.Echo, h ClassesHandler) {
	router = e
	common.ClassesGroup.GET("", h.Classes).Name = string(common.Classes)
	common.ClassesGroup.GET("/hx-classes", h.HxClasses).Name = string(common.HxClasses)
	common.ClassesGroup.GET("/add", h.ShowAddClassForm).Name = string(common.ShowAddClassForm)
	common.ClassesGroup.POST("", h.CreateClass).Name = string(common.CreateClass)
	common.ClassIDGroup.GET("/edit", h.ShowEditClassForm).Name = string(common.ShowEditClassForm)
	common.ClassIDGroup.POST("/edit", h.EditClass).Name = string(common.EditClass)
	common.ClassIDGroup.DELETE("", h.DeleteClass).Name = string(common.DeleteClass)
}

func (h handler) HxClasses(c echo.Context) error {
	log.SetPrefix("Classes: ")
	if err := c.Get("id"); err == nil {
		log.Println("Failed to get teacher ID.")
		log.Println("c.Get(id): ", c.Get("id"))
		return c.String(500, "Failed to get teacher ID.")
	}
	teacherID := c.Get("id").(int)
	classes, err := h.classes.List(teacherID)
	if err != nil {
		log.Println("Failed to list classes: ", err)
		return c.String(500, "Failed to list classes. See server logs for details.")
	}
	if !util.IsHTMX(c) {
		return c.Redirect(303, router.Reverse(string(common.Classes)))
	}
	return util.RenderTempl(views.ClassesTable(classes, router), c, 200)
}

func (h handler) Classes(c echo.Context) error {
	log.SetPrefix("Classes: ")
	if err := c.Get("id"); err == nil {
		log.Println("Failed to get teacher ID.")
		log.Println("c.Get(id): ", c.Get("id"))
		return c.String(500, "Failed to get teacher ID.")
	}
	teacherID := c.Get("id").(int)
	classes, err := h.classes.List(teacherID)
	if err != nil {
		log.Println("Failed to list classes: ", err)
		return c.String(500, "Failed to list classes. See server logs for details.")
	}
	test := router.Reverse("delete-class", 1)
	log.Println("test: ", test)
	currentUrl := c.Request().Header.Get("HX-Current-URL")
	log.Println("Current URL: ", currentUrl)
	// return util.RenderTempl(components.ClassesTable(classes, router, common.ShowAddClassForm, common.DeleteClass), c, 200)
	if util.IsHTMX(c) {
		return util.RenderTempl(views.DashboardPage(classes, router), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(views.DashboardPage(classes, router)), c, 200)

}

func (h handler) ShowAddClassForm(c echo.Context) error {
	log.SetPrefix("Class Handler: ")
	log.Println("Handler: Showing add class form")
	switch util.IsHTMX(c) {
	case true:
		class := models.Class{}
		return util.RenderTempl(views.AddClassForm(&views.AddClassFormProperties{
			IsEdit:     false,
			Class:      &class,
			AddPostURI: router.Reverse(string(common.CreateClass)),
		}), c, 200)
	default:
		return c.Redirect(303, "/")
	}
}

func (h handler) ShowEditClassForm(c echo.Context) error {
	log.SetPrefix("Class Handler: ")
	log.Println("Handler: Showing edit class form")
	classID, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		return c.String(400, "Invalid class ID")
	}
	class, err := h.classes.FindClassByID(classID)
	if err != nil {
		return c.String(500, "Failed to get class. See server logs for details.")
	}
	switch util.IsHTMX(c) {
	case true:
		return util.RenderTempl(views.AddClassForm(&views.AddClassFormProperties{
			IsEdit:      true,
			Class:       class,
			EditPostURI: router.Reverse("edit-class", class.ID),
			AddPostURI:  router.Reverse("add-class"),
		}), c, 200)
	default:
		return c.Redirect(303, "/")
	}
}
func (h handler) DeleteClass(c echo.Context) error {
	if !util.IsHTMX(c) {
		return c.String(400, "Invalid request")
	}
	classID, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		return c.String(400, "Invalid class ID")
	}
	err = h.classes.DeleteClass(classID)
	if err != nil {
		return c.String(500, "Failed to delete class. See server logs for details.")
	}
	return c.NoContent(http.StatusOK)
}

func (h handler) CreateClass(c echo.Context) error {
	log.Println("Handler: Creating class")
	name := c.FormValue("name")
	block, err := strconv.Atoi(c.FormValue("block"))
	if err != nil {
		return c.String(400, "Invalid block")
	}
	teacherID := c.Get("id").(int)
	log.Println(teacherID)
	class, err := h.classes.AddClass(name, block, teacherID)
	if err != nil {
		log.Println("Failed to create class:", err)
		return c.String(500, "Failed to create class. Error:"+err.Error())
	}

	switch util.IsHTMX(c) {
	case true:

		err := util.RenderTempl(views.ClassRowComponent(class, router), c, 201)
		if err != nil {
			return c.String(500, "Failed to render class row component. See server logs for details.")
		}
	default:
		return c.Redirect(303, "/")
	}
	return nil
}

func (h handler) EditClass(c echo.Context) error {
	log.SetPrefix("Class Handler: ")
	log.Println("Handler: Editing class")
	classID, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		return c.String(400, "Invalid class ID")
	}
	name := c.FormValue("name")
	log.Println("Class ID:", classID)
	log.Println("Name:", name)
	class, err := h.classes.UpdateClass(classID, name)
	if err != nil {
		return c.String(500, "Failed to edit class. See server logs for details.")
	}
	return util.RenderTempl(views.ClassRowComponent(class, router), c, 200)
}
