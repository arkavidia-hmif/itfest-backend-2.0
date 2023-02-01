package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/middlewares"
	"itfest-backend-2.0/routes"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.Cors())

	e.Validator = &configs.RequestValidator{Validator: validator.New()}

	routes.LoginRoute(e)
	routes.RegisterRoute(e)
	routes.UserRoute(e)
	routes.ProfileRoute(e)
	routes.PointsRoute(e)

	port := "8080"
	e.Logger.Fatal(e.Start(":" + port))
}
