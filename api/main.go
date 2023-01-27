package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"itfest-backend-2.0/middlewares"
	"itfest-backend-2.0/routes"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.Cors())

	routes.LoginRoute(e)
	routes.RegisterRoute(e)

	port := "8080"
	e.Logger.Fatal(e.Start(":" + port))
}
