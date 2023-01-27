package routes

import (
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/controllers"
	"itfest-backend-2.0/middlewares"
)

func UserRoute(e *echo.Echo) {
	e.GET("/user", controllers.GetUserHandler, middlewares.AuthMiddleware)
	e.POST("/findUser", controllers.FindUserHandler)
}
