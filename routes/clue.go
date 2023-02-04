package routes

import (
	"github.com/labstack/echo/v4"
	controllers "itfest-backend-2.0/controllers"
	"itfest-backend-2.0/middlewares"
)

func ClueRoute(e *echo.Echo) {
	e.GET("/getClue", controllers.ClueHandler, middlewares.AuthMiddleware)
	e.POST("/submitClue", controllers.SubmitClueHandler, middlewares.AuthMiddleware)
	e.POST("/createClue", controllers.CreateClueHandler, middlewares.AuthMiddleware)
}
