package view

import (
	"net/http"
	"path/filepath"
	"snake-scape/internal/middleware"
	"snake-scape/internal/template"

	"github.com/labstack/echo/v4"
)

func ServeFavicon(c echo.Context) error {
	filePath := "favicon.ico"
	fullPath := filepath.Join(".", "static", filePath)
	return c.File(fullPath)
}

func ServeStaticFiles(c echo.Context) error {
	filePath := c.Param("*")
	fullPath := filepath.Join(".", "static", filePath)
	return c.File(fullPath)
}

func Home(c echo.Context) error {
	cc, ok := c.(*middleware.CustomContext)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not get custom context")
	}

	if c.Request().URL.Path != "/" {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return template.Home("Templ Quickstart").Render(cc, c.Response().Writer)
}

