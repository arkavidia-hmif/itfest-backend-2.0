package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/services"
	"itfest-backend-2.0/types"
	"net/http"
	"os"
	"strconv"
)

type GrantPointRequest struct {
	Usercode string `json:"usercode" form:"usercode" query:"usercode"`
	Point    uint   `json:"point" form:"point" query:"point"`
}

func GrantPointHandler(c echo.Context) error {
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
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	user, grantErr := services.GrantPoint(user.ID, request.Point)

	if grantErr != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Data = GrantPointRequest{Usercode: request.Usercode, Point: user.Point}

	return c.JSON(http.StatusOK, response)
}
