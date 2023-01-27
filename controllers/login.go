package controllers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"itfest-backend-2.0/configs"
	"itfest-backend-2.0/middlewares"
	"itfest-backend-2.0/models"
)

type LoginRequest struct {
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

func LoginHandler(c echo.Context) error {
	db := configs.DB.GetConnection()
	config := configs.Config.GetMetadata()
	response := models.Response[string]{}

	request := LoginRequest{}
	if err := c.Bind(&request); err != nil {
		response.Message = "ERROR: BAD REQUEST"
		return c.JSON(http.StatusBadRequest, response)
	}

	user := models.User{}
	username := request.Username
	condition := models.User{Username: username}
	if err := db.Where(&condition).Find(&user).Error; err != nil {
		response.Message = "ERROR: INTERNAL SERVER ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}
	if user.Username == "" {
		response.Message = "ERROR: INVALID USERNAME OR PASSWORD"
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password)); err != nil {
		response.Message = "ERROR: INVALID USERNAME OR PASSWORD"
		return c.JSON(http.StatusUnauthorized, response)
	}

	authClaims := middlewares.AuthClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.ApplicationName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.LoginExpirationDuration)),
		},
		ID: user.ID,
		Role: user.Role,
	}

	unsignedAuthToken := jwt.NewWithClaims(config.JWTSigningMethod, authClaims)
	signedAuthToken, err := unsignedAuthToken.SignedString(config.JWTSignatureKey)
	if err != nil {
		response.Message = "ERROR: JWT SIGNING ERROR"
		return c.JSON(http.StatusInternalServerError, response)
	}

	response.Message = "SUCCESS"
	response.Data = signedAuthToken
	return c.JSON(http.StatusAccepted, response)
}
