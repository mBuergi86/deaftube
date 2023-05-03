package middleware

import (
	"github.com/labstack/echo/v4"
	middleware2 "github.com/labstack/echo/v4/middleware"
)

func Middleware(app *echo.Echo) {
	app.Use(middleware2.Logger())
	app.Use(middleware2.Recover())
}
