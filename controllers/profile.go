package controllers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/guregu/null.v4"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/models"
	"itfest-backend-2.0/services"
	"itfest-backend-2.0/types"
)

type ProfileUpdateRequest struct {
	Email     null.String           `json:"email" validate:"email"`
	BirthDate types.BirthDate       `json:"birthdate"`
	Gender    null.String           `json:"gender"`
	Interests types.CareerInterests `json:"interests"`
}

func GetProfileHandler(c echo.Context) error {
	id := c.Get("id").(uint)

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

func UpdateProfileHandler(c echo.Context) error {
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

	if gender := updateProfile.Gender.String; updateProfile.Gender.Valid && gender != "male" && gender != "female" {
		response.Message = "ERROR gender should be blank, male, or female"
		return c.JSON(http.StatusBadRequest, response)
	}

	id := c.Get("id").(uint)

	db := configs.DB.GetConnection()
	profile := models.Profile{}

	if err := db.FirstOrCreate(&profile, models.Profile{UserID: id}).Error; err != nil {
		response.Message = "ERROR failed to get profile"
		return c.JSON(http.StatusBadRequest, response)
	}

	shouldGrantPoint := !profile.Submitted

	if err := db.Model(&profile).Where("user_id = ?", profile.UserID).Updates(models.Profile{
		Email:     updateProfile.Email,
		BirthDate: updateProfile.BirthDate,
		Gender:    updateProfile.Gender,
		Interests: updateProfile.Interests,
		Submitted: true,
	}).Error; err != nil {
		response.Message = "ERROR update profile"
		return c.JSON(http.StatusBadRequest, response)
	}

	if shouldGrantPoint {
		points, perr := strconv.Atoi(os.Getenv("ADD_PROFILE_POINT"))

		if perr != nil {
			return perr
		}

		admin := models.User{}

		if err := db.Where(models.User{Role: types.Admin}).First(&admin).Error; err != nil {
			return err
		}

		_, err := services.GrantPoint(admin.ID, id, uint(points))

		if err != nil {
			return err
		}
	}

	response.Message = "SUCCESS update profile"
	response.Data = profile

	return c.JSON(http.StatusOK, response)
}
