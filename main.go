package main

import (
	"golang_projects/projects/Echo_framework/handler"
	"golang_projects/projects/Echo_framework/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new server
	server := echo.New()

	server.Use(middleware.Logger())

	isLogedIn := middleware.JWT([]byte("secret"))
	isAdmin := middlewares.IsAdminMdw

	server.GET("/", handler.Hello, isLogedIn)

	server.POST("/login", handler.Login, middleware.BasicAuth(middlewares.BasicAuth))

	// middleware can sort and from left to right
	server.GET("/admin", handler.Hello, isLogedIn, isAdmin)

	// Group API
	groupUser := server.Group("/api/user")
	groupUser.GET("/get", handler.GetUser)
	groupUser.GET("/get_all", handler.GetAllUsers)
	groupUser.GET("/update", handler.UpdateUser, isAdmin)
	groupUser.GET("/delete", handler.DeleteUser, isAdmin)

	server.Logger.Fatal(server.Start(":8888"))

}
