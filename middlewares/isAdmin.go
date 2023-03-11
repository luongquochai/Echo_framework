package middlewares

import (
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func IsAdminMdw(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		is_admin := claims["admin"].(bool)
		log.Printf("DEBUG: isAdmin=%v", is_admin)
		if is_admin {
			next(c)
		}

		return echo.ErrUnauthorized
	}
}
