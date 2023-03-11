package handler

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt"
	"golang_projects/projects/Echo_framework/models"
)

func Hello(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	username :=	claims["username"].(string)
	is_admin := claims["admin"].(bool)

	message := fmt.Sprintf("Hello %s has role admin %v", username, is_admin)

	x := &models.X{
		Text: message,
	}
	return c.JSON(http.StatusOK, x)
}