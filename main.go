package main

import (
	"fmt"
	"os"
	"snake-scape/internal/middleware"
	"snake-scape/internal/view"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Load()

	e := echo.New()

	// Apply middlewares
	e.Use(middleware.CustomContextMiddleware)
	e.Use(middleware.LogMiddleware)
	e.Use(middleware.ParseFormMiddleware)
	// e.Use(middleware.ParseMultipartFormMiddleware)  // Uncomment if needed

	// Static files
	e.GET("/favicon.ico", view.ServeFavicon)
	e.GET("/static/*", view.ServeStaticFiles)

	// Routes
	e.GET("/", view.Home)

	port := os.Getenv("PORT")
	fmt.Printf("Server is running on port %s\n", port)
	e.Logger.Fatal(e.Start(":" + port))
}

