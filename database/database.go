package database

import (
	"database/sql"
	"github.com/mBuergi86/deaftube/utility"
	"log"
)

var (
	dbDriver = "postgres"
	connStr  = utility.ConnectString()
)

func NewDBConnection() *sql.DB {
	//
	db, err := sql.Open(dbDriver, connStr)
	if err != nil {
		log.Fatal(err)
	}

	//
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
