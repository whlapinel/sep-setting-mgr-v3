package common

import (
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/domain/models"

	"github.com/labstack/echo/v4"
)

type RouteName string

const (
	GoogleSignup          RouteName = "google-signup"
	GoogleSignin          RouteName = "google-signin"
	Rooms                 RouteName = "rooms"
	DeleteRoom            RouteName = "delete-room"
	ShowEditRoomForm      RouteName = "show-edit-room-form"
	ShowAddRoomForm       RouteName = "show-add-room-form"
	EditRoom              RouteName = "edit-room"
	ShowEditUserForm      RouteName = "show-edit-user-form"
	EditUser              RouteName = "edit-user"
	Users                 RouteName = "users"
	DeleteUser            RouteName = "delete-user"
	ShowAddClassForm      RouteName = "show-add-class-form"
	DeleteClass           RouteName = "delete-class"
	ShowEditClassForm     RouteName = "show-edit-class-form"
	EditClass             RouteName = "edit-class"
	Classes               RouteName = "classes"
	HxClasses             RouteName = "hx-classes"
	CreateClass           RouteName = "create-class"
	Students              RouteName = "students"
	ShowAddStudentForm    RouteName = "show-add-student-form"
	ShowEditStudentForm   RouteName = "show-edit-student-form"
	DeleteStudent         RouteName = "delete-student"
	EditStudent           RouteName = "edit-student"
	DeleteTestEvent       RouteName = "delete-test-event"
	ShowAddTestEventForm  RouteName = "show-add-test-event-form"
	ShowEditTestEventForm RouteName = "show-edit-test-event-form"
	EditTestEvent         RouteName = "edit-test-event"
	TestEvents            RouteName = "test-events"
	CreateTestEvent       RouteName = "create-test-event"
	CreateStudent         RouteName = "create-student"
	SignupPage            RouteName = "signup-page"
	Signup                RouteName = "signup-post"
	Signout               RouteName = "signout"
	DashboardCalendar     RouteName = "dashboard-calendar"
	DBCalendarDetails     RouteName = "dashboard-calendar-details"
	AdminCalendar         RouteName = "admin-calendar"
	AdminCalendarDetails  RouteName = "admin-calendar-details"
	ShowAssignRoomForm    RouteName = "show-assign-room-form"
	AssignRoom            RouteName = "assign-room"
	CreateRoom            RouteName = "create-room"
	SigninPostRoute       RouteName = "signin-post"
	SigninPage            RouteName = "signin-page"
	Unauthorized          RouteName = "unauthorized"
)

var (

	// /admin
	AdminGroup *echo.Group

	// /admin/calendar
	CalendarGroup *echo.Group

	// /admin/calendar/:date/details
	DayDetailsGroup *echo.Group

	// /dashboard/calendar/:date/details
	DBDayDetailsGroup *echo.Group

	// /admin/calendar/assign-room/:assignment-id
	AssignRoomGroup *echo.Group

	// /admin/rooms
	RoomsGroup *echo.Group

	// /admin/rooms/:room-id
	RoomsIDGroup *echo.Group

	// /dashboard
	DashboardGroup *echo.Group

	// /dashboard/calendar
	DBCalendarGroup *echo.Group

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

	// /admin/users
	UsersGroup *echo.Group

	// /admin/users/:user-id
	UserIDGroup *echo.Group
)

func CreateGroups(e *echo.Echo, userRepo models.UserRepository) {

	AdminGroup = e.Group("/admin", auth.AddCookieToHeader, auth.JWTMiddleware, auth.GetClaims, auth.Authorization(userRepo))

	UsersGroup = AdminGroup.Group("/users")
	UserIDGroup = UsersGroup.Group("/:user-id")

	RoomsGroup = AdminGroup.Group("/rooms")
	RoomsIDGroup = RoomsGroup.Group("/:room-id")

	CalendarGroup = AdminGroup.Group("/calendar")
	DayDetailsGroup = CalendarGroup.Group("/:date/details")

	AssignRoomGroup = CalendarGroup.Group("/assign-room/:assignment-id")

	DashboardGroup = e.Group("/dashboard", auth.AddCookieToHeader, auth.JWTMiddleware, auth.GetClaims)
	DBCalendarGroup = DashboardGroup.Group("/calendar")
	DBDayDetailsGroup = DBCalendarGroup.Group(":date/details")

	ClassesGroup = DashboardGroup.Group("/classes")
	ClassIDGroup = ClassesGroup.Group("/:class-id")

	StudentsGroup = ClassIDGroup.Group("/students")
	StudentIDGroup = StudentsGroup.Group("/:student-id")

	TestEventsGroup = ClassIDGroup.Group("/test-events")
	TestEventsIDGroup = TestEventsGroup.Group("/:test-event-id")
}
