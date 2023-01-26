package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
)

// TODO: @graceclaudia19
func GetUserHandler(c echo.Context) error {
	db := configs.DB.GetConnection()
	response := models.Response[string]{}

	request := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		response.Message = "ERROR: BAD REQUEST"
		return c.JSON(http.StatusBadRequest, response)
	}

	result := models.User{}
	usercode := request["ID"].(string)
	condition := models.User{Usercode: usercode}
	if err := db.Find(&condition, &result).Error; err != nil {
		response.Message = "ERROR: USER NOT FOUND"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Data = result.Name
	return c.JSON(http.StatusOK, response)
}

// TODO: @graceclaudia19
func FindUserHandler(c echo.Context) error {
	return nil
	// return c.JSON(http.StatusOK, repo.FindUser(models.User{}))
}
