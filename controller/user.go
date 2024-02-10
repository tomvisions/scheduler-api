package controller

import (
	"net/http"
	e "scheduler-api/entity"
	m "scheduler-api/model"

	"github.com/labstack/echo/v4"
)

func AddUser(c echo.Context) error {
	var user e.User
	err := c.Bind(&user)

	///if err != nil {
	//		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	//	}

	err = m.AddUser(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find gallery listing")
	}

	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", EmptyValue))

	//return c.JSON(http.StatusOK, gallery)
}

func GetUsers(c echo.Context) error {
	gallery, err := m.GetMainGallery()

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find gallery listing")
	}

	return c.JSON(http.StatusOK, gallery)
}

func GetUser(c echo.Context) error {
	category := c.Param("category")

	gallery, err := m.GetGalleryByCategory(category)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find category")
	}

	return c.JSON(http.StatusOK, gallery)
}
