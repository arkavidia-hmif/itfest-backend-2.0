package repositories

import (
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
)

func GetUser(c echo.Context) models.Response[models.User] {
	db := configs.DB
	response := models.Response[models.User]{}

	user := models.User{}
	if err := db.Find(&user).Error; err != nil {
		response.Message = "ERROR get user"
		return response
	}

	response.Message = "SUCCESS get user"
	response.Data = user
	return response
}

func FindUser(c echo.Context) models.Response[models.User] {
	db := configs.DB
	response := models.Response[models.User]{}

	user := models.User{}
	// TODO
	condition := models.User{Usercode: "1234"}
	if err := db.Where(&condition, &user).Error; err != nil {
		response.Message = "ERROR find user"
		return response
	}

	response.Message = "SUCCESS get user"
	response.Data = user
	return response
}
