package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/types"
)

type FindUserRequest struct {
	Usercode string `json:"usercode" form:"usercode" query:"usercode"`
}

type UserResponse struct {
	Name     string     `json:"name"`
	Username string     `json:"username"`
	Usercode string     `json:"usercode"`
	Role     types.Role `json:"role"`
	Point    uint       `json:"point"`
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

func FindUserHandler(c echo.Context) error {
	db := configs.DB.GetConnection()
	response := models.Response[UserResponse]{}

	role := c.Get("role")
	if role == types.User {
		response.Message = "ERROR: UNAUTHORIZED"
		return c.JSON(http.StatusUnauthorized, response)
	}

	request := FindUserRequest{}
	if err := c.Bind(&request); err != nil {
		response.Message = "ERROR: BAD REQUEST"
		return c.JSON(http.StatusBadRequest, response)
	}

	result := models.User{}
	condition := models.User{Usercode: request.Usercode}
	if err := db.Where(&condition).Find(&result).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
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
