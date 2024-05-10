package main

import (
	//import echo

	"sep_setting_mgr/internal/assets"
	"sep_setting_mgr/internal/domain"
	classesPkg "sep_setting_mgr/internal/features/classes"
	"sep_setting_mgr/internal/features/teacher_dashboard"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()
	classes := domain.NewClasses()
	classes.Add("Math", 1)
	teacher_dashboard.Mount(e, teacher_dashboard.NewHandler(teacher_dashboard.NewService(classes)))
	classesPkg.Mount(e, classesPkg.NewHandler(classesPkg.NewService(classes)))
	// testEventRepo := repository.NewMemoryTestEventRepository()
	// testEventService := services.NewTestEventService(testEventRepo)
	// routes.RegisterTestEventRoutes(e, testEventService)
	// routes.RegisterRoutes(e)
	// insert user service and repo wiring up here
	assets.RegisterStatic(e)
	e.Logger.Fatal(e.Start(":1323"))
}
