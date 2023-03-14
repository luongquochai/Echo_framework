package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string
	Age  int
}

var listUsers = []User{
	{
		Name: "John",
		Age:  18,
	},
	{
		Name: "Jane",
		Age:  20,
	},
	{
		Name: "Hai",
		Age:  24,
	},
	{
		Name: "Jack",
		Age:  30,
	},
}

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "api get user")
}

func UpdateUser(c echo.Context) error {
	return c.String(http.StatusOK, "api update user")
}

func DeleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "api delete user")
}

func GetAllUsers(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(http.StatusOK)

	enc := json.NewEncoder(c.Response())

	for _, user := range listUsers {
		if err := enc.Encode(user); err != nil {
			return err
		}

		c.Response().Flush()
		time.Sleep(1 * time.Second)
	}

	return nil
}
