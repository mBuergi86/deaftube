package router

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/mBuergi86/deaftube/handlers"
	"github.com/mBuergi86/deaftube/repository"
	"github.com/mBuergi86/deaftube/utility"
	"log"
)

func Router(app *fiber.App) {
	connStr := utility.ConnectString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := repository.NewUserRepository(db)

	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", handlers.GetUsers(repo))
	app.Get("/:id", handlers.GetUserByID(repo))
	app.Post("/", handlers.CreateUser(repo))
	app.Put("/:id", handlers.UpdateUser(repo))
	app.Delete("/:id", handlers.DeleteUser(repo))
}
