package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	repo "itfest-backend-2.0/repositories"
)

type exampleController struct{}

// Add new function here then make new one below.
type ExampleController interface {
	GetHelloWorld(c echo.Context) error
}

func NewExampleController() ExampleController {
	return &exampleController{}
}

func (*exampleController) GetHelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, repo.GetHelloWorld())
}
