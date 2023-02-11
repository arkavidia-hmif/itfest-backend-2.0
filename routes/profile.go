package routes

import (
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/controllers"
	"itfest-backend-2.0/middlewares"
)

func ProfileRoute(e *echo.Echo) {
	e.GET("/profile", controllers.GetProfileHandler, middlewares.AuthMiddleware)
	e.POST("/profile", controllers.UpdateProfileHandler, middlewares.AuthMiddleware)
}
