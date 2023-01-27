package routes

import (
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/controllers"
)

func UserRoute(e *echo.Echo) {
	e.GET("/user", controllers.GetUserHandler)
	e.POST("/findUser", controllers.FindUserHandler)
}
