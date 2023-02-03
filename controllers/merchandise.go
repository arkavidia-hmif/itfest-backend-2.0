package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/types"
)

type MerchandiseOrder struct {
	MerchantID uint `json:"merch_id" form:"merch_id" query:"merch_id"`
	Quantity   uint `json:"quantity" form:"quantity" query:"quantity"`
}

type CheckoutRequest struct {
	To      uint               `json:"to" form:"to" query:"to"`
	Payload []MerchandiseOrder `json:"payload" form:"payload" query:"payload"`
}

func GetAllMerchandiseHandler(c echo.Context) error {
	response := models.Response[[]models.Merchandise]{}

	db := configs.DB.GetConnection()
	merchs := []models.Merchandise{}

	if err := db.Find(&merchs).Error; err != nil {
		response.Message = "ERROR: FAILED TO GET ALL MERCH"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "Success"
	response.Data = merchs

	return c.JSON(http.StatusOK, response)
}

func GetMerchandiseHandler(c echo.Context) error {
	response := models.Response[models.Merchandise]{}

	db := configs.DB.GetConnection()
	merch := models.Merchandise{}

	if err := db.First(&merch, c.Param("id")).Error; err != nil {
		response.Message = "ERROR: FAILED TO GET MERCH"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "Success"
	response.Data = merch

	return c.JSON(http.StatusOK, response)
}

func CheckoutHandler(c echo.Context) error {
	role := c.Get("role").(types.Role)
	response := models.Response[string]{}

	if role != types.Admin {
		response.Message = "FORBIDDEN"
		return c.JSON(http.StatusUnauthorized, response)
	}

	// Validating Request Body
	request := CheckoutRequest{}
	if err := c.Bind(&request); err != nil {
		response.Message = "ERROR: BAD REQUEST"
		return c.JSON(http.StatusBadRequest, response)
	}

	db := configs.DB.GetConnection()

	// Get User Point
	user := models.User{}
	if err := db.First(&user, request.To).Error; err != nil {
		response.Message = "ERROR: BAD REQUEST"
		return c.JSON(http.StatusBadRequest, response)
	}

	userPoint := user.Point

	// Get Total Merch Point
	totalMerchPoint := uint(0)

	for _, order := range request.Payload {
		merch := models.Merchandise{}
		if err := db.First(&merch, order.MerchantID).Error; err != nil {
			response.Message = "ERROR: BAD REQUEST"
			return c.JSON(http.StatusBadRequest, response)
		}

		totalMerchPoint += merch.Point * order.Quantity
	}

	// Check if user has enough point
	if userPoint < totalMerchPoint {
		response.Message = "ERROR: NOT ENOUGH POINT"
		return c.JSON(http.StatusBadRequest, response)
	}

	// Update User Point
	if err := db.Model(&user).Where("id = ?", user.ID).Update("point", user.Point-totalMerchPoint).Error; err != nil {
		response.Message = "ERROR: FAILED TO UPDATE POINT"
		return c.JSON(http.StatusInternalServerError, response)
	}

	// Update Merch Stock
	for _, order := range request.Payload {
		merch := models.Merchandise{}
		if err := db.First(&merch, order.MerchantID).Error; err != nil {
			response.Message = "ERROR: BAD REQUEST"
			return c.JSON(http.StatusBadRequest, response)
		}

		if err := db.Model(&merch).Where("id = ?", merch.ID).Update("stock", merch.Stock-order.Quantity).Error; err != nil {
			response.Message = "ERROR: FAILED TO UPDATE STOCK"
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	response.Message = "Success"
	return c.JSON(http.StatusOK, response)
}
