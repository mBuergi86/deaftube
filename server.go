package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mBuergi86/deaftube/middleware"
	"github.com/mBuergi86/deaftube/router"
	"os"
)

func main() {
	app := echo.New()

	router.Router(app)

	middleware.Middleware(app)

	httpPort := os.Getenv("PORT")

	if httpPort == "" {
		httpPort = "8080"
	}

	app.Logger.Fatal(app.Start(":" + httpPort))
}
