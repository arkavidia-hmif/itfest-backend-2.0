package routes

import (
	"github.com/labstack/echo/v4"
	controllers "itfest-backend-2.0/controllers"
)

func LoginRoute(e *echo.Echo) {
	e.POST("/login", controllers.LoginHandler)
}
