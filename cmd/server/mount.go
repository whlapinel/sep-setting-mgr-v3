package main

import (
	"database/sql"
	"sep_setting_mgr/internal/database"
	"sep_setting_mgr/internal/pages/about"
	"sep_setting_mgr/internal/pages/admin"
	"sep_setting_mgr/internal/pages/dashboard"
	"sep_setting_mgr/internal/pages/home"
	"sep_setting_mgr/internal/pages/signin"
	"sep_setting_mgr/internal/pages/signout"
	"sep_setting_mgr/internal/pages/signup"
	"sep_setting_mgr/internal/pages/unauthorized"

	"github.com/labstack/echo/v4"
)

func MountHandlers(e *echo.Echo, db *sql.DB) error {
	users := database.NewUsersRepo(db)
	classes := database.NewClassesRepo(db)
	students := database.NewStudentsRepo(db)
	testEvents := database.NewTestEventsRepo(db)
	home.Mount(e, home.NewHandler())
	signout.Mount(e, signout.NewHandler())
	unauthorized.Mount(e, unauthorized.NewHandler())
	about.Mount(e, about.NewHandler())
	signup.Mount(e, signup.NewHandler(signup.NewService(users)))
	signin.Mount(e, signin.NewHandler(signin.NewService(users)))
	dashboard.Mount(e, dashboard.NewHandler(dashboard.NewService(classes, users, students, testEvents)))
	admin.Mount(e, admin.NewHandler(admin.NewService(users)))
	return nil
}
