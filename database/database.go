package database

import (
	"database/sql"
	"log"

	"github.com/mBuergi86/deaftube/utility"
)

var (
	dbDriver = "postgres"
	connStr  = utility.ConnectString()
)

func NewDBConnection() *sql.DB {
	// get a database connection
	db, err := sql.Open(dbDriver, connStr)
	if err != nil {
		log.Fatal(err)
	}

	// check a database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// a database connection is ready
	return db
}
