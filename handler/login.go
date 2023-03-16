package handler

import (
	"echo_framework/models"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	// echojwt "github.com/labstack/echo-jwt/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func Login(c echo.Context) error {

	username := c.Get("username").(string)
	log.Printf("Username: %v is logged in\n", username)
	is_admin := c.Get("admin").(bool)
	log.Printf("Is admin: %v\n", is_admin)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"admin":    is_admin,
		"exp":      time.Now().Add(180 * time.Second).Unix(),
	})

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		log.Printf("Singed token error: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, &models.LoginResponse{
		Token: t,
	})
}
