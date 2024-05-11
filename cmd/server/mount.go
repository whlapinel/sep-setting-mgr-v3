package main

import (
	"sep_setting_mgr/internal/domain"
	"sep_setting_mgr/internal/pages/teacher_dashboard"
	classesPkg "sep_setting_mgr/internal/pages/teacher_dashboard/htmx-classes"

	"github.com/labstack/echo"
)

func MountHandlers(e *echo.Echo) error {
	classes := domain.NewClasses()
	classes.Add("Math", 1)
	teacher_dashboard.Mount(e, teacher_dashboard.NewHandler(teacher_dashboard.NewService(classes)))
	classesPkg.Mount(e, classesPkg.NewHandler(classesPkg.NewService(classes)))
}
