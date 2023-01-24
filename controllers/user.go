package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	repo "itfest-backend-2.0/repositories"
)

type userController struct{}

// Add new function here then make new one below.
type UserController interface {
	GetUser(c echo.Context) error
	FindUser(c echo.Context) error
}

func NewUserController() UserController {
	return &userController{}
}

func (*userController) GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, repo.GetUser(c))
}

func (*userController) FindUser(c echo.Context) error {
	return c.JSON(http.StatusOK, repo.FindUser(c))
}
