package utility

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func HandlerError(err error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return nil
	}
}

func HandlerBadRequest(err error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		return nil
	}
}

func ConnectString() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	connStr := os.Getenv("CONN_STR")
	return connStr
}
