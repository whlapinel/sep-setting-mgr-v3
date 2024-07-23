package students

import (
	"errors"
	"log"
	"net/http"
	"sep_setting_mgr/internal/domain/models"
	common "sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/views"
	"sep_setting_mgr/internal/services/classes"
	"sep_setting_mgr/internal/services/students"
	"sep_setting_mgr/internal/util"
	"strconv"

	"github.com/labstack/echo/v4"
)

type StudentsHandler interface {
	// GET /dashboard/classes/:class-id/students
	Students(c echo.Context) error

	// GET /dashboard/classes/:class-id/students/add
	ShowAddStudentForm(c echo.Context) error

	// POST /dashboard/classes/:class-id/students
	CreateStudent(c echo.Context) error

	// GET /dashboard/classes/:class-id/students/:student-id/edit
	ShowEditStudentForm(c echo.Context) error

	// POST /dashboard/students/:student-id/edit
	EditStudent(c echo.Context) error

	// DELETE /students/:student-id
	DeleteStudent(c echo.Context) error
}

type handler struct {
	students students.StudentsService
	classes  classes.ClassesService
}

func NewHandler(students students.StudentsService, classes classes.ClassesService) StudentsHandler {
	return &handler{students, classes}
}

var router *echo.Echo

func Mount(e *echo.Echo, h StudentsHandler) {
	router = e
	common.StudentsGroup.GET("", h.Students).Name = string(common.Students)
	common.StudentsGroup.POST("", h.CreateStudent).Name = string(common.CreateStudent)
	common.StudentsGroup.GET("/add", h.ShowAddStudentForm).Name = string(common.ShowAddStudentForm)
	common.StudentIDGroup.GET("/edit", h.ShowEditStudentForm).Name = string(common.ShowEditStudentForm)
	common.StudentIDGroup.POST("", h.EditStudent).Name = string(common.EditStudent)
	common.StudentIDGroup.DELETE("", h.DeleteStudent).Name = string(common.DeleteStudent)
}

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
	class, err := h.classes.FindClassByID(classID)
	if err != nil {
		return err
	}
	students, err := h.students.ListStudents(classID)
	if err != nil {
		return c.String(500, "Failed to list students. See server logs for details.")
	}
	class.Students = students
	return util.RenderTempl(views.StudentTableComponent(class.Students, class, router), c, 200)
}

func (h handler) CreateStudent(c echo.Context) error {
	log.SetPrefix("Handler: ")
	log.Println("Handler: Adding student")
	firstName := c.FormValue("first-name")
	lastName := c.FormValue("last-name")
	oneOnOne := c.FormValue("one-on-one") == "yes"
	log.Println("First name:", firstName)
	log.Println("Last name:", lastName)
	classID, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		return c.String(400, "Invalid class ID")
	}
	student, err := h.students.AddStudent(firstName, lastName, classID, oneOnOne)
	if err != nil {
		if errors.Is(err, util.ErrNotAssigned) {
			message := "Rooms were full for this event and not all students were assigned to a room. Please contact your administrator."
			util.SetMessage(c, message)
		} else {
			return c.String(500, "Failed to add student. See server logs for details.")
		}
	}
	return util.RenderTempl(views.StudentRowComponent(student, &student.Class, router), c, 201)
}

func (h handler) ShowEditStudentForm(c echo.Context) error {
	log.SetPrefix("Handler: ")
	log.Println("Handler: Showing edit student form")
	studentID, err := strconv.Atoi(c.Param("student-id"))
	if err != nil {
		return c.String(400, "Invalid student ID")
	}
	student, err := h.students.FindStudentByID(studentID)
	if err != nil {
		return c.String(500, "Failed to get student. See server logs for details.")
	}
	switch util.IsHTMX(c) {
	case true:
		return util.RenderTempl(views.AddStudentForm(true, student.Class.ID, student), c, 200)
	default:
		return c.Redirect(303, "/")
	}
}

func (h handler) EditStudent(c echo.Context) error {
	log.SetPrefix("Handler: ")
	log.Println("Handler: Editing student")
	firstName := c.FormValue("first-name")
	lastName := c.FormValue("last-name")
	oneOnOne := c.FormValue("one-on-one") == "yes"
	studentID, err := strconv.Atoi(c.Param("student-id"))
	if err != nil {
		return c.String(400, "Invalid student ID")
	}
	student, err := h.students.UpdateStudent(firstName, lastName, oneOnOne, studentID)
	if err != nil {
		return c.String(500, "Failed to update student. See server logs for details.")
	}
	return util.RenderTempl(views.StudentRowComponent(student, &student.Class, router), c, 200)
}

func (h handler) ShowAddStudentForm(c echo.Context) error {
	classID, err := strconv.Atoi(c.Param("class-id"))
	if err != nil {
		return c.String(400, "Invalid class ID")
	}
	log.SetPrefix("Handler: ")
	log.Println("Handler: Showing add student form")
	switch util.IsHTMX(c) {
	case true:
		return util.RenderTempl(views.AddStudentForm(false, classID, &models.Student{}), c, 200)
	default:
		return c.Redirect(303, "/")
	}
}

func (h handler) DeleteStudent(c echo.Context) error {
	if !util.IsHTMX(c) {
		return c.String(400, "Invalid request")
	}
	studentID, err := strconv.Atoi(c.Param("student-id"))
	if err != nil {
		return c.String(400, "Invalid student ID")
	}
	err = h.students.DeleteStudent(studentID)
	if err != nil {
		return c.String(500, "Failed to delete student. See server logs for details.")
	}
	return c.NoContent(http.StatusOK)
}
