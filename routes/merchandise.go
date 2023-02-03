package routes

import (
	"github.com/labstack/echo/v4"
	controllers "itfest-backend-2.0/controllers"
	"itfest-backend-2.0/middlewares"
)

func MerchandiseRoute(e *echo.Echo) {
	e.GET("/getAllMerchandise", controllers.GetAllMerchandiseHandler, middlewares.AuthMiddleware)
	e.GET("/getMerchandise/:id", controllers.GetMerchandiseHandler, middlewares.AuthMiddleware) // TODO: Make Sure Query Param Comply with API Contract
	e.POST("/checkout", controllers.LoginHandler, middlewares.AuthMiddleware)
}
