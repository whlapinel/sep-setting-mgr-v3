package main

import (
	"database/sql"
	"sep_setting_mgr/internal/domain/services"
	"sep_setting_mgr/internal/handlers/about"
	"sep_setting_mgr/internal/handlers/admin"
	adminapplications "sep_setting_mgr/internal/handlers/admin_applications"
	"sep_setting_mgr/internal/handlers/applications"
	"sep_setting_mgr/internal/handlers/calendar"
	"sep_setting_mgr/internal/handlers/classes"
	"sep_setting_mgr/internal/handlers/dashboard"
	common "sep_setting_mgr/internal/handlers/handlerscommon"
	"sep_setting_mgr/internal/handlers/home"
	"sep_setting_mgr/internal/handlers/rooms"
	"sep_setting_mgr/internal/handlers/signin"
	"sep_setting_mgr/internal/handlers/signout"
	"sep_setting_mgr/internal/handlers/signup"
	"sep_setting_mgr/internal/handlers/students"
	testevents "sep_setting_mgr/internal/handlers/test_events"
	"sep_setting_mgr/internal/handlers/unauthorized"
	"sep_setting_mgr/internal/handlers/users"
	"sep_setting_mgr/internal/repositories"
	adminService "sep_setting_mgr/internal/services/admin"
	applicationService "sep_setting_mgr/internal/services/applications"
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
	usersRepo := repositories.NewUsersRepo(db)
	classesRepo := repositories.NewClassesRepo(db)
	studentsRepo := repositories.NewStudentsRepo(db)
	testEventsRepo := repositories.NewTestEventsRepo(db)
	roomsRepo := repositories.NewRoomsRepo(db)
	assignmentsRepo := repositories.NewAssignmentsRepo(db)
	applicationsRepo := repositories.NewApplicationRepo(db)

	// initialize domain services
	assignmentsDomainService := services.NewAssignmentsService(assignmentsRepo, roomsRepo, testEventsRepo, classesRepo, studentsRepo)

	// initialize application services
	assignmentAppService := assignments.NewService(assignmentsRepo, roomsRepo, testEventsRepo)
	usersService := usersService.NewService(usersRepo)
	classesService := classesService.NewService(classesRepo, studentsRepo)
	testEventsService := testEventsService.NewService(testEventsRepo, classesRepo, *assignmentsDomainService)
	studentsService := studentsService.NewService(studentsRepo, classesRepo, *assignmentsDomainService)
	roomsService := roomsService.NewService(roomsRepo, *assignmentsDomainService)
	signupService := signupService.NewService(usersRepo)
	signinService := signinService.NewService(usersRepo)
	adminService := adminService.NewService(usersRepo, applicationsRepo)
	applicationService := applicationService.NewService(applicationsRepo, usersRepo)

	// define routes
	common.CreateGroups(e, usersRepo)

	// initialize handlers
	usersHandler := users.NewHandler(usersService)
	classesHandler := classes.NewHandler(classesService)
	testEventsHandler := testevents.NewHandler(testEventsService, assignmentAppService)
	studentsHandler := students.NewHandler(studentsService, classesService)
	roomsHandler := rooms.NewHandler(roomsService)
	calendarHandler := calendar.NewHandler(assignmentAppService, roomsService, testEventsService, studentsService)

	// mount handlers
	users.Mount(e, usersHandler)
	dashboard.Mount(e, dashboard.NewHandler(classesService, assignmentAppService, testEventsService, studentsService, roomsService))
	admin.Mount(e, admin.NewHandler(adminService))
	signup.Mount(e, signup.NewHandler(signupService))
	signin.Mount(e, signin.NewHandler(signinService))
	signout.Mount(e, signout.NewHandler())
	classes.Mount(e, classesHandler)
	testevents.Mount(e, testEventsHandler)
	students.Mount(e, studentsHandler)
	rooms.Mount(e, roomsHandler)
	calendar.Mount(e, calendarHandler)
	about.Mount(e, about.NewHandler())
	home.Mount(e, home.NewHandler())
	unauthorized.Mount(e, unauthorized.NewHandler())
	applications.Mount(e, applications.NewHandler(applicationService))
	adminapplications.Mount(e, adminapplications.NewHandler(applicationService))
	return nil
}
