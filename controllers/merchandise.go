package controllers

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/types"
)

type AddMerchandiseRequest struct {
	Name     string `json:"name" form:"name" query:"name" validate:"required"`
	Stock    uint   `json:"stock" form:"stock" query:"stock" validate:"required,min=0"`
	Point    uint   `json:"point" form:"point" query:"point" validate:"required,min=1"`
	Usercode string `json:"usercode" form:"usercode" query:"usercode" validate:"required"`
}

type DeleteMerchandiseRequest struct {
	ID uint `json:"id" form:"id" query:"id" validate:"required"`
}

type MerchandiseOrder struct {
	MerchantID uint `json:"merch_id" form:"merch_id" query:"merch_id"`
	Quantity   uint `json:"quantity" form:"quantity" query:"quantity"`
}

type CheckoutRequest struct {
	To      string             `json:"to" form:"to"`
	Payload []MerchandiseOrder `json:"payload" form:"payload" binding:"dive"`
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
		Name:     newMerchandise.Name,
		Stock:    newMerchandise.Stock,
		Point:    newMerchandise.Point,
		Usercode: newMerchandise.Usercode,
	}

	if err := db.Create(&merchandise).Error; err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Data = merchandise

	return c.JSON(http.StatusOK, response)
}

func GetAllMerchandiseHandler(c echo.Context) error {
	response := models.Response[[]models.Merchandise]{}

	db := configs.DB.GetConnection()
	merchs := []models.Merchandise{}

	if err := db.Preload("User", func(d *gorm.DB) *gorm.DB {
		return d.Select("Name", "Usercode", "ID")
	}).Find(&merchs).Error; err != nil {
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

	if err := db.Preload("User", func(d *gorm.DB) *gorm.DB {
		return d.Select("Name", "Usercode", "ID")
	}).First(&merch, c.Param("id")).Error; err != nil {
		response.Message = "ERROR: FAILED TO GET MERCH"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "Success"
	response.Data = merch

	return c.JSON(http.StatusOK, response)
}

func DeleteMerchandiseHandler(c echo.Context) error {
	db := configs.DB.GetConnection()

	role := c.Get("role").(types.Role)
	response := models.Response[string]{}

	if role != types.Admin {
		response.Message = "FORBIDDEN"
		return c.JSON(http.StatusUnauthorized, response)
	}

	dmRequest := DeleteMerchandiseRequest{}
	if err := c.Bind(&dmRequest); err != nil {
		response.Message = "ERROR: BAD REQUEST"
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Delete(&models.Merchandise{}, dmRequest.ID).Error; err != nil {
		response.Message = "ERROR: FAILED TO DELETE MERCHANDISE"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "SUCCESS: Delete Merchandise With ID:" + strconv.Itoa(int(dmRequest.ID))
	return c.JSON(http.StatusCreated, response)
}

func CheckoutHandler(c echo.Context) error {
	role := c.Get("role").(types.Role)
	response := models.Response[[]models.Merchandise]{}

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
	condition := models.User{Usercode: request.To}
	if err := db.Where(&condition).Find(&user).Error; err != nil {
		response.Message = "ERROR: USERCODE NOT FOUND"
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

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Update User Point
	if err := tx.Model(&user).Where("id = ?", user.ID).Update("point", user.Point-totalMerchPoint).Error; err != nil {
		tx.Rollback()
		response.Message = "ERROR: FAILED TO UPDATE POINT"
		return c.JSON(http.StatusInternalServerError, response)
	}

	// Update Merch Stock
	data := []models.Merchandise{}
	for _, order := range request.Payload {
		merch := models.Merchandise{}
		if err := tx.First(&merch, order.MerchantID).Error; err != nil {
			tx.Rollback()
			response.Message = "ERROR: BAD REQUEST"
			return c.JSON(http.StatusBadRequest, response)
		}

		data = append(data, merch)

		if err := tx.Model(&merch).Where("id = ?", merch.ID).Update("stock", merch.Stock-order.Quantity).Error; err != nil {
			tx.Rollback()
			response.Message = "ERROR: FAILED TO UPDATE STOCK"
			return c.JSON(http.StatusInternalServerError, response)
		}
	}

	adminId := c.Get("id").(uint)
	userId := user.ID

	log := models.Log{
		FromId: userId,
		ToId:   adminId,
		Point:  totalMerchPoint,
	}

	if err := tx.Create(&log).Error; err != nil {
		tx.Rollback()
		response.Message = "ERROR: FAILED TO CREATE LOG"
		return c.JSON(http.StatusInternalServerError, response)
	}

	tx.Commit()
	response.Message = "SUCCESS: CHECKOUT"
	response.Data = data
	return c.JSON(http.StatusOK, response)
}
