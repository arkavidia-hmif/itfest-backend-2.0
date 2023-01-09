package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	con "itfest-backend-2.0/controllers"
)

func Run() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Declare new controller
	exampleCon := con.NewExampleController()

	// Routes
	e.GET("/example", exampleCon.GetHelloWorld)

	// Start server
	port := "8080"
	e.Logger.Fatal(e.Start(":" + port))
}
