package repository

import (
	"database/sql"
	"github.com/mBuergi86/deaftube/entities"
	"github.com/mBuergi86/deaftube/utility"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var (
	dbDriver = "postgres"
	connStr  = utility.ConnectString()
)

var u = &entities.SUsers{
	Firstname:   "Markus",
	Lastname:    "Bürgi",
	Username:    "MBuergi",
	Email:       "test@test.com",
	ChannelName: "cool",
	Password:    "12345",
	Role:        entities.Admin,
}

func NewConnect() *sql.DB {
	db, err := sql.Open(dbDriver, connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return db
}

func TestUserRepo_CreateUser(t *testing.T) {
	db := NewConnect()
	repo := NewUserRepository(db)
	defer func() {
		db.Close()
	}()
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	const createUser = `-- name: CreateUser :exec
insert into users (
                   firstname,
                   lastname,
                   username,
                   email,
                   channel_name,
                   password,
                   created_at,
                   update_at) values ( $1, $2, $3, $4, $5, $6, now(), now())
returning *;`

	_, err := repo.db.Exec(createUser, u)
	if err != nil {
		assert.Error(t, err)
	}
}
