package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mBuergi86/deaftube/database"
	"log"
)

func Middleware(app *fiber.App) fiber.Handler {
	db := database.NewDBConnection()

	return func(ctx *fiber.Ctx) error {

		err := db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		//
		log.Printf("INFO: %s %s\n", ctx.Method(), ctx.Path())

		//
		return ctx.Next()
	}
}
