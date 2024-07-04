package domain

import (
	"github.com/labstack/echo/v4"
)

type DashboardService interface {
	// List returns a copy of the todos list
	List(teacherID int) ([]*Class, error)
	AddClass(name string, block int, teacherID int) (*Class, error)
	DeleteClass(classID int) error
	AddStudent(firstName string, lastName string, classID int) (*Student, error)
	DeleteStudent(studentID int) error
	ListStudents(classID int) ([]*Student, error)
	FindClassByID(classID int) (*Class, error)
	ListAllTestEvents(classID int) (*TestEvents, error)
	CreateTestEvent(classID int, testName string, testDate string) (*TestEvent, error)
	DeleteTestEvent(testEventID int) error
}

type DashboardHandler interface {
	// GET /dashboard
	DashboardHandler(c echo.Context) error

	// POST /dashboard/classes
	CreateClass(c echo.Context) error

	// DELETE /dashboard/classes/:class-id
	DeleteClass(c echo.Context) error

	// GET /dashboard/classes/:class-id/students
	Students(c echo.Context) error

	// POST /dashboard/classes/:class-id/students
	AddStudent(c echo.Context) error

	// DELETE /students/:student-id
	DeleteStudent(c echo.Context) error

	// GET /dashboard/classes/:class-id/test-events
	TestEvents(c echo.Context) error

	// POST /dashboard/classes/:class-id/test-events
	CreateTestEvent(c echo.Context) error

	// DELETE /test-events/test-event-id
	DeleteTestEvent(c echo.Context) error
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
