package routes

import (
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/controllers"
	"itfest-backend-2.0/middlewares"
)

func MerchandiseRoute(e *echo.Echo) {
	e.GET("/getAllMerchandise", controllers.GetAllMerchandiseHandler, middlewares.AuthMiddleware)
	e.GET("/getMerchandise/:id", controllers.GetMerchandiseHandler, middlewares.AuthMiddleware)
	e.POST("/checkout", controllers.CheckoutHandler, middlewares.AuthMiddleware)
	e.POST("/addMerchandise", controllers.AddMerchandiseHandler, middlewares.AuthMiddleware)
	e.POST("/deleteMerchandise", controllers.DeleteMerchandiseHandler, middlewares.AuthMiddleware)
}
