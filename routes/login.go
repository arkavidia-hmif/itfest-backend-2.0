package routes

import (
	"github.com/labstack/echo/v4"
	controllers "itfest-backend-2.0/controllers"
	"itfest-backend-2.0/middlewares"
)

func LoginRoute(e *echo.Echo) {
	e.POST("/login", controllers.Testing, middlewares.AuthMiddleware)
}
