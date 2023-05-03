package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mBuergi86/deaftube/entities"
	"time"
)

type UserRepository interface {
	GetUsers() ([]entities.SUsers, error)
	GetUserID(id uuid.UUID) (entities.SUsers, error)
	CreateUser(user entities.SUsers) error
	UpdateUser(id uuid.UUID, user entities.SUsers) error
	DeleteUser(id uuid.UUID) error
}

type UserRepo struct {
	db *sqlx.DB
	entities.SUsers
}

func NewUserRepository(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) GetUsers() ([]entities.SUsers, error) {
	var users []entities.SUsers
	err := u.db.Select(&users, "SELECT * FROM users ORDER BY firstname ASC")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepo) GetUserID(id uuid.UUID) (entities.SUsers, error) {
	var users entities.SUsers
	err := u.db.Select(&users, "SELECT * FROM users WHERE uuid=$1", id)
	if err != nil {
		return entities.SUsers{}, err
	}
	return users, nil
}

func (u *UserRepo) CreateUser(user entities.SUsers) error {
	newUUID, _ := uuid.NewUUID()
	createdAt, updatedAt := time.Now(), time.Now()
	_, err := u.db.Exec(`
        INSERT INTO users (uuid, name, rename, username, role, email, password, photo_url, kannel_name, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    `, newUUID, user.Firstname, user.Lastname, user.Username, "user", user.Email, user.Password, user.PhotoUrl, user.KannelName, createdAt, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) UpdateUser(id uuid.UUID, user entities.SUsers) error {
	newUUID, _ := uuid.NewUUID()
	updatedAt := time.Now()

	_, err := u.db.Exec(`
        UPDATE users 
        SET name = $1, rename = $2 
        WHERE uuid = $11;
		`, newUUID, user.Firstname, user.Lastname, user.Username, user.Role, user.Password, user.PhotoUrl, user.KannelName, updatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) DeleteUser(id uuid.UUID) error {
	//usersMap := map[uuid.UUID]*entities.SUsers{}
	panic("implement me")
}
