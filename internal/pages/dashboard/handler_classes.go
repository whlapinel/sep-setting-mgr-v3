package dashboard

import (
	"log"
	"net/http"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/pages/dashboard/components"
	"sep_setting_mgr/internal/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h handler) ShowAddClassForm(c echo.Context) error {
	log.SetPrefix("Class Handler: ")
	log.Println("Handler: Showing add class form")
	switch util.IsHTMX(c) {
	case true:
		class := models.Class{}
		return util.RenderTempl(components.AddClassForm(&components.AddClassFormProperties{
			IsEdit:     false,
			Class:      &class,
			AddPostURI: router.Reverse(string(addClass)),
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
	class, err := h.service.FindClassByID(classID)
	if err != nil {
		return c.String(500, "Failed to get class. See server logs for details.")
	}
	switch util.IsHTMX(c) {
	case true:
		return util.RenderTempl(components.AddClassForm(&components.AddClassFormProperties{
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
	err = h.service.DeleteClass(classID)
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
	class, err := h.service.AddClass(name, block, teacherID)
	if err != nil {
		log.Println("Failed to create class:", err)
		return c.String(500, "Failed to create class. Error:"+err.Error())
	}

	switch util.IsHTMX(c) {
	case true:

		err := util.RenderTempl(components.ClassRowComponent(class, router, deleteClass), c, 201)
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
	class, err := h.service.UpdateClass(classID, name)
	if err != nil {
		return c.String(500, "Failed to edit class. See server logs for details.")
	}
	return util.RenderTempl(components.ClassRowComponent(class, router, deleteClass), c, 200)
}
