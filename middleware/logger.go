package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/mBuergi86/deaftube/database"
	"log"
)

func Middleware(app *fiber.App) fiber.Handler {
	app.Use(logger.New())
	app.Use(recover.New())
	db := database.NewDBConnection()

	return func(ctx *fiber.Ctx) error {

		err := db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("INFO: %s %s\n", ctx.Method(), ctx.Path())

		return ctx.Next()
	}
}
