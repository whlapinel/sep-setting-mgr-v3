package domain

import "github.com/labstack/echo/v4"

type DashboardService interface {
	// List returns a copy of the todos list
	List(teacherID int) ([]*Class, error)
	AddClass(name string, block int, teacherID int) (*Class, error)
	DeleteClass(classID int) error
	AddStudent(firstName string, lastName string, classID int) (*Student, error)
	DeleteStudent(studentID int) error
	ListStudents(classID int) ([]*Student, error)
	FindClassByID(classID int) (*Class, error)
}

type DashboardHandler interface {
	// Dashboard : GET /dashboard
	DashboardHandler(c echo.Context) error

	// Create : POST /dashboard/classes
	CreateClass(c echo.Context) error

	// Delete : DELETE /dashboard/classes/:classID
	DeleteClass(c echo.Context) error

	// Students : GET /dashboard/classes/:classID/students
	Students(c echo.Context) error

	// AddStudent : POST /dashboard/classes/:classID/students
	AddStudent(c echo.Context) error

	// DeleteStudent : DELETE /dashboard/students/:studentID
	DeleteStudent(c echo.Context) error
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
