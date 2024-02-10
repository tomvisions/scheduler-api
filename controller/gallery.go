package controller

import (
	"net/http"
	m "scheduler-api/model"

	"github.com/labstack/echo/v4"
)

func MainGallery(c echo.Context) error {
	gallery, err := m.GetMainGallery()

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find gallery listing")
	}

	return c.JSON(http.StatusOK, gallery)
}

func GalleryByCategory(c echo.Context) error {
	category := c.Param("category")

	gallery, err := m.GetGalleryByCategory(category)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find category")
	}

	return c.JSON(http.StatusOK, gallery)
}

func GalleryByTag(c echo.Context) error {
	tag := c.Param("tag")
	gallery, err := m.GetGalleryByTag(tag)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find tag")
	}

	return c.JSON(http.StatusOK, gallery)
}
