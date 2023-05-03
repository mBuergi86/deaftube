package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mBuergi86/deaftube/entities"
	"github.com/mBuergi86/deaftube/repository"
	"github.com/mBuergi86/deaftube/utility"
	"net/http"
	"sync"
)

var (
	lock sync.Mutex
)

func GetUsers(r repository.UserRepository) echo.HandlerFunc {
	lock.Lock()
	defer lock.Unlock()
	return func(c echo.Context) error {
		data, err := r.GetUsers()
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.JSON(http.StatusOK, data)
	}
}

func GetUserByID(r repository.UserRepository) echo.HandlerFunc {
	lock.Lock()
	defer lock.Unlock()
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return utility.HandlerBadRequest(err)(c)
		}
		data, err := r.GetUserID(id)
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.JSON(http.StatusOK, data)
	}
}

func CreateUser(r repository.UserRepository) echo.HandlerFunc {
	lock.Lock()
	defer lock.Unlock()
	return func(c echo.Context) error {
		name := c.FormValue("name")
		rename := c.FormValue("rename")
		username := c.FormValue("username")
		email := c.FormValue("email")
		password := c.FormValue("password")
		u := entities.SUsers{Firstname: name, Lastname: rename, Username: username, Email: email, Password: password}
		if err := c.Bind(u); err != nil {
			return err
		}
		err := r.CreateUser(u)
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.JSON(http.StatusOK, "Successful")
	}
}

func UpdateUser(r repository.UserRepository) echo.HandlerFunc {
	lock.Lock()
	defer lock.Unlock()
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return utility.HandlerBadRequest(err)(c)
		}
		name := c.FormValue("name")
		rename := c.FormValue("rename")
		username := c.FormValue("username")
		email := c.FormValue("email")
		password := c.FormValue("password")
		u := entities.SUsers{Firstname: name, Lastname: rename, Username: username, Email: email, Password: password}
		if err := c.Bind(u); err != nil {
			return err
		}
		err = r.UpdateUser(id, u)
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.JSON(http.StatusOK, "Successful")
	}
}

func DeleteUser(r repository.UserRepository) echo.HandlerFunc {
	lock.Lock()
	defer lock.Unlock()
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return utility.HandlerBadRequest(err)(c)
		}
		data, err := r.GetUserID(id)
		if err != nil {
			return utility.HandlerError(err)(c)
		}
		return c.JSON(http.StatusOK, data)
	}
}
