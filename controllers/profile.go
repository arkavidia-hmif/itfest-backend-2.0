package controllers

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/types"
	"net/http"
)

type ProfileUpdateRequest struct {
	Email     sql.NullString        `json:"email" validate:"email"`
	BirthDate sql.NullTime          `json:"birthdate" validate:"datetime"`
	Gender    types.Gender          `json:"gender"`
	Interests types.CareerInterests `json:"interests"`
}

type profileController struct{}

type ProfileController interface {
	GetProfile(c echo.Context) error
	UpdateProfile(c echo.Context) error
}

func NewProfileController() ProfileController {
	return &profileController{}
}

func (*profileController) GetProfile(c echo.Context) error {
	id := c.Get("id").(int)

	db := configs.DB.GetConnection()
	profile := models.Profile{}

	response := models.Response[models.Profile]{}

	if err := db.FirstOrCreate(&profile, models.Profile{UserID: id}).Error; err != nil {
		response.Message = "ERROR failed to get profile"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "SUCCESS get user"
	response.Data = profile

	return c.JSON(http.StatusOK, response)
}

func (*profileController) UpdateProfile(c echo.Context) error {
	updateProfile := new(ProfileUpdateRequest)
	response := models.Response[models.Profile]{}

	if err := c.Bind(updateProfile); err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(updateProfile); err != nil {
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	id := c.Get("id").(int)

	db := configs.DB.GetConnection()
	profile := models.Profile{}

	if err := db.FirstOrCreate(&profile, models.Profile{UserID: id}).Error; err != nil {
		response.Message = "ERROR failed to get profile"
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Model(&profile).Updates(updateProfile).Error; err != nil {
		response.Message = "ERROR update profile"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Message = "SUCCESS update profile"
	response.Data = profile

	return c.JSON(http.StatusOK, response)
}
