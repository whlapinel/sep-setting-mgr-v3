package pages

import (
	"sep_setting_mgr/internal/domain/models"
	"time"

	"github.com/labstack/echo/v4"
)

type DashboardService interface {
	// classes service
	List(teacherID int) ([]*models.Class, error)
	AddClass(name string, block int, teacherID int) (*models.Class, error)
	DeleteClass(classID int) error
	UpdateClass(classID int, name string) (*models.Class, error)
	FindClassByID(classID int) (*models.Class, error)

	// students  service
	AddStudent(firstName string, lastName string, classID int, oneOneOne bool) (*models.Student, error)
	UpdateStudent(firstName string, lastName string, oneOnOne bool, studentID int) (*models.Student, error)
	FindStudentByID(studentID int) (*models.Student, error)
	DeleteStudent(studentID int) error
	ListStudents(classID int) ([]*models.Student, error)

	// test events
	ListAllTestEvents(classID int) (models.TestEvents, error)
	CreateTestEvent(classID int, testName string, testDate string) (*models.TestEvent, error)
	DeleteTestEvent(testEventID int) error
	FindTestEventByID(testEventID int) (*models.TestEvent, error)

	// assignments
	GetAssignments(teacherID int, start, end time.Time) (models.Assignments, error)
	CreateAssignment(student *models.Student, testEvent *models.TestEvent) (*models.Assignment, error)
	DeleteAssignments(studentID int) error
}

type DashboardHandler interface {
	// GET /dashboard
	Redirect(c echo.Context) error

	// GET /classes
	Classes(c echo.Context) error

	// GET /dashboard/classes
	DashboardHandler(c echo.Context) error

	// GET /dashboard/calendar
	ShowCalendar(c echo.Context) error

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

	// GET /dashboard/classes/:class-id/students
	Students(c echo.Context) error

	// GET /dashboard/classes/:class-id/students/add
	ShowAddStudentForm(c echo.Context) error

	// POST /dashboard/classes/:class-id/students
	AddStudent(c echo.Context) error

	// GET /dashboard/classes/:class-id/students/:student-id/edit
	ShowEditStudentForm(c echo.Context) error

	// POST /dashboard/students/:student-id/edit
	EditStudent(c echo.Context) error

	// DELETE /students/:student-id
	DeleteStudent(c echo.Context) error

	// GET /dashboard/classes/:class-id/test-events
	TestEvents(c echo.Context) error

	// GET /dashboard/classes/:class-id/test-events/add
	ShowAddTestEventForm(c echo.Context) error

	// GET /dashboard/classes/:class-id/test-events/:test-event-id/edit
	ShowEditTestEventForm(c echo.Context) error

	// POST /dashboard/classes/:class-id/test-events
	CreateTestEvent(c echo.Context) error

	// DELETE /test-events/test-event-id
	DeleteTestEvent(c echo.Context) error
}

type AdminHandler interface {
	// GET /admin
	AdminHandler(c echo.Context) error

	// GET /admin/rooms
	Rooms(c echo.Context) error

	// GET /admin/rooms/add
	ShowAddRoomForm(c echo.Context) error

	// POST /admin/rooms
	CreateRoom(c echo.Context) error

	// GET /admin/calendar
	Calendar(c echo.Context) error

	// GET /admin/users
	Users(c echo.Context) error

	// Middleware for /admin/* routes
	Authorization(next echo.HandlerFunc) echo.HandlerFunc
}

type AdminService interface {
	GetAllAssignments() (models.Assignments, error)
	ListUsers() ([]*models.User, error)
	AddRoom(*models.Room) (id int, err error)
	ListRooms() ([]*models.Room, error)
	FindRoomByID(id int) (*models.Room, error)
	IsAdmin(userID int) bool
}

type HomeHandler interface {
	// Dashboard : GET /
	Home(e echo.Context) error
}

type SigninHandler interface {
	// signin : GET /
	SignInHandler(e echo.Context) error
	// signin : POST /
	HxHandleSignin(e echo.Context) error
}

type SigninService interface {
	VerifyCredentials(email string, password string) (bool, error)
	GetUserID(email string) int
}

type SignoutHandler interface {
	// signout : POST /
	HxHandleSignOut(c echo.Context) error
}

type SignupHandler interface {
	// signin : GET /
	SignUpHandler(c echo.Context) error
	// signin : POST /
	HxHandleSignUp(c echo.Context) error
}

type SignupService interface {
	CreateUser(username string, password string) (bool, error)
}

type UnauthorizedHandler interface {
	// redirect after middleware credential check fails
	UnauthorizedHandler(c echo.Context) error
}
