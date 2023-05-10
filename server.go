package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mBuergi86/deaftube/middleware"
	"github.com/mBuergi86/deaftube/router"
	"log"
	"os"
)

func main() {
	app := fiber.New()
	middleware.Middleware(app)
	router.Router(app)

	httpPort := os.Getenv("PORT")

	if httpPort == "" {
		httpPort = "8080"
	}

	log.Fatal(app.Listen(":" + httpPort))
}
