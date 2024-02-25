package routes

import (
	"scheduler-api/controller"

	"github.com/labstack/echo/v4"
)

func SetUserRoutes(e *echo.Echo) {
	e.POST("/user", controller.AddUser)
	e.GET("/user/page-index/:page-index/page-size/:page-size/:field/:order", controller.GetUsers)
	e.GET("/user/id/:id", controller.GetUserById)
	// e.GET("/gallery/category/:category", controller.GalleryByCategory, paramValidation)
	// e.GET("/gallery/tag/:tag", controller.GalleryByTag, paramValidation)
}
