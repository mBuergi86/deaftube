package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/mBuergi86/deaftube/entities"
)

type UserRepository interface {
	GetUsers() ([]entities.SUsers, error)
	GetUserID(id uuid.UUID) (entities.SUsers, error)
	CreateUser(user entities.SUsers) error
	UpdateUser(id uuid.UUID, user entities.SUsers) error
	DeleteUser(id uuid.UUID) error
}

type UserRepo struct {
	db *sql.DB
	entities.SUsers
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

const getUsers = `-- name: GetUsers :many
select * from users order by firstname;`

func (u *UserRepo) GetUsers() ([]entities.SUsers, error) {
	ctx := context.TODO()
	row := u.db.QueryRowContext(ctx, getUsers)
	var users []entities.SUsers
	err := row.Scan(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

const getUser = `-- name: GetUserID :one
select * from users where id = $1 limit 1;`

func (u *UserRepo) GetUserID(id uuid.UUID) (entities.SUsers, error) {
	ctx := context.TODO()
	row := u.db.QueryRowContext(ctx, getUser, id)
	var users entities.SUsers
	err := row.Scan(&users)
	if err != nil {
		return entities.SUsers{}, err
	}
	return users, nil
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

func (u *UserRepo) CreateUser(arg entities.SUsers) error {
	ctx := context.TODO()
	_, err := u.db.ExecContext(ctx, createUser, arg.Firstname, arg.Lastname, arg.Username, arg.Email, arg.ChannelName, arg.Password)
	if err != nil {
		return err
	}
	return nil
}

const updateUser = `-- name: UpdateUser :exec
update users set
    firstname = coalesce($2, firstname),
    lastname = coalesce($3, lastname),
    username = coalesce($4, username),
    email = coalesce($5, email),
    channel_name = coalesce($6, channel_name),
    password = coalesce($7, password),
    update_at = now()
where id = $1;`

func (u *UserRepo) UpdateUser(id uuid.UUID, arg entities.SUsers) error {
	ctx := context.TODO()
	_, err := u.db.ExecContext(ctx, updateUser, id, arg.Firstname, arg.Lastname, arg.Username, arg.Email, arg.ChannelName, arg.Password)
	if err != nil {
		return err
	}
	return nil
}

const deleteUser = `-- name: DeleteUser :exec
delete from users where id = $1;`

func (u *UserRepo) DeleteUser(id uuid.UUID) error {
	ctx := context.TODO()
	_, err := u.db.ExecContext(ctx, deleteUser, id)
	if err != nil {
		return err
	}
	return nil
}
