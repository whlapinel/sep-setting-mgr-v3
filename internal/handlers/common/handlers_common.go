package common

import (
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/domain/models"

	"github.com/labstack/echo/v4"
)

type RouteName string

const (
	Rooms                RouteName = "rooms"
	ShowAddClassForm     RouteName = "show-add-class-form"
	DeleteClass          RouteName = "delete-class"
	ShowEditClassForm    RouteName = "show-edit-class-form"
	EditClass            RouteName = "edit-class"
	Classes              RouteName = "classes"
	HxClasses            RouteName = "hx-classes"
	CreateClass          RouteName = "create-class"
	Students             RouteName = "students"
	ShowAddStudentForm   RouteName = "show-add-student-form"
	ShowEditStudentForm  RouteName = "show-edit-student-form"
	DeleteStudent        RouteName = "delete-student"
	EditStudent          RouteName = "edit-student"
	DeleteTestEvent      RouteName = "delete-test-event"
	ShowAddTestEventForm RouteName = "show-add-test-event-form"
	TestEvents           RouteName = "test-events"
	CreateTestEvent      RouteName = "create-test-event"
	CreateStudent        RouteName = "create-student"
	SignupPage           RouteName = "signup-page"
	Signup               RouteName = "signup-post"
	Signout              RouteName = "signout"
	Calendar             RouteName = "calendar"
	ShowAddRoomForm      RouteName = "show-add-room-form"
	CreateRoom           RouteName = "create-room"
	SigninPostRoute      RouteName = "signin-post"
	SigninPage           RouteName = "signin-page"
	Unauthorized         RouteName = "unauthorized"
	Users                RouteName = "users"
)

var (
	// /admin
	AdminGroup *echo.Group

	// /admin/calendar
	CalendarGroup *echo.Group

	// /admin/rooms
	RoomsGroup *echo.Group

	// /dashboard
	DashboardGroup *echo.Group

	// /dashboard/classes
	ClassesGroup *echo.Group

	// /dashboard/classes/:class-id
	ClassIDGroup *echo.Group

	// /dashboard/classes/:class-id/students
	StudentsGroup *echo.Group

	// /dashboard/classes/:class-id/test-events
	TestEventsGroup *echo.Group

	// /dashboard/classes/:class-id/test-events/:test-event-id
	TestEventsIDGroup *echo.Group

	// /dashboard/classes/:class-id/students/:student-id
	StudentIDGroup *echo.Group
)

func CreateGroups(e *echo.Echo, userRepo models.UserRepository) {
	AdminGroup = e.Group("/admin", auth.AddCookieToHeader, auth.JWTMiddleware, auth.GetClaims, auth.Authorization(userRepo))
	RoomsGroup = AdminGroup.Group("/rooms")
	CalendarGroup = AdminGroup.Group("/calendar")

	DashboardGroup = e.Group("/dashboard", auth.AddCookieToHeader, auth.JWTMiddleware, auth.GetClaims)

	ClassesGroup = DashboardGroup.Group("/classes")
	ClassIDGroup = ClassesGroup.Group("/:class-id")

	StudentsGroup = ClassIDGroup.Group("/students")
	StudentIDGroup = StudentsGroup.Group("/:student-id")

	TestEventsGroup = ClassIDGroup.Group("/test-events")
	TestEventsIDGroup = TestEventsGroup.Group("/:test-event-id")
}
