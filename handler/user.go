package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	name := c.QueryParam("name")
	user := &User{
		Name: name,
	}

	glog.Infof("id %d", user.Id)
	o := orm.NewOrm()

	err := o.Read(user, "Name")

	if err != nil {
		glog.Errorf("Error reading user: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := &User{}
	if err := c.Bind(user); err != nil {
		glog.Errorf("bind user error: %v", err)
		return err
	}
	glog.Info("request update user: %+v", user)
	o := orm.NewOrm()
	_, err := o.Update(user, "Name")

	if err != nil {
		glog.Errorf("Update user %s error: %v", user.Name, err)
		return err
	}
	userUpdate := &User{
		Name: user.Name,
	}
	o.Read(userUpdate, "Name")
	return c.JSON(http.StatusOK, userUpdate)

}

func DeleteUser(c echo.Context) error {
	s := c.FormValue("id")
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return err
	}
	glog.Info("Deleting user %d", id)
	user := &User{
		Id: id,
	}
	o := orm.NewOrm()
	row, err := o.Delete(user)
	if err != nil {
		glog.Errorf("Delete user %d failed: %v", row, err)
		return err
	}

	return c.String(http.StatusOK, fmt.Sprintf("Delete user id %d at %d", id, row))

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
