package router

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/mBuergi86/deaftube/database"
	"github.com/mBuergi86/deaftube/handlers"
	"github.com/mBuergi86/deaftube/repository"
)

func Router(app *fiber.App) {
	db := database.NewDBConnection()
	repo := repository.NewUserRepository(db)

	app.Get("/", handlers.GetUsers(repo))
	app.Get("/:id", handlers.GetUserByID(repo))
	app.Post("/", handlers.CreateUser(repo))
	app.Put("/:id", handlers.UpdateUser(repo))
	app.Delete("/:id", handlers.DeleteUser(repo))
}
