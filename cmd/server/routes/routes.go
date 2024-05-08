package routes

//import echo
import "github.com/labstack/echo"

func RegisterRoutes(e *echo.Echo) {
	RegisterTestEventRoutes(e)
}