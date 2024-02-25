package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	e "scheduler-api/entity"
	m "scheduler-api/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddUser(c echo.Context) error {
	fmt.Printf("echo value: %v\n", c)
	var user e.User
	err := c.Bind(&user)

	///if err != nil {
	//		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	//	}

	err = m.AddUser(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find user listing")
	}

	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", EmptyValue))

	//return c.JSON(http.StatusOK, gallery)
}

func GetUsers(c echo.Context) error {
	pageIndex, err := strconv.ParseUint(c.Param("page-index"), 10, 64)
	pageSize, err := strconv.ParseUint(c.Param("page-size"), 10, 64)
	field := c.Param("field")
	order := c.Param("order")

	users, err := m.GetUsers(pageIndex, pageSize, field, order)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find user listing")
	}

	usersBytes, err := json.Marshal(users)
	uersJson := ConvertStructToJSON(usersBytes)

	return c.JSON(http.StatusOK, uersJson)
}

func GetUserById(c echo.Context) error {

	userId := c.Param("id")
	user, err := m.GetUserById(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find user")
	}
	//json.Marshal()
	//user.UsherGroup = json.Unmarshal(user.UsherGroup)
	return c.JSON(http.StatusOK, user)
}
