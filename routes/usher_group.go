package routes

import (
	"scheduler-api/controller"

	"github.com/labstack/echo/v4"
)

func SetUsherGroupRoutes(e *echo.Echo) {
	e.POST("/usher-group", controller.AddUserUsherGroup)
	e.GET("/usher-group/page-index/:page-index/page-size/:page-size/:field/:order", controller.GetUsherGroups)
	e.GET("/usher-group/id/:id", controller.GetUsherGroupsById)
	e.GET("/usher-group/key-value", controller.GetUsherGroupsById)
	// e.GET("/gallery/category/:category", controller.GalleryByCategory, paramValidation)
	// e.GET("/gallery/tag/:tag", controller.GalleryByTag, paramValidation)
}
