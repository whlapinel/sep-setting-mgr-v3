package common

import (
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/domain/models"

	"github.com/labstack/echo/v4"
)

const (
	GoogleSignup                        string = "google-signup"
	GoogleSignin                        string = "google-signin"
	Registration                        string = "registration"
	Rooms                               string = "rooms"
	DeleteRoom                          string = "delete-room"
	ShowEditRoomForm                    string = "show-edit-room-form"
	ShowAddRoomForm                     string = "show-add-room-form"
	EditRoom                            string = "edit-room"
	ShowEditUserForm                    string = "show-edit-user-form"
	EditUser                            string = "edit-user"
	Users                               string = "users"
	DeleteUser                          string = "delete-user"
	ShowAddClassForm                    string = "show-add-class-form"
	DeleteClass                         string = "delete-class"
	ShowEditClassForm                   string = "show-edit-class-form"
	EditClass                           string = "edit-class"
	Classes                             string = "classes"
	HxClasses                           string = "hx-classes"
	CreateClass                         string = "create-class"
	Students                            string = "students"
	ShowAddStudentForm                  string = "show-add-student-form"
	ShowEditStudentForm                 string = "show-edit-student-form"
	DeleteStudent                       string = "delete-student"
	EditStudent                         string = "edit-student"
	DeleteTestEvent                     string = "delete-test-event"
	ShowAddTestEventForm                string = "show-add-test-event-form"
	ShowEditTestEventForm               string = "show-edit-test-event-form"
	EditTestEvent                       string = "edit-test-event"
	TestEvents                          string = "test-events"
	CreateTestEvent                     string = "create-test-event"
	CreateStudent                       string = "create-student"
	SignupPage                          string = "signup-page"
	Signup                              string = "signup-post"
	Signout                             string = "signout"
	DashboardCalendar                   string = "dashboard-calendar"
	DBCalendarDetails                   string = "dashboard-calendar-details"
	AdminCalendar                       string = "admin-calendar"
	AdminCalendarDetails                string = "admin-calendar-details"
	ShowAssignRoomForm                  string = "show-assign-room-form"
	AssignRoom                          string = "assign-room"
	CreateRoom                          string = "create-room"
	SigninPostRoute                     string = "signin-post"
	SigninPage                          string = "signin-page"
	Unauthorized                        string = "unauthorized"
	UnauthorizedWithPage                string = "unauthorized-with-page"
	UnauthorizedWithPageAndReason       string = "unauthorized-with-page-and-reason"
	UnauthorizedWithPageReasonAndUserID string = "unauthorized-with-page-reason-and-user-id"
	ApplicationsPage                    string = "applications-page"
	ApplyForRole                        string = "apply-for-role"
	Applications                        string = "admin-applications"     // GET /admin/applications
	AdjudicateApplication               string = "adjudicate-application" // POST /admin/applications/:app-id/:action
)

var (
	// /admin
	AdminGroup *echo.Group

	// /admin/applications
	AdminApplicationsGroup *echo.Group

	// /applications
	ApplicationsGroup *echo.Group

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

	AdminGroup = e.Group("/admin", auth.AddCookieToHeader, auth.JWTMiddleware, auth.GetClaims, auth.Authorization(userRepo, models.AdminRole))

	UsersGroup = AdminGroup.Group("/users")
	UserIDGroup = UsersGroup.Group("/:user-id")

	RoomsGroup = AdminGroup.Group("/rooms")
	RoomsIDGroup = RoomsGroup.Group("/:room-id")

	CalendarGroup = AdminGroup.Group("/calendar")
	DayDetailsGroup = CalendarGroup.Group("/:date/details")

	AssignRoomGroup = CalendarGroup.Group("/assign-room/:assignment-id")
	AdminApplicationsGroup = AdminGroup.Group("/applications")

	ApplicationsGroup = e.Group("/applications", auth.AddCookieToHeader, auth.JWTMiddleware, auth.GetClaims)

	DashboardGroup = e.Group("/dashboard", auth.AddCookieToHeader, auth.JWTMiddleware, auth.GetClaims, auth.Authorization(userRepo, models.TeacherRole))
	DBCalendarGroup = DashboardGroup.Group("/calendar")
	DBDayDetailsGroup = DBCalendarGroup.Group(":date/details")

	ClassesGroup = DashboardGroup.Group("/classes")
	ClassIDGroup = ClassesGroup.Group("/:class-id")

	StudentsGroup = ClassIDGroup.Group("/students")
	StudentIDGroup = StudentsGroup.Group("/:student-id")

	TestEventsGroup = ClassIDGroup.Group("/test-events")
	TestEventsIDGroup = TestEventsGroup.Group("/:test-event-id")
}
