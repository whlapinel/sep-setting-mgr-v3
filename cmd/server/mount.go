package main

import (
	"database/sql"
	"sep_setting_mgr/internal/database"
	"sep_setting_mgr/internal/domain/services"
	"sep_setting_mgr/internal/handlers/about"
	"sep_setting_mgr/internal/handlers/admin"
	"sep_setting_mgr/internal/handlers/calendar"
	"sep_setting_mgr/internal/handlers/classes"
	common "sep_setting_mgr/internal/handlers/common"
	"sep_setting_mgr/internal/handlers/dashboard"
	"sep_setting_mgr/internal/handlers/home"
	"sep_setting_mgr/internal/handlers/rooms"
	"sep_setting_mgr/internal/handlers/signin"
	"sep_setting_mgr/internal/handlers/signout"
	"sep_setting_mgr/internal/handlers/signup"
	"sep_setting_mgr/internal/handlers/students"
	testevents "sep_setting_mgr/internal/handlers/test_events"
	"sep_setting_mgr/internal/handlers/unauthorized"
	"sep_setting_mgr/internal/handlers/users"
	adminService "sep_setting_mgr/internal/services/admin"
	"sep_setting_mgr/internal/services/assignments"
	classesService "sep_setting_mgr/internal/services/classes"
	roomsService "sep_setting_mgr/internal/services/rooms"
	signinService "sep_setting_mgr/internal/services/signin"
	signupService "sep_setting_mgr/internal/services/signup"
	studentsService "sep_setting_mgr/internal/services/students"
	testEventsService "sep_setting_mgr/internal/services/test_events"
	usersService "sep_setting_mgr/internal/services/users"

	"github.com/labstack/echo/v4"
)

func MountHandlers(e *echo.Echo, db *sql.DB) error {
	// initialize repositories
	usersRepo := database.NewUsersRepo(db)
	classesRepo := database.NewClassesRepo(db)
	studentsRepo := database.NewStudentsRepo(db)
	testEventsRepo := database.NewTestEventsRepo(db)
	roomsRepo := database.NewRoomsRepo(db)
	assignmentsRepo := database.NewAssignmentsRepo(db)

	// initialize domain services
	assignmentsDomainService := services.NewAssignmentsService(assignmentsRepo, roomsRepo, testEventsRepo, classesRepo, studentsRepo)

	// initialize application services
	assignmentAppService := assignments.NewService(assignmentsRepo, roomsRepo, testEventsRepo)
	usersService := usersService.NewService(usersRepo, roomsRepo, assignmentsRepo)
	classesService := classesService.NewService(classesRepo, studentsRepo)
	testEventsService := testEventsService.NewService(testEventsRepo, studentsRepo, classesRepo, assignmentsRepo, roomsRepo, *assignmentsDomainService)
	studentsService := studentsService.NewService(studentsRepo, classesRepo, *assignmentsDomainService)
	roomsService := roomsService.NewService(roomsRepo)
	signupService := signupService.NewService(usersRepo)
	signinService := signinService.NewService(usersRepo)
	adminService := adminService.NewService(usersRepo)

	// define routes
	common.CreateGroups(e, usersRepo)
	common.PassRouter(e)

	// initialize handlers
	usersHandler := users.NewHandler(usersService)
	classesHandler := classes.NewHandler(classesService)
	testEventsHandler := testevents.NewHandler(testEventsService, assignmentAppService)
	studentsHandler := students.NewHandler(studentsService, classesService)
	roomsHandler := rooms.NewHandler(roomsService)
	calendarHandler := calendar.NewHandler(*assignmentsDomainService)

	// mount handlers
	users.Mount(e, usersHandler)
	classes.Mount(e, classesHandler)
	testevents.Mount(e, testEventsHandler)
	students.Mount(e, studentsHandler)
	rooms.Mount(e, roomsHandler)
	about.Mount(e, about.NewHandler())
	calendar.Mount(e, calendarHandler)
	home.Mount(e, home.NewHandler())
	signout.Mount(e, signout.NewHandler())
	unauthorized.Mount(e, unauthorized.NewHandler())
	signup.Mount(e, signup.NewHandler(signupService))
	signin.Mount(e, signin.NewHandler(signinService))
	admin.Mount(e, admin.NewHandler(adminService))
	dashboard.Mount(e, dashboard.NewHandler(classesService, assignmentAppService))
	return nil
}
