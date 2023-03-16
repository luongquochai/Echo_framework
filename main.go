package main

import (
	"echo_framework/handler"
	"echo_framework/middlewares"

	"github.com/beego/beego/v2/client/orm"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// need to register models in init
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// need to register default database
	err := orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/echotest?charset=utf8")
	if err != nil {
		glog.Fatal("Failed to register database: %v", err)
	}

	//DB alias
	name := "default"

	// Drop table and re-create
	force := false

	// print log
	verbose := true

	// Create table
	err = orm.RunSyncdb(name, force, verbose)

	if err != nil {
		glog.Fatalf("Failed to sync database: %v", err)
	}
}

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
	groupUser.PUT("/add", handler.AddUser)
	groupUser.GET("/get", handler.GetUser)
	groupUser.GET("/get_all", handler.GetAllUsers)
	groupUser.GET("/update", handler.UpdateUser, isAdmin)
	groupUser.GET("/delete", handler.DeleteUser, isAdmin)

	server.Logger.Fatal(server.Start(":8888"))

}
