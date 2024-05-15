package main

import (
	"database/sql"
	"sep_setting_mgr/internal/database"
	"sep_setting_mgr/internal/pages/about"
	"sep_setting_mgr/internal/pages/dashboard"
	"sep_setting_mgr/internal/pages/home"
	"sep_setting_mgr/internal/pages/signin"
	"sep_setting_mgr/internal/pages/signup"
	"sep_setting_mgr/internal/pages/unauthorized"

	"github.com/labstack/echo/v4"
)

func MountHandlers(e *echo.Echo, db *sql.DB) error {
	users := database.NewUsersRepo(db)
	home.Mount(e, home.NewHandler(home.NewService()))
	signin.Mount(e, signin.NewHandler(signin.NewService(users)))
	signup.Mount(e, signup.NewHandler(signup.NewService(users)))
	unauthorized.Mount(e, unauthorized.NewHandler())
	about.Mount(e, about.NewHandler())
	dashboard.Mount(e, dashboard.NewHandler(dashboard.NewService(database.NewClassesRepo(db), database.NewUsersRepo(db))))
	return nil
}
