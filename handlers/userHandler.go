package handlers

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mBuergi86/deaftube/entities"
	"github.com/mBuergi86/deaftube/repository"
	"github.com/mBuergi86/deaftube/utility"
)

var (
	lock sync.Mutex
)

// the users will be completely from database
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

// an user is verified with an ID from database
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

// an new user will be recorded in the database
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

// an modified user will be changed in the database
func UpdateUser(r repository.UserRepository) fiber.Handler {
	lock.Lock()
	defer lock.Unlock()
	return func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return utility.HandlerBadRequest(err)(c)
		}
		u := new(entities.SUsers)
		if err := c.BodyParser(u); err != nil {
			return err
		}
		err = r.UpdateUser(id, *u)
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.Status(fiber.StatusOK).JSON("Successful")
	}
}

// an user will be deleted in the database
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
