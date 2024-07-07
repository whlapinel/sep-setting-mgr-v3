package constants

type RouteSegment string

type Param string

const (
	// Dashboard
	Dashboard RouteSegment = "dashboard"
	// Classes
	Classes RouteSegment = "classes"
	// Calendar
	Calendar RouteSegment = "calendar"
	// Students
	Students RouteSegment = "students"
	// TestEvents
	TestEvents RouteSegment = "test-events"
	// ClassID
	ClassID Param = "class-id"
	// StudentID
	StudentID Param = "student-id"
	// TestEventID
	TestEventID Param = "test-event-id"

	// Admin
	Admin RouteSegment = "admin"
	// Home
	Home RouteSegment = "home"
	// Signin
	Signin RouteSegment = "signin"
	// Signout
	Signout RouteSegment = "signout"
	// Signup
	Signup RouteSegment = "signup"
	// Unauthorized
	Unauthorized RouteSegment = "unauthorized"
)

func Root(r RouteSegment) string {
	return "/" + string(r)
}

func (s RouteSegment) AddSegment(child RouteSegment) string {
	return "/" + string(s) + "/" + string(child)
}

func (s RouteSegment) AddParam(param Param) string {
	return string(s) + "/:" + string(param)
}
