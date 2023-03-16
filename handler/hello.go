package handler

import (
	"echo_framework/models"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	username := claims["username"].(string)
	is_admin := claims["admin"].(bool)

	message := fmt.Sprintf("Hello %s has role admin %v", username, is_admin)

	x := &models.X{
		Text: message,
	}
	return c.JSON(http.StatusOK, x)
}
