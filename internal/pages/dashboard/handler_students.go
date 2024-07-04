package dashboard

import (
	"log"
	"net/http"
	"sep_setting_mgr/internal/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h handler) Students(c echo.Context) error {
	if !util.IsHTMX(c) {
		return c.String(400, "Invalid request")
	}
	idParam := c.Param("class-id")
	classID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.String(400, "Invalid class ID")
	}
	log.Println("classID: ", classID)
	class, err := h.service.FindClassByID(classID)
	if err != nil {
		return err
	}
	students, err := h.service.ListStudents(classID)
	if err != nil {
		return c.String(500, "Failed to list students. See server logs for details.")
	}
	class.Students = students
	return util.RenderTempl(StudentTableComponent(class.Students, classID), c, 200)
}

func (h handler) AddStudent(c echo.Context) error {
	log.Println("Handler: Adding student")
	firstName := c.FormValue("first-name")
	lastName := c.FormValue("last-name")
	log.Println("First name:", firstName)
	log.Println("Last name:", lastName)
	classID, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		return c.String(400, "Invalid class ID")
	}
	student, err := h.service.AddStudent(firstName, lastName, classID)
	if err != nil {
		return c.String(500, "Failed to add student. See server logs for details.")
	}
	return util.RenderTempl(StudentRowComponent(student), c, 201)
}

func (h handler) DeleteStudent(c echo.Context) error {
	if !util.IsHTMX(c) {
		return c.String(400, "Invalid request")
	}
	studentID, err := strconv.Atoi(c.Param("student-id"))
	if err != nil {
		return c.String(400, "Invalid student ID")
	}
	err = h.service.DeleteStudent(studentID)
	if err != nil {
		return c.String(500, "Failed to delete student. See server logs for details.")
	}
	return c.NoContent(http.StatusOK)
}
