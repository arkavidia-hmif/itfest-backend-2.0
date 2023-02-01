package routes

import (
	"github.com/labstack/echo/v4"
	controllers "itfest-backend-2.0/controllers"
)

func ClueRoute(e *echo.Echo) {
	e.GET("/getClue", controllers.ClueHandler)
	e.POST("/submitClue", controllers.SubmitClueHandler)
}
