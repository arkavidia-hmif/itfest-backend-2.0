package routes

import (
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/controllers"
	"itfest-backend-2.0/middlewares"
)

func MerchandiseRoute(e *echo.Echo) {
	e.POST("/addMerchandise", controllers.AddMerchandiseHandler, middlewares.AuthMiddleware)
}
