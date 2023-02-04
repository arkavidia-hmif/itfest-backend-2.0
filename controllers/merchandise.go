package controllers

import (
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/types"
	"net/http"
)

type AddMerchandiseRequest struct {
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Stock uint   `json:"stock" form:"stock" query:"stock" validate:"required,min=0"`
	Point uint   `json:"point" form:"point" query:"point" validate:"required,min=1"`
}

func AddMerchandiseHandler(c echo.Context) error {
	role := c.Get("role").(types.Role)
	response := models.Response[models.Merchandise]{}
	newMerchandise := new(AddMerchandiseRequest)

	if role != types.Admin {
		response.Message = "FORBIDDEN"
		return c.JSON(http.StatusUnauthorized, response)
	}

	if err := c.Bind(newMerchandise); err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(newMerchandise); err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	db := configs.DB.GetConnection()
	merchandise := models.Merchandise{
		Name:  newMerchandise.Name,
		Stock: newMerchandise.Stock,
		Point: newMerchandise.Point,
	}

	if err := db.Create(&merchandise).Error; err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Data = merchandise

	return c.JSON(http.StatusOK, response)
}
