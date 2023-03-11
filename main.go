package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang_projects/projects/Echo_framework/handler"
	"golang_projects/projects/Echo_framework/middlewares"
)

func main() {
	// Create a new server
	server := echo.New()

	server.Use(middleware.Logger())

	isLoggIn := middleware.JWT([]byte("secret"))
	server.GET("/", handler.Hello, isLoggIn)

	server.POST("/login", handler.Login, middleware.BasicAuth(middlewares.BasicAuth))

	server.Logger.Fatal(server.Start(":8888"))

}


