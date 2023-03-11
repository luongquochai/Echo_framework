package middlewares

import (
	"github.com/labstack/echo/v4"
)

func BasicAuth(username string, password string, c echo.Context) (bool, error) {
	if username == "admin" && password == "123" {
		//Transfer the username -> login
		c.Set("username", username)
		c.Set("admin", true)
		return true, nil
	} else {
		c.Set("username", username)
		c.Set("admin", false)
		return true, nil
	}
}
