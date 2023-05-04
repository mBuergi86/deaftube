package router

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/mBuergi86/deaftube/handlers"
	"github.com/mBuergi86/deaftube/repository"
	"github.com/mBuergi86/deaftube/utility"
	"log"
)

func Router(app *echo.Echo) {
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

	app.GET("/", handlers.GetUsers(repo))
	app.GET("/:id", handlers.GetUserByID(repo))
	app.POST("/", handlers.CreateUser(repo))
	app.PUT("/:id", handlers.UpdateUser(repo))
	app.DELETE("/:id", handlers.DeleteUser(repo))
}
