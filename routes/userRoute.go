package routes

import (
	controllers "itfest-backend-2.0/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo) {
	userCon := controllers.NewUserController();

	e.GET("/user", userCon.GetUser)
	e.POST("/findUser", userCon.FindUser)
}
