package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mBuergi86/deaftube/entities"
	"github.com/mBuergi86/deaftube/repository"
	"github.com/mBuergi86/deaftube/utility"
	"sync"
)

var (
	lock sync.Mutex
)

// GetUsers the users will be completely from database
func GetUsers(r repository.UserRepository) fiber.Handler {
	lock.Lock()
	defer lock.Unlock()
	return func(c *fiber.Ctx) error {
		data, err := r.GetUsers()
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.Status(fiber.StatusOK).JSON(data)
	}
}

// GetUserByID an user is verified with an ID from database
func GetUserByID(r repository.UserRepository) fiber.Handler {
	lock.Lock()
	defer lock.Unlock()
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return utility.HandlerBadRequest(err)(c)
		}
		data, err := r.GetUserID(id)
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.Status(fiber.StatusOK).JSON(data)
	}
}

// CreateUser an new user will be recorded in the database
func CreateUser(r repository.UserRepository) fiber.Handler {
	lock.Lock()
	defer lock.Unlock()
	return func(c *fiber.Ctx) error {
		u := new(entities.SUsers)
		if err := c.BodyParser(u); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request playload",
				"error":   err.Error(),
			})
		}
		err := r.CreateUser(*u)
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User created successfully",
			"data":    u,
		})
	}
}

// UpdateUser an modified user will be changed in the database
func UpdateUser(r repository.UserRepository) fiber.Handler {
	lock.Lock()
	defer lock.Unlock()
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return utility.HandlerBadRequest(err)(c)
		}
		update := new(entities.SUsers)
		if err := c.BodyParser(update); err != nil {
			return utility.HandlerBadRequest(err)(c)
		}

		var users []entities.SUsers
		user := entities.SUsers{}

		if update.Firstname != " " {
			user = entities.SUsers{Firstname: update.Firstname}
		}
		if update.Lastname != " " {
			user = entities.SUsers{Lastname: update.Lastname}
		}
		if update.Username != "" {
			user.Username = update.Username
		}
		if update.Email != "" {
			user.Email = update.Email
		}
		if update.ChannelName != "" {
			user.ChannelName = update.ChannelName
		}
		if update.Password != "" {
			user.Password = update.Password
		}
		if update.PhotoUrl != "" {
			user.PhotoUrl = update.PhotoUrl
		}
		if update.Role != "" {
			user.Role = update.Role
		}

		users = append(users, user)

		err = r.UpdateUser(id, users)
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.Status(fiber.StatusOK).JSON("Successful")
	}
}

// DeleteUser an user will be deleted in the database
func DeleteUser(r repository.UserRepository) fiber.Handler {
	lock.Lock()
	defer lock.Unlock()
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return utility.HandlerBadRequest(err)(c)
		}
		err = r.DeleteUser(id)
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.Status(fiber.StatusOK).JSON("Successful")
	}
}
