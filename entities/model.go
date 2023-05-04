package entities

import (
	"github.com/google/uuid"
	"time"
)

type Role string

const (
	User  Role = "user"
	Admin Role = "admin"
)

type SUsers struct {
	ID          uuid.UUID `db:"id"`
	Firstname   string    `db:"firstname"`
	Lastname    string    `db:"lastname"`
	Username    string    `db:"username"`
	Email       string    `db:"email"`
	ChannelName string    `db:"channel_name"`
	Password    string    `db:"password"`
	PhotoUrl    string    `db:"photo_url"`
	Role        Role      `db:"role"`
	CreatedAt   time.Time `db:"created_at"`
	UpdateAt    time.Time `db:"update_at"`
}
