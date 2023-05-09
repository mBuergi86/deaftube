package repository

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/mBuergi86/deaftube/database"
	"github.com/mBuergi86/deaftube/entities"
	"log"
)

type UserRepository interface {
	GetUsers() ([]entities.SUsers, error)
	GetUserID(id uuid.UUID) (entities.SUsers, error)
	CreateUser(user entities.SUsers) error
	UpdateUser(id uuid.UUID, user []entities.SUsers) error
	DeleteUser(id uuid.UUID) error
}

type UserRepo struct {
	db *sql.DB
	entities.SUsers
}

func NewUserRepository(db *sql.DB) *UserRepo {
	return &UserRepo{db: database.NewDBConnection()}
}

const getUsers = `-- name: GetUsers :many
select firstname, lastname, username, channel_name, role, created_at, update_at from users order by firstname;`

func (u *UserRepo) GetUsers() ([]entities.SUsers, error) {
	rows, err := u.db.Query(getUsers)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []entities.SUsers

	for rows.Next() {
		var user entities.SUsers

		err := rows.Scan(
			&user.Firstname, &user.Lastname, &user.Username, &user.ChannelName, &user.Role, &user.CreatedAt, &user.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	return users, nil
}

const getUser = `-- name: GetUserID :one
select firstname, lastname, username, role from users where id = $1;`

func (u *UserRepo) GetUserID(id uuid.UUID) (entities.SUsers, error) {
	row := u.db.QueryRow(getUser, id)

	var user entities.SUsers
	err := row.Scan(
		&user.Firstname, &user.Lastname, &user.Username, &user.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.SUsers{}, errors.New("user not found")
		}
		return entities.SUsers{}, err
	}
	return user, nil
}

const createUser = `-- name: CreateUser :exec
insert into users (
                   firstname,
                   lastname,
                   username,
                   email,
                   channel_name,
                   password,
                   role,
                   created_at,
                   update_at) values ( $1, $2, $3, $4, $5, $6, $7, now(), now())
returning *;`

func (u *UserRepo) CreateUser(arg entities.SUsers) error {
	_, err := u.db.Exec(createUser, arg.Firstname, arg.Lastname, arg.Username, arg.Email, arg.ChannelName, arg.Password, arg.Role)
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
    photo_url = coalesce($8, photo_url),
    role = coalesce($9, role),
    update_at = now()
where id = $1;`

func (u *UserRepo) UpdateUser(id uuid.UUID, user []entities.SUsers) error {
	_, err := u.db.Exec(updateUser, id, user)
	if err != nil {
		return err
	}
	return nil
}

const deleteUser = `-- name: DeleteUser :exec
delete from users where id = $1;`

func (u *UserRepo) DeleteUser(id uuid.UUID) error {
	_, err := u.db.Exec(deleteUser, id)
	if err != nil {
		return err
	}
	return nil
}
