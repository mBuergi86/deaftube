package utility

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

func HandlerError(err error) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return nil
	}
}

func HandlerBadRequest(err error) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
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
