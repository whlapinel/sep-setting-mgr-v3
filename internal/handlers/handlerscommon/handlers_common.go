package handlerscommon

import (
	"sep_setting_mgr/internal/auth"
	"sep_setting_mgr/internal/domain/models"

	"github.com/labstack/echo/v4"
)

type RouteName string

func (r RouteName) String() string {
	return string(r)
}

const (
	AboutPage RouteName = "GET /about"

	AdminPage RouteName = "GET /admin"

	// auth routes
	GoogleSignup          RouteName = "google-signup"
	GoogleSignin          RouteName = "google-signin"
	Registration          RouteName = "registration"
	Rooms                 RouteName = "rooms"
	DeleteRoom            RouteName = "delete-room"
	ShowEditRoomForm      RouteName = "show-edit-room-form"
	ShowAddRoomForm       RouteName = "show-add-room-form"
	EditRoom              RouteName = "edit-room"
	ShowEditUserForm      RouteName = "show-edit-user-form"
	EditUser              RouteName = "edit-user"
	Users                 RouteName = "users"
	DeleteUser            RouteName = "delete-user"
	Dashboard             RouteName = "GET /dashboard"
	ShowAddClassForm      RouteName = "GET /dashboard/classes/add"
	DeleteClass           RouteName = "DELETE /dashboard/classes/:class-id"
	ShowEditClassForm     RouteName = "GET /dashboard/classes/:class-id/edit"
	EditClass             RouteName = "POST /dashboard/classes/:class-id/edit"
	Classes               RouteName = "GET /dashboard/classes"
	HxClasses             RouteName = "GET /dashboard/classes/hx-classes"
	CreateClass           RouteName = "POST /dashboard/classes"
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
	Signup                RouteName = "GET /signup"
	Signout               RouteName = "signout"
	DashboardCalendar     RouteName = "GET /dashboard/calendar"
	DBCalendarDetails     RouteName = "GET /dashboard/calendar/:date/details"
	AdminCalendar         RouteName = "GET /admin/calendar/:date"
	AdminCalendarDetails  RouteName = "GET /admin/calendar/:date/details"
	ShowAssignRoomForm    RouteName = "show-assign-room-form"
	AssignRoom            RouteName = "assign-room"
	CreateRoom            RouteName = "create-room"
	PromoteRoom           RouteName = "POST /admin/rooms/:room-id/promote"

	RefreshToken RouteName = "POST /refresh-token"

	SigninPage                          RouteName = "signin-page"
	Unauthorized                        RouteName = "unauthorized"
	UnauthorizedWithPage                RouteName = "unauthorized-with-page"
	UnauthorizedWithPageAndReason       RouteName = "unauthorized-with-page-and-reason"
	UnauthorizedWithPageReasonAndUserID RouteName = "unauthorized-with-page-reason-and-user-id"
	ApplicationsPage                    RouteName = "applications-page"
	ApplyForRole                        RouteName = "apply-for-role"
	Applications                        RouteName = "GET /admin/applications"
	AdjudicateApplication               RouteName = "POST /admin/applications/:app-id/:action"
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
