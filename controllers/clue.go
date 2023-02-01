package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func getRandomClueId() uint {
	return 1
}

func ClueHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "hi")
}

func SubmitClueHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "hi")
}
