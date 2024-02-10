package controller

import (
	"net/http"
	m "scheduler-api/model"

	"github.com/labstack/echo/v4"
)

func GetUsherGroups(c echo.Context) error {
	gallery, err := m.GetMainGallery()

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find gallery listing")
	}

	return c.JSON(http.StatusOK, gallery)
}

func GetUsherGroup(c echo.Context) error {
	category := c.Param("category")

	gallery, err := m.GetGalleryByCategory(category)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find category")
	}

	return c.JSON(http.StatusOK, gallery)
}
