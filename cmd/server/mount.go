package main

import (
	"database/sql"
	"sep_setting_mgr/internal/database"
	"sep_setting_mgr/internal/pages/home"
	"sep_setting_mgr/internal/pages/teacher_dashboard"
	classesPkg "sep_setting_mgr/internal/pages/teacher_dashboard/htmx-classes"

	"github.com/labstack/echo"
)

func MountHandlers(e *echo.Echo, db *sql.DB) error {
	teacher_dashboard.Mount(
		e, teacher_dashboard.NewHandler(
			teacher_dashboard.NewService(
				database.NewClassesRepo(
					db))))
	classesPkg.Mount(
		e, classesPkg.NewHandler(
			classesPkg.NewService(
				database.NewClassesRepo(
					db))))
	home.Mount(e, home.NewHandler(home.NewService()))
	return nil
}
