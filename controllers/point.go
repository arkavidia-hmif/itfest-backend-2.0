package controllers

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/services"
	"itfest-backend-2.0/types"
)

type GrantPointRequest struct {
	Usercode string `json:"usercode" form:"usercode" query:"usercode"`
	Point    uint   `json:"point" form:"point" query:"point"`
}

func GrantPointHandler(c echo.Context) error {
	id := c.Get("id").(uint)
	role := c.Get("role").(types.Role)
	response := models.Response[GrantPointRequest]{}

	if role != types.Startup {
		response.Message = "FORBIDDEN"
		return c.JSON(http.StatusUnauthorized, response)
	}

	request := GrantPointRequest{}
	if err := c.Bind(&request); err != nil {
		response.Message = "ERROR: BAD REQUEST"
		return c.JSON(http.StatusBadRequest, response)
	}

	maxGrant, err := strconv.Atoi(os.Getenv("MAX_GRANT_POINT"))

	if err != nil {
		return err
	}

	if request.Point > uint(maxGrant) {
		response.Message = fmt.Sprintf("ERROR: Added point should be less than %d", maxGrant)
		return c.JSON(http.StatusBadRequest, response)
	}

	db := configs.DB.GetConnection()

	user := models.User{}
	usercode := request.Usercode

	if err := db.Where(models.User{Usercode: usercode}).First(&user).Error; err != nil {
		response.Message = "ERROR: User Not Found"
		return c.JSON(http.StatusNotFound, response)
	}

	user, grantErr := services.GrantPoint(id, user.ID, request.Point)

	if grantErr != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Data = GrantPointRequest{Usercode: request.Usercode, Point: user.Point}

	return c.JSON(http.StatusOK, response)
}

func GetHistoriesHandler(c echo.Context) error {
	id := c.Get("id").(uint)
	response := models.Response[[]models.Log]{}

	db := configs.DB.GetConnection()
	var logs []models.Log

	err := db.Debug().Where(models.Log{FromId: id}).Or(models.Log{ToId: id}).
		Preload("From", func(d *gorm.DB) *gorm.DB {
			return d.Select("Name", "Usercode", "ID")
		}).
		Preload("To", func(d *gorm.DB) *gorm.DB {
			return d.Select("Name", "Usercode", "ID")
		}).
		Find(&logs).Error

	if err != nil {
		response.Message = "ERROR: FAILED TO QUERY POINT LOGS"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "SUCCESS"
	response.Data = logs
	return c.JSON(http.StatusOK, response)
}
