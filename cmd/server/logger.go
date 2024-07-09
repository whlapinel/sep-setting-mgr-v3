package main

import (
	"fmt"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

var logger = middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	LogURI:     true,
	LogStatus:  true,
	LogMethod:  true,
	LogLatency: true,
	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
		const (
			reset  = "\033[0m"
			red    = "\033[31m"
			green  = "\033[32m"
			yellow = "\033[33m"
		)
		statusColor := reset
		if v.Status >= 400 {
			statusColor = red
		} else if v.Status >= 300 {
			statusColor = yellow
		} else {
			statusColor = green
		}
		methodWidth := 6
		uriWidth := 25
		statusWidth := 6
		customWidth := 12
		latencyWidth := 15
		value, _ := c.Get("id").(int)
		logLine := fmt.Sprintf("%-*s %-*s %s%-*d%s %-*d %-*s",
			methodWidth, v.Method,
			uriWidth, v.URI,
			statusColor,
			statusWidth, v.Status,
			reset,
			customWidth, value,
			latencyWidth, v.Latency,
		)
		fmt.Println(logLine)
		return nil
	},
})
