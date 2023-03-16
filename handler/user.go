package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/golang/glog"
	"github.com/labstack/echo/v4"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id    int64  `orm:"auto" json:"id" form:"id" query:"id"`
	Name  string `json:"name" form:"name" query:"name"`
	Age   int    `json:"age" form:"age" query:"age"`
	Phone string `json:"phone" form:"phone" query:"phone`
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

func AddUser(c echo.Context) error {
	user := &User{}
	err := c.Bind(user)
	if err != nil {
		glog.Errorf("Bind user error: %v", err)
		return err
	}

	glog.Infof("%v", user)
	o := orm.NewOrm()

	id, err := o.Insert(user)

	if err != nil {
		glog.Errorf("Insert user error: %v", err)
		return err
	}

	glog.Infof("Insert user at row %d", id)
	return c.JSON(http.StatusOK, user)
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
