package dashboard

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"

	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/domain/models"
	"sep_setting_mgr/internal/domain/pages"
	"sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/layouts"
	"sep_setting_mgr/internal/pages/dashboard/components"
	"sep_setting_mgr/internal/util"
)

type (
	handler struct {
		service pages.DashboardService
	}
)

func NewHandler(svc pages.DashboardService) pages.DashboardHandler {
	return &handler{service: svc}
}

var router *echo.Echo
var showAddClassFormRoute *echo.Route

const (
	deleteClass      components.RouteName = "delete-class"
	editClassForm    components.RouteName = "edit-class-form"
	editClass        components.RouteName = "edit-class"
	classes          components.RouteName = "classes"
	addClass         components.RouteName = "add-class"
	students         components.RouteName = "students"
	addStudentForm   components.RouteName = "add-student-form"
	editStudentForm  components.RouteName = "edit-student-form"
	deleteStudent    components.RouteName = "delete-student"
	editStudent      components.RouteName = "edit-student"
	deleteTestEvent  components.RouteName = "delete-test-event"
	addTestEventForm components.RouteName = "add-test-event-form"
	testEvents       components.RouteName = "test-events"
	addTestEvent     components.RouteName = "add-test-event"
	addStudent       components.RouteName = "add-student"
)

func Mount(e *echo.Echo, h pages.DashboardHandler) {
	router = e
	r := e.Group("/dashboard")
	r.Use(auth.AddCookieToHeader)
	r.Use(auth.JWTMiddleware)
	r.Use(auth.GetClaims)
	r.GET("", h.Redirect)
	r.GET("/hx-classes", h.Classes).Name = string(classes)
	r.GET("/calendar", h.ShowCalendar)
	classesGroup := r.Group("/classes")
	showAddClassFormRoute = classesGroup.GET("/add", h.ShowAddClassForm)
	classesGroup.GET("", h.DashboardHandler).Name = string(classes)
	classesGroup.POST("", h.CreateClass).Name = string(addClass)
	classIDgroup := classesGroup.Group("/:class-id")
	classIDgroup.GET("/edit", h.ShowEditClassForm).Name = string(editClassForm)
	classIDgroup.POST("/edit", h.EditClass).Name = string(editClass)
	classIDgroup.DELETE("", h.DeleteClass).Name = string(deleteClass)

	r.DELETE("/students/:student-id", h.DeleteStudent).Name = string(deleteStudent)
	r.POST("/students/:student-id", h.EditStudent).Name = string(editStudent)
	classIDgroup.GET("/students", h.Students).Name = string(students)
	classIDgroup.GET("/students/add", h.ShowAddStudentForm).Name = string(addStudentForm)
	studentIDgroup := classIDgroup.Group("/students/:student-id")
	studentIDgroup.GET("/edit", h.ShowEditStudentForm).Name = string(editStudentForm)
	classIDgroup.POST("/students", h.AddStudent).Name = string(addStudent)

	classIDgroup.GET("/test-events/add", h.ShowAddTestEventForm).Name = string(addTestEventForm)
	classIDgroup.GET("/test-events", h.TestEvents).Name = string(testEvents)
	classIDgroup.POST("/test-events", h.CreateTestEvent).Name = string(addTestEvent)
	r.DELETE("/test-events/:test-event-id", h.DeleteTestEvent).Name = string(deleteTestEvent)

}

func (h handler) Redirect(c echo.Context) error {
	return c.Redirect(303, "/dashboard/classes")
}

func (h handler) DashboardHandler(c echo.Context) error {
	log.SetPrefix("DashboardHandler: ")
	teacherID := c.Get("id").(int)
	classes, err := h.service.List(teacherID)
	if err != nil {
		log.Println("Failed to list classes: ", err)
		return c.String(500, "Failed to list classes. See server logs for details.")
	}
	if util.IsHTMX(c) {
		return util.RenderTempl(components.DashboardPage(classes, router, common.ShowAddClassForm, deleteClass), c, 200)
	}
	return util.RenderTempl(layouts.MainLayout(components.DashboardPage(classes, router, common.ShowAddClassForm, deleteClass)), c, 200)
}

func (h handler) Classes(c echo.Context) error {
	log.SetPrefix("Classes: ")
	teacherID := c.Get("id").(int)
	classes, err := h.service.List(teacherID)
	if err != nil {
		log.Println("Failed to list classes: ", err)
		return c.String(500, "Failed to list classes. See server logs for details.")
	}
	test := router.Reverse("delete-class", 1)
	log.Println("test: ", test)
	return util.RenderTempl(components.ClassesTable(classes, router, common.ShowAddClassForm, deleteClass), c, 200)
}

func (h handler) ShowCalendar(c echo.Context) error {
	log.SetPrefix("ShowCalendar: ")
	teacherID := c.Get("id").(int)
	classes, err := h.service.List(teacherID)
	if err != nil {
		return err
	}
	var assignments models.Assignments
	for _, class := range classes {
		eventAssignments, err := h.service.GetAssignments(class.ID, time.Now(), time.Now().AddDate(0, 1, 0))
		log.Println("eventAssignments: ", eventAssignments)
		log.Println("len(eventAssignments): ", len(eventAssignments))
		if err != nil {
			log.Println("Failed to get assignments: ", err)
			return c.String(500, "Failed to get assignments. See server logs for details.")
		}
		assignments = append(assignments, eventAssignments...)
	}
	log.Println("len(assignments): ", len(assignments))
	calendar := util.RenderTempl(components.CalendarComponent(assignments), c, 200)
	if util.IsHTMX(c) {
		return calendar
	}
	return calendar
}
