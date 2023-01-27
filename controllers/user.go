package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/types"
)

type UserResponse struct {
	Name     string
	Username string
	Usercode string
	Role     types.Role
	Point    uint
}

func GetUserHandler(c echo.Context) error {
	db := configs.DB.GetConnection()
	response := models.Response[UserResponse]{}

	id := c.Get("id")
	result := models.User{}
	if err := db.First(&result, id).Error; err != nil {
		response.Message = "ERROR: USER NOT FOUND"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "SUCCESS"
	response.Data = UserResponse{
		Name:     result.Name,
		Username: result.Username,
		Usercode: result.Usercode,
		Role:     result.Role,
		Point:    result.Point,
	}
	return c.JSON(http.StatusOK, response)
}

// TODO: @graceclaudia19
func FindUserHandler(c echo.Context) error {
	return nil
	// return c.JSON(http.StatusOK, repo.FindUser(models.User{}))
}
