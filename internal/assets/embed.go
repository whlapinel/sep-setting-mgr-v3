package assets

import (
	"embed"
	"net/http"

	"github.com/labstack/echo"
)

//go:embed dist
var Assets embed.FS

func RegisterStatic(e *echo.Echo) {
	e.GET("/dist/*", echo.WrapHandler(http.FileServerFS(Assets)))
}
