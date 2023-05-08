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

func CreateUser(r repository.UserRepository) fiber.Handler {
	lock.Lock()
	defer lock.Unlock()
	return func(c *fiber.Ctx) error {
		u := new(entities.SUsers)
		if err := c.BodyParser(u); err != nil {
			return err
		}
		err := r.CreateUser(*u)
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.Status(fiber.StatusOK).JSON("Successful")
	}
}

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
