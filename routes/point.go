package routes

import (
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/controllers"
	"itfest-backend-2.0/middlewares"
)

func PointsRoute(e *echo.Echo) {
	e.GET("/getHistories", controllers.GetHistoriesHandler, middlewares.AuthMiddleware)
	e.POST("/grantPoint", controllers.GrantPointHandler, middlewares.AuthMiddleware)
}
