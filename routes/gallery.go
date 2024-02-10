package routes

import (
	"scheduler-api/controller"

	"github.com/labstack/echo/v4"
)

func SetGalleryRoutes(e *echo.Echo) {
	e.GET("/gallery", controller.MainGallery)
	e.GET("/gallery/category/:category", controller.GalleryByCategory, paramValidation)
	e.GET("/gallery/tag/:tag", controller.GalleryByTag, paramValidation)
}
