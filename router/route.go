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

	app.Get("/user", handlers.GetUsers(repo))
	app.Get("/user/:id", handlers.GetUserByID(repo))
	app.Post("/user", handlers.CreateUser(repo))
	app.Patch("/user/:id", handlers.UpdateUser(repo))
	app.Delete("/user/:id", handlers.DeleteUser(repo))
}
